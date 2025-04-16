import templates from "./templates.js";
import { updateNavbar } from "./header.js";
window.loadPage = function (page) {
    const app = document.getElementById("app");
    switch (page) {
        case "home":
            app.innerHTML = templates.posts + templates.dms;
            history.pushState({}, "", "/");
            break;
        case "profile":
            loadProfilePage();
            break;
        case "create":
            app.innerHTML = createPost;
            break;
        default:
            app.innerHTML = "<h2>Page not found</h2>";
    }
};

// import { updateNavbar } from "./header.js";
function loadPageFromPath() {
    const path = window.location.pathname;
    switch (path) {
        case "/":
        case "/home":
            loadPage("home");
            break;
        case "/profile":
            loadPage("profile");
            break;
        case "/create":
            loadPage("create");
            break;
        default:
            loadPage("notfound"); // Optional fallback
    }
}
window.addEventListener("popstate", loadPageFromPath);

async function loadProfilePage() {
    const app = document.getElementById("app");
    try {
        const res = await fetch("/api/profile", { credentials: "include" });
        if (!res.ok) {
            const errData = await res.json();
            throw new Error(errData.error || "Error loading profile");
        }

        const data = await res.json();
        app.innerHTML = ` 
        <center>
        <div class="profile">
        <h2>Welcome, ${data.username}</h2>
        <p>First Name: ${data.first_name}</p>
        <p>Last Name: ${data.last_name}</p>
        <p>Age: ${data.age}</p>
        <p>Gender: ${data.gender}</p>
        <p>Email: ${data.email}</p>
        <a class="change-password" href="#" > Change Password </a>
        <br />
        <a href="/" onclick="loadPage('home'); return false;">Go Back</a>
        </div>
        </center>`;

        history.pushState({}, "", "/profile");
    } catch (err) {
        console.error("Failed to load profile:", err);
        app.innerHTML = "<h2>Error loading profile.</h2>";
    }
}

let profile = document.getElementById("profile");
if (profile) {
    profile.addEventListener("click", function (e) {
        e.preventDefault();
        loadPage("profile");
        history.pushState(null, "", "/profile");
    });
}

// function to load the header and nav
document.addEventListener("DOMContentLoaded", async () => {
    const app = document.getElementById("app");

    try {
        const response = await fetch("/api/check-auth", {
            credentials: "include",
        });
        const authData = await response.json();
        const authenticated = authData.authenticated;

        if (authenticated) {
            setHeader(authenticated);

            loadPageFromPath();
        } else {
            app.innerHTML = templates.auth;
        }
    } catch (error) {
        console.error("Error loading:", error);
        app.innerHTML = "<h2>Error loading the page.</h2>";
    }
});

function setHeader(authStatus) {
    const headerWrapper = document.createElement("div");
    headerWrapper.innerHTML = templates.header;
    document.body.insertBefore(headerWrapper, app);
    injectStylesheet("/front-end/styles/header.css");
    injectStylesheet("/front-end/styles/dms.css");
    injectStylesheet("/front-end/styles/style.css");
    updateNavbar(authStatus);
}

function injectStylesheet(href) {
    if (!document.querySelector(`link[href="${href}"]`)) {
        const link = document.createElement("link");
        link.rel = "stylesheet";
        link.href = href;
        document.head.appendChild(link);
    }
}
