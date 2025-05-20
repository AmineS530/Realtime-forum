import templates from "./templates.js";
import { updateNavbar } from "./header.js";

window.loadPage = function (page, event) {
    if (event) {
        event.preventDefault();
    }

    const app = document.getElementById("app");
    switch (page) {
        case "home":
            setHeader(true);
            app.innerHTML = templates.posts + templates.dms + templates.postCreation;
            loadUsers();
            loadPosts({ mode: "replace" });
            setupPostCreator();
            history.pushState({}, "", "/");
            break;
        case "profile":
            loadProfilePage();
            break;
        default:
            showErrorPage(404, "Page not found");
    }
};

function loadPageFromPath() {
    const path = window.location.pathname;
    const app = document.getElementById("app");
    switch (path) {
        case "/":
        case "/home":
            loadPage("home");
            break;
        case "/profile":
            loadPage("profile");
            break;
        default:
            showErrorPage(404, "Page not found");
    }
}


async function loadProfilePage() {
    const app = document.getElementById("app");
    try {
        const res = await fetch("/api/profile", {
            method: "GET",
            credentials: "include",
            headers: {
                'X-Requested-With': 'XMLHttpRequest',
                "Content-Type": "application/json",
            },
        });

        if (!res.ok) {
            const errData = await res.json();
            throw new Error(errData.error || "Error loading profile");
        }

        const data = await res.json();
        app.innerHTML = ` 
        <link rel="stylesheet" href="/front-end/styles/style.css" />
        <div class="profile">
        <h2>Welcome, ${data.username}</h2>
        <p>First Name: ${data.first_name}</p>
        <p>Last Name: ${data.last_name}</p>
        <p>Age: ${data.age}</p>
        <p>Gender: ${data.gender}</p>
        <p>Email: ${data.email}</p>
        <a href="/" onclick="loadPage('home', event)">Go Back</a>
        </div>`+ templates.dms;
        await loadUsers();
        history.pushState({}, "", "/profile");
    } catch (err) {
        console.error("Failed to load profile:", err);
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
            headers: {
                'X-Requested-With': 'XMLHttpRequest',
                "Content-Type": "application/json",
            }
        });
        if (!response.ok) {
            handleApiError(response.status, "Authentication failed");
            return;
        }
        const authData = await response.json();
        const authenticated = authData.authenticated;
        if (authenticated) {
            await setHeader(authenticated);
            loadPageFromPath();
        } else {
            app.innerHTML = templates.auth;
            bindInputTrimming();
        }
    } catch (error) {
        console.error("Authentication check failed:", error);
        showErrorPage(500, "Unable to verify authentication");
    }
});

async function loadUsers() {
    const container = document.querySelector(".chat-users");
    if (!container) {
        console.warn("chat-users container not found in the DOM.");
        return;
    }


    try {
        const response = await fetch("/api/v1/get/users", {
            method: "GET",
            headers: {
                "X-Requested-With": "XMLHttpRequest",
                "Content-Type": "application/json",
            },
        });

        if (!response.ok) {
            showErrorPage(response.status, "Unable to load users");
            return;
        }

        const users = await response.json();
        if (!users) return;
        
        users.forEach((user) => {
            const userDiv = document.createElement("div");
            userDiv.className = "chat-user";
            userDiv.id = "message-select";
            userDiv.textContent = `${user.online ? "🟢" : "🔴"} ${user.username}`;
            userDiv.setAttribute("username", user.username);
            userDiv.onclick = () => setupChat(user.username);
            container.appendChild(userDiv);
        });

    } catch (error) {
        console.error("Error loading users:", error);
        showErrorPage(500, "Network error or unable to load users");
    }
}

function setupChat(username) {
    document.getElementById("chat-username").textContent = username;
    const userList = document.querySelector(".chat-users");
    const chatBox = document.getElementById("chat-box");
    const inputGroup = document.querySelector(".input-group");
    const userSearch = document.getElementById("userSearch");
    userList.addEventListener("click", (e) => {
        
        // Display chat box and input
        userSearch.style.display = "none";
        chatBox.style.display = "flex";
        chatBox.title = username;
        inputGroup.style.display = "flex";
        userList.style.display = "none";
        document.getElementById("discussion").style.height ="100%";
        // Load chat history
        changeDiscussion(username);
    });
}


async function setHeader(authStatus) {
    if (document.getElementById("main-header")) return;
    const wrapper = document.createElement("div");
    wrapper.innerHTML = templates.header;

    const headerNode = wrapper.firstElementChild;
    headerNode.id = "main-header";
    document.body.insertBefore(headerNode, document.body.firstChild);
    updateNavbar(authStatus);

    injectStylesheet("/front-end/styles/header.css");
    injectStylesheet("/front-end/styles/dms.css");
    injectStylesheet("/front-end/styles/style.css");
}

async function setupPostCreator() {
    const postSection = document.getElementById("create-post-section");
    const closeBtn = document.getElementById("close-post-creator");
    const form = document.getElementById("post-form");

    // Close modal
    closeBtn?.addEventListener("click", () => {
        document.body.classList.remove("dimmed");
        postSection.style.display = "none";
        saveDraft();
    });

    // Save draft to localStorage
    function saveDraft() {
        const draft = {
            title: document.getElementById("post-title").value,
            content: document.getElementById("post-content").value,
            category: document.getElementById("post-category").value,
        };
        localStorage.setItem("postDraft", JSON.stringify(draft));
    }

    // Load draft if exists
    const draft = JSON.parse(localStorage.getItem("postDraft"));
    if (draft) {
        document.getElementById("post-title").value = draft.title;
        document.getElementById("post-content").value = draft.content;
        document.getElementById("post-category").value = draft.category;
    }

    // Clear draft on submit
    form?.addEventListener("submit", async (e) => {
        e.preventDefault();
        await submitPost(e);
        await loadPosts({ mode: "prepend" })
        localStorage.removeItem("postDraft");
        form.reset();
        document.body.classList.remove("dimmed");
        postSection.style.display = "none";
    });
}

async function submitPost() {
    const payload = {
        title: document.getElementById("post-title").value.trim(),
        content: escapeHTML(document.getElementById("post-content").value.trim()),
        category: document.getElementById("post-category").value.trim(),
    };

    if (payload.title.length < 3 || payload.title.length > 30) {
        return showNotification("Title length must be between 3 and 30 characters", "error");
    }
    if (payload.category.length > 30) {
        return showNotification("Category length exceeded the allowed limit", "error");
    }
    if (payload.content.length < 10 || payload.content.length > 1500) {
        return showNotification("Content length must be between 10 and 1000 characters", "error");
    }

    try {
        const res = await fetch("/api/v1/post/createPost", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "X-Requested-With": "XMLHttpRequest",
            },
            body: JSON.stringify(payload),
        });

        if (!res.ok) throw new Error("Submission failed");

        showNotification("Post submitted successfully", "success");

        // Optionally clear inputs after submission
        document.getElementById("post-form").reset();

    } catch (err) {
        console.error("Error submitting post:", err);
    }

}

document.addEventListener("click", async (event) => {
    if (event.target.id === "submit-comment") {
        event.preventDefault();
        await submitComment(event);
    }
});

async function submitComment(event) {
    const container = event.target.closest(".comment-container");
    const textarea = container.querySelector(".comment-textarea");
    const comment = textarea.value.trim();
    if (!comment) return showNotification("Comment cannot be empty.", "error");

    const postDiv = container.closest(".post");
    const postId = postDiv?.id;

    if (!postId) return showNotification("Post ID not found.", "error");
    if (String(comment).length > 1000) return showNotification("Comment length exceeded the allowed limit","error")
    const payload = { content: comment, post_id: postId };
    try {
        const res = await fetch("/api/v1/post/createComment", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "X-Requested-With": "XMLHttpRequest"
            },
            body: JSON.stringify(payload),
        });

        if (!res.ok) throw new Error("Failed to post comment");

        const result = await res.json();

        textarea.value = "";
        const newCommentHTML = commentTemplate({
            content: comment,
            author: document.getElementById("username").textContent,
            creation_time: new Date().toISOString()
        });
        postDiv.querySelector(".comments")?.insertAdjacentHTML("afterbegin", newCommentHTML);

        showNotification("Comment posted!", "success");
    } catch (err) {
        console.error("Comment error:", err);
        showNotification("Error posting comment", "error");
    }
}

console.log("Loaded inedex.js")