const WebAppName = "EpicHub";
document.title = WebAppName;

function bindInputTrimming() {
    const inputs = document.querySelectorAll("input, textarea");
    inputs.forEach((input) => {
        if (input.type !== "password") {
            input.addEventListener("blur", function () {
                this.value = this.value.trim();
            });
        }
    });
}

const sounds = {
    message: new Audio("/front-end/sounds/messageNotif.mp3"),
    alert: new Audio("/front-end/sounds/alert.mp3"),
    notification: new Audio("/front-end/sounds/notification.mp3")
};

function playSound(name) {
    const sound = sounds[name];
    sound.volume = 0.75
    if (sound) {
        sound.currentTime = 0;
        sound.play().catch(e => console.warn("Playback failed:", e));
    }
}
// Optimized stylesheet injection with cache busting
function injectStylesheet(href) {
    if (!document.querySelector(`link[href="${href}"]`)) {
        const link = document.createElement("link");
        link.rel = "stylesheet";
        link.href = href;
        document.head.appendChild(link);
    }
}

function delay(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
}

let notificationCooldown = false;

function showNotification(message, type = "success", sound = true) {
    if (notificationCooldown) return;

    notificationCooldown = true;

    const notification = document.createElement("div");
    notification.className = `notification ${type}`;
    notification.textContent = message;

    document.body.appendChild(notification);

    // Slide down
    requestAnimationFrame(() => {
        notification.style.top = "40px";
    });

    // Play sound
    if (sound) {
        if (type === "success" || type === "info") {
            playSound("notification");
        } else {
            playSound("alert");
        }
    }

    setTimeout(() => {
        notification.style.top = "-100px";
        notification.style.opacity = "0";

        notification.addEventListener("transitionend", () => {
            notification.remove();
        }, { once: true });

        // Reset cooldown after animation finishes
        setTimeout(() => {
            notificationCooldown = false;
        }, 500);
    }, 1500);
}

function searchUser() {
    const input = document.getElementById("userSearch");
    const filter = input.value.toLowerCase();
    const usersContainer = document.querySelector(".chat-users");
    const users = usersContainer.querySelectorAll(".chat-user");

    users.forEach(user => {
        const username = user.textContent.toLowerCase();
        user.style.display = username.includes(filter) ? "block" : "none";
    });
}

async function viewPosts(event) {
    await loadPosts({
        offset: document.querySelectorAll("#app .post").length,
        mode: "append",
        target: event.target,
    });
}

const commentTemplate = (comment) => `<div class="comment" id="post_id-${comment.pid
    }">
  <div class="comment-header">
    <span class="comment-author">Published by <strong>${comment.author
    }</strong></span>
    <span class="comment-date">${new Date(
        comment.creation_time
    ).toLocaleString()}</span>
  </div>
  <div class="comment-body">
    <p>${escapeHTML(comment.content)}</p>
  </div>
</div>
`;

function escapeHTML(str) {
    return String(str ?? "")
        .replace(/&/g, "&amp;")
        .replace(/</g, "&lt;")
        .replace(/>/g, "&gt;")
        .replace(/"/g, "&quot;")
        .replace(/'/g, "&#039;");
}
async function handleApiError(res) {
    const contentType = res.headers.get('content-type');

    if (contentType.includes('text/html')) {
        return;
    }

    // Handle JSON errors normally
    const error = await res.json().catch(() => ({}));
    showErrorPage(res.status, error.message || "API request failed");
}

function closeChat() {
    const chatBox = document.getElementById("chat-box");
    const inputGroup = document.querySelector(".input-group");
    const userList = document.querySelector(".chat-users");
    const discussion = document.getElementById("discussion");
    const userSearch = document.getElementById("userSearch");

    discussion.innerHTML = "";
    chatBox.style.display = "none";
    inputGroup.style.display = "none";
    userList.style.display = "flex";
    userSearch.style.display = "flex";
}

function showErrorPage(errorCode, errorMessage) {
    // Ensure DOM is ready before proceeding
    if (!document.body) {
        document.addEventListener('DOMContentLoaded', () => showErrorPage(errorCode, errorMessage));
        return;
    }

    // Inject error page styles
    injectStylesheet("/front-end/styles/errors.css");

    // Error type definitions
    const ERROR_TYPES = {
        404: {
            title: "Page Not Found",
            description: "The page you are looking for might have been removed or does not exist."
        },
        500: {
            title: "Internal Server Error",
            description: "Something went wrong on our end. Please try again later."
        },
        403: {
            title: "Access Denied",
            description: "You do not have permission to access this page."
        },
        default: {
            title: "Unexpected Error",
            description: "An unexpected error occurred. Please try again."
        }
    };

    // Get the appropriate container
    const targetContainer = document.getElementById('app') || document.body;

    // Clear previous content and render error template
    targetContainer.innerHTML = errorPageTemplate;

    // Get error type details
    const errorType = ERROR_TYPES[errorCode] || ERROR_TYPES.default;
    // Update DOM elements safely
    const setElementText = (id, text) => {
        const el = document.getElementById(id);
        if (el) el.textContent = text;
    };

    setElementText('error-page-title', `${errorType.title} (Error ${errorCode})`);
    setElementText('error-page-description', errorMessage || errorType.description);

    // Show additional error details if available
    const detailsEl = document.getElementById('error-page-details');
    const messageEl = document.getElementById('error-page-message');
    if (errorMessage && detailsEl && messageEl) {
        messageEl.textContent = errorMessage;
        detailsEl.style.display = 'block';
    }

    // Handle home button click
    const homeBtn = document.getElementById('error-page-home-btn');
    if (homeBtn) {
        homeBtn.addEventListener('click', (e) => {
            e.preventDefault();
                loadPage("home");
        });
    }
}
const errorPageTemplate = `
<link rel="stylesheet" href="/front-end/styles/errors.css" />
<div class="error-page-container">
    <div class="error-page-icon">⚠️</div>
    <h1 class="error-page-title" id="error-page-title">Page Not Found</h1>
    <p class="error-page-description" id="error-page-description">The page you are looking for might have been removed or does not exist.</p>
    <div class="error-page-actions">
        <button class="error-page-home-btn" id="error-page-home-btn">Go Home</button>
    </div>
    <div class="error-page-details" id="error-page-details" style="display: none;">
        <pre id="error-page-message" style="text-align: center"></pre>
    </div>
</div>
`;

console.log("Loaded convenience.js");