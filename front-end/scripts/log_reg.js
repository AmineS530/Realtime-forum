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
        console.log(e, 'azer')
        e.preventDefault();
        showLoginForm();
    };

    // Link to Login
    // goToLogin.addEventListener("click", (e) => {
    //     e.preventDefault();
    //     showLoginForm();
    // });
});

function fetching(e, target) {
    e.preventDefault();
    let data = {};
    for (const element of e.target) {
        if (element.value != '') {
            console.log(element.value, element.id, element);
            data[element.id] = element.value;
        }
    }
    console.log(e.target, e.this, data, target);
    let a = fetch(target, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    }).then((response) => {
        if (!response.ok) {
            throw new Error(' Network error \n\tstatus code :'+ response.status);
        }
        return response.json();
    }).catch((error) => {
            console.log('Error:', error);
            // Return error object with a message
            return { 'error': error.message };
    });
    a.then((result) => {
        console.log('Response as object:', result);
        for (let key in result) {
            if (result.hasOwnProperty(key)) {
                alert(`${key}: ${result[key]}`);
            }
        }
    });
}