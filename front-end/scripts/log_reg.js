let current = "login";

document.addEventListener("DOMContentLoaded", () => {
    const formSection = document.querySelector(".form-section");
    const auth = document.getElementById("auth");
    const slider = document.querySelector(".slider");

    if (!auth || !formSection || !slider) {
        console.error("Required elements not found in DOM");
        return;
    }

    if (current === "login") {
        showLoginForm();
    } else if (current === "registration") {
        showSignUpForm();
    } else {
        auth.style.display = "flex";
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

// Notification function
function showNotification(message, type = "OK") {
    const notification = document.createElement("div");
    notification.innerText = message;
    notification.className = `notification ${type}`;
    document.body.appendChild(notification);
    setTimeout(() => notification.remove(), 3000);
}

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

// // Function to check authentication status and show login form if needed
// async function fetchData(endpoint) {
//     try {
//         const response = await fetch(endpoint, {
//             method: "GET",
//             credentials: "include",
//         });

//         const result = await response.json();

//         if (response.status === 401 && result.login === "required") {
//             showNotification("You are not logged in.", "warning");
//             showLoginForm();
//             return;
//         }

//         console.log("Fetched data:", result);
//     } catch (error) {
//         showNotification("Error fetching data.", "error");
//     }
// }
