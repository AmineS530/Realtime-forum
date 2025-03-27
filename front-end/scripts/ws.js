let socket;

// Connect to WebSocket
function connectWebSocket() {
    socket = new WebSocket("ws://localhost:8080/ws");

    socket.onopen = () => {
        console.log("Connected to WebSocket server");
    };

    socket.onmessage = (event) => {
        console.log("Received:", event.data);
        displayMessage(event.data); // Call a function to update the UI
    };

    socket.onclose = () => {
        console.warn("WebSocket closed. Reconnecting...");
        setTimeout(connectWebSocket, 3000); // Auto-reconnect
    };

    socket.onerror = (error) => {
        console.error("WebSocket error:", error);
    };
}

// Send message to WebSocket server
function sendMessage(message) {
    if (socket && socket.readyState === WebSocket.OPEN) {
        socket.send(message);
    } else {
        console.warn("WebSocket is not connected.");
    }
}

// Display message in the UI
function displayMessage(message) {
    const container = document.getElementById("messages");
    const msgElement = document.createElement("div");
    msgElement.textContent = message;
    container.appendChild(msgElement);
}

// Connect WebSocket when the page loads
document.addEventListener("DOMContentLoaded", connectWebSocket);
