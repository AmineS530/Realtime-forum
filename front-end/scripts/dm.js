function dms_ToggleShowSidebar(event) {
    document.getElementById("backdrop").classList.toggle("show");
}
function playNotificationSound() {
    // Create a new Audio object and specify the sound file
    var audio = new Audio('https://www.soundjay.com/buttons/sounds/button-10.mp3'); // You can use a different URL or a local file

    // Play the sound
    audio.play();
}

let socket;
window.retrysocket = function () {
    socket = new WebSocket(`ws://${window.location.host}/api/v1/ws`);
    
    socket.onopen = function (event) {
        console.log("Connected to WebSocket server");
    };
    
    socket.onmessage = function (event) {
        const msg = JSON.parse(event.data)
        console.log("Received message:", event.data,"Parsed message:", msg);
        if (msg.sender !== "internal") {
            if (['system',document.getElementById("username").text,discussion.previousElementSibling.value].includes(msg.sender)) {
                discussion.innerHTML += `<li>[${msg.sender}] : ${msg.message}</li>`;
            }else {
                discussion.innerHTML += `<li>[system]received a new message from ${msg.sender}.</li>`;
                showNotification("new Message from :"+msg.sender)
            }
            playNotificationSound();
        } else {
            if (msg.type == "toggle") {
                discussion.previousElementSibling.querySelector(`[value="${msg.username}"]`).text = (msg.online?'🟢 ':'🔴 ') + msg.username;
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
              'page' : 0,
              'X-Requested-With': 'XMLHttpRequest'
            }
        })
        .then(response => response!== null ?response.json():data=[])
        .then(data => {
            console.log("azerazerazernbfhqbfhbqfhqbsfjhbqjsbfhqbsdfhqbsdfq",data)
            let formattedHistory = data && data.length===10 ?`<button onclick="window.MessageScroll()" >load more messages</button>`:"";
            if (data) {
                data.forEach(message => {
                    formattedHistory += `<li title="${new Date(message.time).toLocaleString()}">[${message.sender}] : ${message.message}</li>`
                });
            }
            console.log("azerazerazerazerazer",formattedHistory)
            elem.nextElementSibling.innerHTML = formattedHistory
            elem.nextElementSibling.nextElementSibling.value = elem.value
        })
        .catch(error => console.log('Error:', error))
        .finally(elem.disabled = false);
}

function  sendDm(event) {
    // console.log(event.target.attributes.value.value)
    console.log("azer",event.target.previousElementSibling.previousElementSibling.value,event.target[0].value,event);
    let message = new Message(event.target.previousElementSibling.previousElementSibling.value,event.target[0].value);
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
                  'page' : (elem.nextElementSibling.childElementCount-1)/10 | 0,
                  'X-Requested-With': 'XMLHttpRequest'
                }
            })
            .then(response => response!== null ?response.json():data=[])
            .then(data => {
                let formattedHistory = "";
                if (data) {
                    data.forEach(message => {
                        formattedHistory += `<li title="${new Date(message.time).toLocaleString()}">[${message.sender}] : ${message.message}</li>`
                    });
                }
                discussion.children[0].insertAdjacentHTML('afterend',formattedHistory)
                if (data.length!==10) {
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


console.log("loaded dm.js")