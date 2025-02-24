function dms_ToggleShowSidebar(event) {
    console.log("dms_ShowSidebar",document.getElementById("backdrop").classList, event);
    document.getElementById("backdrop").classList.toggle("show");
    console.log("dms_ShowSidebar after",document.getElementById("backdrop").classList);
}

let socket = new WebSocket("ws://localhost:9090/api/v1/ws");

socket.onopen = function (event) {
    console.log("Connected to WebSocket server");
};

socket.onmessage = function (event) {
    console.log("Received message:", event.data);
};

socket.onclose = function (event) {
    console.log("Disconnected from WebSocket server");
};

socket.onerror = function (error) {
    console.error("WebSocket error:", error);
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
azer={}
function  sendDm(event) {
    // azer = event.;
    console.log(event.target[0].value);
    let message = new transmission("direct_message", {
        receiver : "guest1",
        message : event.target[0].value
    });
    // event.preventDefault();
    // let message = {
    //     type: "broadcast",
    //     content: {
    //         receiver : "guest1",
    //         message : messageInput.value
    //     }
    // };
    // sendMessage(JSON.stringify(message));
    // messageInput.value = "";
    // return false;
}

class transmission {
    constructor(type, contents) {
        if (typeof type !== 'string') {
            throw new Error("Type must be a string");
        }
        
        // Ensure 'contents' is defined
        if (contents === undefined || contents === null) {
            throw new Error("Contents must be defined");
        }

        // Initialize the body object
        this.body = {};
        this.body.type = type;
        this.body.content = contents;

        this.setHandler(type);
    }

    setHandler(type) {
        // Set the appropriate handler for the given type
        switch (type) {
            case "direct_message":
                this.verify = this.directMessageverify;
                break;
            default:
                throw new Error("Unsupported transmission type");
        }
    }

    directMessageverify() {
        if (typeof this.body.content.receiver!=='string') {
            return false
        }

        if (typeof this.body.content.message!=='string') {
            return false
        }
    }

    send() {
        socket.send(JSON.stringify(this.body));
        console.log("Sent message:", JSON.stringify(this.body));
    }
}