document.addEventListener("DOMContentLoaded", () => {
    formSection = document.querySelector(".form-section");
    const goToRegister = document.getElementById("go-to-register");
    const goToLogin = document.getElementById("go-to-login");
    const auth = document.getElementById("auth");
    slider = document.querySelector(".slider");

    
    // Detect the current page based on URL path
    const currentHash = window.location.hash;
    if (currentHash === "#login") {
        showLoginForm();
    } else if (currentHash === "#registration") {
        showSignUpForm();
    } else {
        auth.style.display = "flex";
    }

    // Link to Sign Up
    goToRegister.addEventListener("click", (e) => {
        e.preventDefault();
        showSignUpForm();
    });

    // Link to Login
    goToLogin.addEventListener("click", (e) => {
        e.preventDefault();
        showLoginForm();
    });
});

function showLoginForm() {
    auth.style.display = "flex";
    formSection.style.transform = "translateX(0)";
    slider.style.left = "0";
    if (window.location.hash !== "#login"){
        window.location.hash = "#login";
    };
}

// Function to switch to Sign Up Form
function showSignUpForm() {
    auth.style.display = "flex";
    formSection.style.transform = "translateX(-50%)";
    slider.style.left = "50%";
    if (window.location.hash !== "#registration"){
        window.location.hash = "#registration";
    };
}

async function fetching(event, endpoint) {
    event.preventDefault(); // Prevent form from refreshing the page

    const form = event.target;
    const formData = new FormData(form);
    
    // Convert FormData to JSON object
    const jsonData = Object.fromEntries(formData.entries());
    console.log(JSON.stringify(jsonData));
    try {
        const response = await fetch(endpoint, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(jsonData),
        });

        const result = await response.json();

        if (response.ok) {
            alert("Success: " + result.message);
            console.log(result);
        } else {
            alert("Error: " + result.error || "Something went wrong");
            console.error(result);
        }
    } catch (error) {
        console.error("Fetch error:", error);
        alert("Network error. Please try again.");
    }
}