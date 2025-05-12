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

async function viewComments(event) {
    const button = event?.target;
    const postDiv = button.closest(".post");

    if (!postDiv) return console.warn("No post container found.");

    const pid = postDiv.id;
    const container = postDiv.querySelector(".comment-container");

    const currentCount = container?.querySelectorAll(".comment")?.length || 0;

    await loadComments({
        // pid,
        // offset: currentCount,
        mode: "append",
        target: button
    });
}

console.log("Loaded convenience.js")