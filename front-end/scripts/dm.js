function dms_ToggleShowSidebar(event) {
    document.getElementById("backdrop").classList.toggle("show");
}

let socket = new WebSocket(`ws://${window.location.host}/api/v1/ws`);

socket.onopen = function (event) {
    console.log("Connected to WebSocket server");
};

socket.onmessage = function (event) {
    const msg = JSON.parse(event.data)
    console.log("Received message:", event.data,"Parsed message:", msg);
    discussion.innerHTML += ['system',window.uname,discussion.previousElementSibling.value.slice(3)].includes(msg.sender)?
                            `<li>[${msg.sender}] : ${msg.message}</li>`:
                            `<li>[system]received a new message from ${msg.sender}.</li>`;
};

socket.onclose = function (event) {
    console.log("Disconnected from WebSocket server");
};

socket.onerror = function (error) {
    console.error("WebSocket error:", error);
    // showNotification("connection lost reload the page for dms to work", "error")
};

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
              'target': elem.value.slice(3),
              'page' : elem.nextElementSibling.childElementCount-1/10,
              'X-Requested-With': 'XMLHttpRequest'
            }
        })
        .then(response => response!== null ?response.json():data=[])
        .then(data => {
            console.log("azerazerazernbfhqbfhbqfhqbsfjhbqjsbfhqbsdfhqbsdfq",data)
            let formattedHistory = `<button onclick="window.setupMessageScroll()" >load more messages</button>`;
            if (data) {
                data.forEach(message => {
                    formattedHistory += `<li>[${message.sender}] : ${message.message}</li>`
                });
            }
            console.log("azerazerazerazerazer",formattedHistory)
            elem.nextElementSibling.innerHTML = formattedHistory
            elem.nextElementSibling.nextElementSibling.value = elem.value
        })
        .catch(error => console.error('Error:', error))
        .finally(elem.disabled = false);
}

function  sendDm(event) {
    // console.log(event.target.attributes.value.value)
    console.log("azer",event.target.value,event.target[0].value,event);
    let message = new Message(event.target.value.slice(3),event.target[0].value);
    message.send()
}

class Message {
    constructor(dest,contents) {
        if (typeof dest!=='string') {
            throw new Error("Type must be a string");        }

        if (typeof contents!=='string') {
            throw new Error("Type must be a string");
        }
        // Initialize the body object
        this.body = {receiver: dest,message: contents};
    }

    send() {
        socket.send(JSON.stringify(this.body));
    }
}
console.log("loaded dm.js")

const usename = document.cookie.match(/session-name=(\S+)==;/)