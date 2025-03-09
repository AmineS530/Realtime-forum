document.addEventListener("DOMContentLoaded", function () {
    // Select all input fields and textareas
    const inputs = document.querySelectorAll("input, textarea");

    inputs.forEach((input) => {
        if (input.type !== "password") {
            input.addEventListener("blur", function () {
                this.value = this.value.replace(/^\s+|\s+$/g, ""); // Remove leading and trailing spaces
            });
        }
    });
});

var Logout = document.getElementById("logout");
if (typeof Logout != "undefined" && Logout != null) {
    Logout.addEventListener("click", function (event) {
        event.preventDefault()
        fetch("/api/logout", {
            method: "POST",
            credentials: "include"
        })
            .then(() => {
                window.location.href = "/";
            })
            .catch((error) => console.error("Logout failed:", error));
    });
}


function showNotification(message, type = "OK") {
    const notification = document.createElement("div");
    notification.innerText = message;
    notification.className = `notification ${type}`; // Add styling
    document.body.appendChild(notification);
    setTimeout(() => notification.remove(), 3000); // Auto-remove after 3s
}
