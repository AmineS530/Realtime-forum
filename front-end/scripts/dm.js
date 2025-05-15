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
        const msg = JSON.parse(event.data)
        if (msg.sender !== "internal") {
            if (['system',document.getElementById("username").text,discussion.previousElementSibling.value].includes(msg.sender)) {
                msg.time = new Date()
                discussion.innerHTML += messages(msg);
            }else {
                discussion.innerHTML += `<li>[system]received a new message from ${msg.sender}.</li>`;
                showNotification("new Message from :"+msg.sender, "success", false)
            }
            playSound("message")
        } else {
            switch (msg.type) {
            case "toggle":
                discussion.previousElementSibling.querySelector(`[value="${msg.username}"]`).text = (msg.online?'🟢 ':'🔴 ') + msg.username;
                break;
            case "typing":
                discussion.previousElementSibling.querySelector(`[value="${msg.username}"]`).text = '🟢 '+ msg.username + "⌨"
                break;
            case "stoptyping":
                discussion.previousElementSibling.querySelector(`[value="${msg.username}"]`).text = '🟢 '+ msg.username
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
        // showNotification("connection lost reload the page for dms to work", "error")
    };
}

function sendMessage(message) {
    socket.send(message);
    console.log("Sent message:", message);
}

let message123 = {
    type: "broadcast",
    content: {
        receiver : "guest1",
        message : "Hello, everyone!"
    }
};

function changeDiscussion(elem) {
    elem.disabled = true;
    fetch('/api/v1/get/dmhistory', {
            method: 'GET',
            headers: {
              'target': elem.value,
              'page' : "",
              'X-Requested-With': 'XMLHttpRequest'
            }
        })
        .then(response => response!== null ?response.json():data=[])
        .then(data => {
            let formattedHistory = data && data.length===10 ?`<button onclick="window.MessageScroll()" >load more messages</button>`:"";
            if (data) {
                data.forEach(message => {
                    formattedHistory += messages(message)
                });
            }
            elem.nextElementSibling.innerHTML = formattedHistory
            elem.nextElementSibling.nextElementSibling.value = elem.value
        })
        .catch(error => console.log('Error:', error))
        .finally(elem.disabled = false);
}

function  sendDm(event) {
    let message = new Message(discussion.previousElementSibling.value,event.target[0].value);
    event.target.reset()
    message.send()
}

class Message {
    constructor(dest, contents) {
        if (typeof dest !== 'string' || !dest.trim()) {
            throw new Error("Destination must be a non-empty string");
        }

        if (typeof contents !== 'string' || !contents.trim()) {
            throw new Error("Message content must be a non-empty string");
        }

        this.body = {
            receiver: dest.trim(),
            message: escapeHTML(contents.trim())
        };
    }

    send() {
        socket.send(JSON.stringify(this.body));
    }
}

const usename = document.cookie.match(/session-name=(\S+)==;/)

let isLoadingMessages = false;

window.MessageScroll = function () {

        console.log('scrolling')
        if (isLoadingMessages) return;
        
        isLoadingMessages = true;
        try {
            elem = document.querySelector('#message-select');
            fetch('/api/v1/get/dmhistory', {
                method: 'GET',
                headers: {
                  'target': elem.value,
                  'page' : new Date(discussion.children[1].title).toISOString(),
                  'X-Requested-With': 'XMLHttpRequest'
                }
            })
            .then(response => response!== null ?response.json():data=[])
            .then(data => {
                let formattedHistory = "";
                if (data) {
                    data.forEach(message => {
                        formattedHistory += messages(message)
                    });
                }
                discussion.children[0].insertAdjacentHTML('afterend',formattedHistory)
                if (data && data.length!==10) {
                    discussion.children[0].remove()
                }
                
                elem.nextElementSibling.nextElementSibling.value = elem.value
            })
            .catch(error => console.error('Error:', error))
            .finally(elem.disabled = false);

        } catch (error) {
            showErrorPage(error.status, error.message);
        } finally {
            isLoadingMessages = false;
        }
        
};


let isTyping = false;
let typingTimer;
const typingDelay = 1000;

function startTyping() {
    if (!isTyping && !discussion.previousElementSibling.disabled) {
        socket.send(`typing:${discussion.previousElementSibling.value}`)
    }
    isTyping = true;
    
    clearTimeout(typingTimer);
    typingTimer = setTimeout(() => {
        isTyping = false;
        console.log('User stopped typing');
        socket.send(`stoptyping:${discussion.previousElementSibling.value}`)
    }, typingDelay);
};

const messages = (message) => `
<li class="message" title="${new Date(message.time).toLocaleString()}">
    <div class="message-meta">
        <span class="message-sender">${message.sender}</span>
        <span class="message-time">
            ${new Date(message.time).toLocaleTimeString([], { month: '2-digit', day: '2-digit', hour: '2-digit', minute: '2-digit' })}
        </span>
    </div>
    <div class="message-content">${escapeHTML(message.message)}</div>
</li>
`;

console.log("loaded dm.js")