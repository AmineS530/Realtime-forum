let current = "login";

document.addEventListener("DOMContentLoaded", () => {
    const formSection = document.querySelector(".form-section");
    const auth = document.getElementById("auth");
    const slider = document.querySelector(".slider");

    if (!auth || !formSection || !slider) {
        console.error("Required elements not found in DOM");
        return;
    }

    // Detect the current page based on URL path
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

// Function to switch to Sign Up Form
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
    // Select all input fields and textareas
    const inputs = document.querySelectorAll("input, textarea");

    inputs.forEach((input) => {
        if (input.type !== "password") {
            input.addEventListener("blur", function () {
                this.value = this.value.replace(/^\s+|\s+$/g, "");
            });
        }
    });
});

async function fetching(event, endpoint) {
    event.preventDefault(); // Prevent form from refreshing the page

    const form = event.target;
    const formData = new FormData(form);

    // Convert FormData to JSON object
    const jsonData = Object.fromEntries(formData.entries());

    try {
        const response = await fetch(endpoint, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(jsonData),
        });

        const contentType = response.headers.get("content-type");
        let result;

        if (contentType && contentType.includes("application/json")) {
            result = await response.json();
        } else {
            result = { error: await response.text() };
        }

        if (response.ok) {
            alert("Success: " + (result.message || "Operation successful"));
            console.log(result);
        } else {
            alert("Error: " + (result.error || "Something went wrong"));
            console.error(result);
        }
    } catch (error) {
        console.error("Fetch error:", error);
        alert("Network error. Please try again.");
    }
    console.log(jsonData);
}
