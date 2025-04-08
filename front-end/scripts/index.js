import { two_bubbles, svg_logout } from "./svg.js";
import { updateNavbar } from "./header.js";

const nav = `
<link rel="stylesheet" href="/front-end/styles/style.css" />

<ul class="nav">
    <li>
        <a href="#" onclick="dms_ToggleShowSidebar(event)" title="Messages">
            ${two_bubbles}
        </a>
    </li>
    <li>
        <a id="logout" href="#" title="Logout">
            ${svg_logout}
        </a>
    </li>
</ul>
`;
const header =`
    <header>
      <!-- Logo -->
      <div class="logo">
        <h1><a href="/">Forum</a></h1>
      </div>
      <!-- Navigation -->
      <nav>
        <ul class="nav"></ul>
      </nav>
    </header>
    <br />
`

const auth = `
<link rel="stylesheet" href="/front-end/styles/log-reg.css" />
<div id="auth">
    <div class="container">
        <!-- Buttons -->
        <div class="btn">
            <button id="login-btn" onclick="showLoginForm()">Login</button>
            <button id="register-btn" onclick="showSignUpForm()">Sign Up</button>
        </div>

        <!-- Slider -->
        <div class="slider"></div>
        <!-- Form Section -->
        <div class="form-section">
            <!-- Login Form -->
            <div class="login-box">
                <form onsubmit="fetching(event,'/api/login')" method="post">
                    <label for="name_or_email">Username or Email</label>
                    <input type="text" id="name_or_email" name="name_or_email" required />
                    <label for="password">Password</label>
                    <input type="password" id="logpassword" name="password" required />
                    <button type="submit">Login</button>
                    <p>
                        Don't have an account?
                        <a style="cursor: pointer;" onclick="showSignUpForm()" >Sign Up</a>
                    </p>
                </form>
            </div>

            <!-- Sign Up Form -->
            <div class="register-box">
                <form onsubmit="fetching(event,'/api/register')" method="post">
                    <label for="username">Username</label>
                    <input type="text" id="username" name="Username" required />
                    <label for="email">Email</label>
                    <input type="email" id="email" name="Email" required />
                    <label for="password">Password</label>
                    <input type="password" id="regpassword" name="Password" required />
                    <label for="password_confirmation">Confirm Password</label>
                    <input type="password" id="password_confirmation" name="Password_confirmation" required />
                    <label for="age">Age</label>
                    <input type="number" id="age" name="Age" min="15" max="90" required />
                    <label for="gender">Gender </label>
                    <select id="gender" name="Gender" required>
                        <option value="">Select Gender</option>
                        <option value="male">Male</option>
                        <option value="female">Female</option>
                        <option value="Attack helicopter">Attack helicopter</option>
                    </select>
                    <label for="first-name">First Name</label>
                    <input type="text" id="first-name" name="First_Name" required />
                    <label for="last-name">Last Name</label>
                    <input type="text" id="last-name" name="Last_Name" required />
                    <button type="submit">Sign Up</button>
                    <p>
                        Already have an account? <a style="cursor: pointer;" onclick="showLoginForm()" >Login</a>
                    </p>
                </form>
            </div>
        </div>
    </div>
</div>
`
function injectStylesheet(href) {
    if (!document.querySelector(`link[href="${href}"]`)) {
        const link = document.createElement('link');
        link.rel = 'stylesheet';
        link.href = href;
        document.head.appendChild(link);
    }
}

// on load

// function to load the header and nav
document.addEventListener("DOMContentLoaded", async () => {
    const app = document.getElementById("app");

    try {
        const response = await fetch("/api/check-auth", {
            credentials: "include"
        });

        const authData = await response.json();
        const authenticated = authData.authenticated;;

        console.log("Auth data:", authData);

        if (authenticated) {
            const headerWrapper = document.createElement("div");
            headerWrapper.innerHTML = header;
            document.body.insertBefore(headerWrapper, app);

            injectStylesheet("/front-end/styles/header.css");

            updateNavbar(authenticated);
        } else {
            app.innerHTML = auth
        }
    } catch (error) {
        console.error("Error loading content:", error);
        app.innerHTML = "<h2>Error loading the page. Please try again.</h2>";
    }
});

// function loadPrivateMessages() {
//     document.getElementById("private-messages-container").innerHTML = "Loading messages...";
//     fetch("/api/messages")
//         .then(res => res.text())
//         .then(html => {
//             document.getElementById("private-messages-container").innerHTML = html;
//         });
// }
