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

function showNotification(message, type = "OK") {
    const notification = document.createElement("div");
    notification.innerText = message;
    notification.className = `notification ${type}`;
    document.body.appendChild(notification);
    setTimeout(() => notification.remove(), 3000);
}
