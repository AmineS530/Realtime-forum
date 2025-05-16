function dms_ToggleShowSidebar(event) {
    document.getElementById("backdrop").classList.toggle("show");
}

let socket;

window.retrysocket = function () {
    socket = new WebSocket(`ws://${window.location.host}/api/v1/ws`);

    socket.onopen = function (event) {
        console.log("Connected to WebSocket server");
    };

    socket.onmessage = function (event) {
        const msg = JSON.parse(event.data);
        if (msg.sender !== "internal") {
            if (
                [
                    "system",
                    document.getElementById("username").text,
                    discussion.previousElementSibling.value,
                ].includes(msg.sender)
            ) {
                msg.time = new Date();
                discussion.innerHTML += messages(msg);
            } else {
                discussion.innerHTML += `<li>[system] You have received a new message from ${msg.sender}.</li>`;
                showNotification("new Message from :" + msg.sender, "success", false);
            }
            playSound("message");
        } else {
            switch (msg.type) {
                case "toggle":
                    document.querySelector(
                        `.chat-user[username="${msg.username}"]`
                    ).innerHTML = (msg.online ? "🟢 " : "🔴 ") + msg.username;
                    break;
                case "typing":
                    document.getElementById("chat-username").textContent = msg.username ;
                    document.getElementById("typing").classList.remove("hidden");
                    break;
                case "stoptyping":
                    document.getElementById("chat-username").textContent = msg.username;
                    document.getElementById("typing").classList.add("hidden");
                    break;
            }
        }
    };

    socket.onclose = function (event) {
        console.log("Disconnected from WebSocket server");
        discussion.innerHTML += `<li>[system]: web socket closed refresh the page to be able to send messages.</li>`;
    };

    socket.onerror = function (error) {
        console.error("WebSocket error:", error);
        discussion.innerHTML += `<li>[system]: web socket eror ${error}.</li>`;
    };
}

function sendMessage(message) {
    socket.send(message);
    console.log("Sent message:", message);
}

function sendDm(event) {
    let receiver = document.getElementById("chat-username").textContent;
    let message = new Message(receiver, event.target[0].value);
    event.target.reset();
    message.send();
}

class Message {
    constructor(dest, contents) {
        if (typeof dest !== "string" || !dest.trim()) {
            throw new Error("Destination must be a non-empty string");
        }

        if (typeof contents !== "string" || !contents.trim()) {
            throw new Error("Message content must be a non-empty string");
        }

        this.body = {
            receiver: dest.trim(),
            message: escapeHTML(contents.trim()),
        };
    }

    send() {
        socket.send(JSON.stringify(this.body));
    }
}

let isLoadingMessages = false;

function throttle(func, limit) {
    let lastCall = 0;
    return function (...args) {
        const now = Date.now();
        if (now - lastCall >= limit) {
            lastCall = now;
            func(...args);
        }
    };
}

function loadMoreMessages() {
    if (isLoadingMessages) return;

    // Only trigger when user is near the top of the container
    const discussionElem = document.getElementById("discussion");
    if (discussionElem.scrollTop > 50) return;

    isLoadingMessages = true;
    const selectElem = document.getElementById("chat-username");
    const username = selectElem.textContent;

    const oldestMessageTimestamp = new Date(
        discussionElem.children[0].title
    ).toISOString();
    // Record scroll position before loading new content
    const oldScrollHeight = discussionElem.scrollHeight;
    const oldScrollTop = discussionElem.scrollTop;

    try {
        fetch("/api/v1/get/dmhistory", {
            method: "GET",
            headers: {
                target: username,
                page: oldestMessageTimestamp,
                "X-Requested-With": "XMLHttpRequest",
            },
        })
            .then((response) => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.json();
            })
            .then((data) => {
                // Format the additional messages
                let formattedHistory = "";
                if (data && data.length > 0) {
                    data.forEach((message) => {
                        formattedHistory += messages(message);
                    });

                    // Insert new messages at the beginning
                    discussionElem.insertAdjacentHTML("afterbegin", formattedHistory);

                    // Maintain scroll position after adding content
                    discussionElem.scrollTop =
                        discussionElem.scrollHeight - oldScrollHeight + oldScrollTop;

                    // Only show "no more messages" if we got fewer than 10
                    if (data.length < 1) {
                        showNotification("No more messages!", "info");
                    }
                } else {
                    showNotification("No more messages!", "info");
                }
            })
            .catch((error) => {
                console.error("Error:", error);
                showNotification(
                    "Error loading messages. Please try again later.",
                    "error"
                );
            })
            .finally(() => {
                isLoadingMessages = false;
            });
    } catch (error) {
        showErrorPage(error.status, error.message);
        isLoadingMessages = false;
    }
}

function changeDiscussion(username) {
    const selectElem = document.getElementById("message-select");
    selectElem.disabled = true;

    fetch("/api/v1/get/dmhistory", {
        method: "GET",
        headers: {
            target: username,
            page: "",
            "X-Requested-With": "XMLHttpRequest",
        },
    })
        .then((response) => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then((data) => {
            // Format and append each message
            let formattedHistory = "";
            if (data && data.length > 0) {
                data.forEach((message) => {
                    formattedHistory += messages(message);
                });
            }

            // Update the discussion container
            const discussionElem = document.getElementById("discussion");
            discussionElem.innerHTML = formattedHistory;

            // Update selected user in any dependent elements
            if (
                selectElem.nextElementSibling &&
                selectElem.nextElementSibling.nextElementSibling
            ) {
                selectElem.nextElementSibling.nextElementSibling.value = username;
            }

            // Set up scroll event listener after loading initial messages
            setupScrollListener();
        })
        .catch((error) => {
            console.error("Error:", error);
            showNotification("Error loading messages.", "error");
        })
        .finally(() => {
            selectElem.disabled = false;
        });
}

function setupScrollListener() {
    const discussionElem = document.getElementById("discussion");

    // Remove any existing listener to prevent duplicates
    if (window.scrollListener) {
        discussionElem.removeEventListener("scroll", window.scrollListener);
    }

    // Add new scroll event listener with throttling
    window.scrollListener = throttle(loadMoreMessages, 500);
    discussionElem.addEventListener("scroll", window.scrollListener);
}

let isTyping = false;
let typingTimer;
const typingDelay = 1000;

function startTyping() {
    if (!isTyping && !discussion.previousElementSibling.disabled) {
        socket.send(`typing:${document.getElementById("chat-username").textContent}`);
    }
    isTyping = true;

    clearTimeout(typingTimer);
    typingTimer = setTimeout(() => {
        isTyping = false;
        socket.send(`stoptyping:${document.getElementById("chat-username").textContent}`);
    }, typingDelay);
}

const messages = (message) => `
<li class="message" title="${new Date(message.time).toLocaleString()}">
    <div class="message-meta">
        <span class="message-sender">${message.sender}</span>
        <span class="message-time">
            ${new Date(message.time).toLocaleTimeString([], {
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
})}
        </span>
    </div>
    <div class="message-content">${escapeHTML(message.message)}</div>
</li>
`;

console.log("loaded dm.js");