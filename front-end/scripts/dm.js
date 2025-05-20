function dms_ToggleShowSidebar(event) {
    const backdrop = document.getElementById("backdrop")
    if (backdrop) {
        backdrop.classList.toggle("show");
        if (document.getElementById("chat-box").title != "") {
            closeChat()
        }
    } else {
        showNotification("Go back to home page to send a message", "info");
    }
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
                    discussion.previousElementSibling.title,
                ].includes(msg.sender)
            ) {
                msg.time = new Date();
                discussion.innerHTML += messages(msg);
                discussion.scrollTop = discussion.scrollHeight;
            } else {
                // discussion.innerHTML += `<li>[system] You have received a new message from ${msg.sender}.</li>`;
                showNotification("new Message from " + msg.sender, "success", false);
            }
            if (msg.sender !== "system") {
                let container = document.querySelector(".chat-users")
                container.insertBefore(
                    container.querySelector(`[username="${msg.sender}"]`)||container.querySelector(`[username="${msg.receiver}"]`),
                    container.firstChild
                )
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
                    if (msg.username == document.getElementById("chat-username").textContent) {
                        document.getElementById("typing").classList.remove("hidden");
                    }
                    break;
                case "stoptyping":
                    if (msg.username == document.getElementById("chat-username").textContent) {
                        document.getElementById("typing").classList.add("hidden");
                    }
                    break;
            }
        }
    };

    socket.onclose = function (event) {
        console.log("Disconnected from WebSocket server");
        document.getElementById("backdrop").classList.add("hidden");
        document.getElementById("nav").classList.add("hidden");
        document.body.classList.add("dimmed");
        const errorBannerHTML = `
        <div id="error-banner" class="error-banner">
            <p>Connection lost. Please refresh the page to reconnect.</p>
            <button id="error-ok-btn">OK</button>
        </div>`;
        document.body.insertAdjacentHTML("afterbegin", errorBannerHTML);
        document.getElementById("error-ok-btn").addEventListener("click", () => {
            window.location.reload();
        });

        socket.onerror = function (error) {
            console.error("WebSocket error:", error);
            discussion.innerHTML += `<li>[system]: web socket error ${error}.</li>`;
        };
    }
}

const sendDm = throttle(function (event) {
    let receiver = document.getElementById("chat-username").textContent;
    let input = String(event.target[0].value)
    let message = new Message(receiver, input.trim());
    event.target.reset();
    if (message) {
        message.send();
    }
}, 200);

class Message {
    constructor(dest, contents) {
        if (typeof dest !== "string" || !dest.trim()) {
            throw new Error("Destination must be a non-empty string");
        }

        if (typeof contents !== "string") {
            showNotification("Cannot send empty message", "info");
            return false;
        }

        this.body = {
            receiver: dest.trim(),
            message: escapeHTML(contents.trim()),
        };
    }

    send() {
        if (!(this.body.receiver.trim()) || !(this.body.message.trim())) {
            showNotification("couldn't send message", "error")
        } else {
            socket.send(JSON.stringify(this.body));
        }
    }
}

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

async function fetchDMHistory(username, page = "") {
    try {
        const response = await fetch("/api/v1/get/dmhistory", {
            method: "GET",
            headers: {
                target: username,
                page: page,
                "X-Requested-With": "XMLHttpRequest",
            },
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        return await response.json();
    } catch (error) {
        throw error;
    }
}

let hasMoreMessages = true;
let lastFetchedTimestamp = null;
let isLoadingMessages = false;

async function loadMoreMessages() {
    if (isLoadingMessages) return;

    const discussionElem = document.getElementById("discussion");
    if (!discussionElem) return;

    // Check if we're near the top (within 100px)
    if (discussionElem.scrollTop > 50) return;

    isLoadingMessages = true;

    const username = document.getElementById("chat-username")?.textContent;
    if (!username) {
        isLoadingMessages = false;
        return;
    }

    try {
        const oldestMessage = discussionElem.firstElementChild;
        if (!oldestMessage) {
            isLoadingMessages = false;
            return;
        }

        // Get timestamp from data-timestamp attribute (must be ISO format)
        const oldestMessageTimestamp = oldestMessage.getAttribute("data-timestamp");

        // Prevent fetching same messages again
        if (lastFetchedTimestamp === oldestMessageTimestamp) {
            isLoadingMessages = false;
            return;
        }
        lastFetchedTimestamp = oldestMessageTimestamp;

        // Store scroll position before loading
        const oldScrollHeight = discussionElem.scrollHeight;
        const oldScrollTop = discussionElem.scrollTop;

        const data = await fetchDMHistory(username, oldestMessageTimestamp);

        if (data && data.length > 1) {
            let formattedHistory = "";
            data.forEach((message) => {
                formattedHistory += messages(message);
            });

            discussionElem.insertAdjacentHTML("afterbegin", formattedHistory);
            if (data.length < 10) {
                hasMoreMessages = false;
                showNotification("No more messages!", "info");
            }
            // Maintain scroll position
            const newScrollHeight = discussionElem.scrollHeight;
            discussionElem.scrollTop = newScrollHeight - oldScrollHeight + oldScrollTop;
        } else {
            hasMoreMessages = false;
        }
    } catch (error) {
        console.error("Error:", error);
        showNotification("Error loading messages. Please try again later.", "error");
    } finally {
        isLoadingMessages = false;
    }
}

async function changeDiscussion(username) {
    // Reset tracking when changing discussions
    hasMoreMessages = true;
    lastFetchedTimestamp = null;

    try {
        const data = await fetchDMHistory(username);
        const discussionElem = document.getElementById("discussion");

        let formattedHistory = "";
        if (data && data.length > 0) {
            data.forEach((message) => {
                formattedHistory += messages(message);
            });
        }
        discussionElem.innerHTML = formattedHistory;
        setupScrollListener();
        discussionElem.scrollTop = discussionElem.scrollHeight;
    } catch (error) {
        console.error("Error:", error);
        showNotification("Error loading messages", "error");
    }
}

function setupScrollListener() {
    const discussionElem = document.getElementById("discussion");
    if (!discussionElem) return;

    // Throttled scroll handler
    window.scrollListener = throttle(() => {
        if (discussionElem.scrollTop <= 50) {
            loadMoreMessages();
        }
    }, 400);

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
<li class="message" title="${new Date(message.time).toLocaleString()}" data-timestamp="${message.time}">
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