import templates from "./templates.js";
import { updateNavbar } from "./header.js";

window.loadPage = function (page) {
    const app = document.getElementById("app");
    switch (page) {
        case "home":
            setHeader(true);
            loadUsers();
            app.innerHTML = templates.posts + templates.dms + templates.postCreation;
            setupPostCreator();
            loadPosts({ mode: "replace" });
            history.pushState({}, "", "/");
            break;
        case "profile":
            loadProfilePage();
            loadUsers();
            break;
        default:
            app.innerHTML = "<h2>Page not foundo</h2>";
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
            app.innerHTML = "<h2>Page not foundp</h2>";
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
       <!-- <a class="change-password" href="#" > Change Password </a> 
        <br /> -->
        <a href="/" onclick="loadPage('home'); return false;">Go Back</a>
        </div>`+ templates.dms;
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
            headers: {
                'X-Requested-With': 'XMLHttpRequest',
                "Content-Type": "application/json",
            }
        });
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
        console.error("Error loading:", error);
        app.innerHTML = "<h2>Error loading the page.</h2>";
    }
});

async function loadUsers() {
    const discussion = document.getElementById("discussion");

    try {
        const response = await fetch("/api/v1/get/users", {
            method: "GET",
            headers: {
                'X-Requested-With': 'XMLHttpRequest',
                "Content-Type": "application/json",
            }
        });

        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const data = await response.json();
        console.log("User data:", data);

        let formattedHistory = "";
        data.forEach((user) => {
            formattedHistory += `<option>${user.online?'🟢':'🔴'} ${user.username}</option>`;
        });

        console.log("Formatted options:", formattedHistory);
        document.getElementById("message-select").innerHTML += formattedHistory;

    } catch (error) {
        console.error("Error loading users:", error);
    }
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

document.addEventListener("DOMContentLoaded", async () => {
    await setupPostCreator();
});

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
        showNotification("Post submitted successfully!", "success");
    });
}

async function submitPost() {
    const payload = {
        title: document.getElementById("post-title").value.trim(),
        content: escapeHTML(document.getElementById("post-content").value.trim()),
        category: document.getElementById("post-category").value.trim(),
    };

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

        console.log("Post submitted successfully");

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

    if (!postId) return alert("Post ID not found.");

    const payload = { content: comment, post_id: postId };
    console.log("Comment payload:", payload.content);
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
        console.log("Comment posted:", result);
        textarea.value = "";
        const newCommentHTML = commentTemplate({
            content: comment,
            author: result.username,
            creation_time: new Date().toISOString()
        });
        postDiv.querySelector(".comments")?.insertAdjacentHTML("afterbegin", newCommentHTML);

        showNotification("Comment posted!", "success");
    } catch (err) {
        console.error("Comment error:", err);
        showNotification("Error posting comment", "error");
    }
}

function escapeHTML(str) {
    return str.replace(/[&<>"']/g, (match) => ({
        '&': '&amp;',
        '<': '&lt;',
        '>': '&gt;',
        '"': '&quot;',
        "'": '&#039;',
    }[match]));
}

console.log("Loaded inedex.js")