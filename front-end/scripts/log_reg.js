document.addEventListener("DOMContentLoaded", () => {
    const loginBtn = document.getElementById("login-btn");
    const registerBtn = document.getElementById("register-btn");
    const goToRegister = document.getElementById("go-to-register");
    const goToLogin = document.getElementById("go-to-login");
    const formSection = document.querySelector(".form-section");
    const slider = document.querySelector(".slider");

    // Function to switch to Login Form
    function showLoginForm() {
        formSection.style.transform = "translateX(0)";
        slider.style.left = "0";
    }

    // Function to switch to Sign Up Form
    function showSignUpForm() {
        formSection.style.transform = "translateX(-50%)";
        slider.style.left = "50%";
    }

    // Detect the current page based on URL path
    const currentPath = window.location.hash;
    if (currentPath === "#login") {
        showLoginForm();
    } else if (currentPath === "#registration") {
        showSignUpForm();
    }

    // Event Listeners for Buttons
    loginBtn.addEventListener("click", () => {
        showLoginForm();
    });

    registerBtn.addEventListener("click", () => {
        showSignUpForm();
    });

    // Link to Sign Up
    goToRegister.addEventListener("click", (e) => {
        e.preventDefault();
        showSignUpForm();
    });

    function goToRegistera(e) {
        console.log(e,'azer')
        e.preventDefault();
        showLoginForm();
    };

    // Link to Login
    // goToLogin.addEventListener("click", (e) => {
    //     e.preventDefault();
    //     showLoginForm();
    // });
});
