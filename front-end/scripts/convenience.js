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
