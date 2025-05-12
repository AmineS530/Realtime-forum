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


function injectStylesheet(href) {
    if (!document.querySelector(`link[href="${href}"]`)) {
        const link = document.createElement("link");
        link.rel = "stylesheet";
        link.href = href;
        document.head.appendChild(link);
    }
}

function delay(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function showNotification(message, type = "OK") {
    const notification = document.createElement("div");
    notification.innerText = message;
    notification.className = `notification ${type}`;
    document.body.appendChild(notification);
    setTimeout(() => notification.remove(), 3000);
}


async function viewPosts(event) {
    await loadPosts({
        offset: document.querySelectorAll("#app .post").length,
        mode: "append",
        target: event.target
    });
}


const commentTemplate = (comment) => `<div class="comment" id="post_id-${comment.pid}">
  <div class="comment-header">
    <span class="comment-author">Published by <strong>${comment.author}</strong></span>
    <span class="comment-date">${new Date(comment.creation_time).toLocaleString()}</span>
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

console.log("Loaded convenience.js")