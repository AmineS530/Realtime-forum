let current = "login";

document.addEventListener("DOMContentLoaded", () => {
    const formSection = document.querySelector(".form-section");
    const auth = document.getElementById("auth");
    const slider = document.querySelector(".slider");

    if (auth && formSection && slider) {
        if (current === "login") {
            showLoginForm();
        } else if (current === "registration") {
            showSignUpForm();
        } else {
            auth.style.display = "flex";
        }
    }
});

function showLoginForm() {
    const auth = document.getElementById("auth");
    const formSection = document.querySelector(".form-section");
    const slider = document.querySelector(".slider");

    if (!auth || !formSection || !slider) return;

    auth.style.display = "flex";
    formSection.style.transform = "translateX(0)";
    slider.style.left = "0";
    current = "login";
}

function showSignUpForm() {
    const auth = document.getElementById("auth");
    const formSection = document.querySelector(".form-section");
    const slider = document.querySelector(".slider");

    if (!auth || !formSection || !slider) return;

    auth.style.display = "flex";
    formSection.style.transform = "translateX(-50%)";
    slider.style.left = "50%";
    current = "registration";
}

document.addEventListener("DOMContentLoaded", function () {
    const inputs = document.querySelectorAll("input, textarea");

    inputs.forEach((input) => {
        if (input.type !== "password") {
            input.addEventListener("blur", function () {
                this.value = this.value.trim();
            });
        }
    });
});



// Function to handle login requests
async function fetching(event, endpoint) {
    event.preventDefault();

    const form = event.target;
    const formData = new FormData(form);
    const jsonData = Object.fromEntries(formData.entries());

    try {
        const response = await fetch(endpoint, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(jsonData),
            credentials: "include",
        });

        const contentType = response.headers.get("content-type");
        let result;
        if (contentType && contentType.includes("application/json")) {
            result = await response.json();
        } else {
            result = { error: await response.text() };
        }

        if (response.ok) {
            showNotification("Login successful!", "success");
            setTimeout(() => {
                window.location.href = "/";
            }, 1000);
        } else {
            showNotification("Login failed: " + result.error, "error");
        }
    } catch (error) {
        showNotification("Network error. Please try again.", "error");
    }
}
