let commentLimit = undefined;

window.loadComments = async function loadComments(input, offset = 0) {
    const isEvent = !(input?.mode);
    const event = isEvent ? input : null;
    const mode = isEvent ? "default" : input.mode;
    const pid = isEvent ? event?.target?.parentElement?.id : input.pid;
    const parent = document.getElementById(pid);

    if (!pid || !parent) return console.warn("Missing post ID or container");

    let comments = [];
    try {
        const res = await fetch(`/api/v1/get/comments?pid=${pid}${commentLimit ? `&limit=${commentLimit}` : ""}${offset ? `&offset=${offset}` : ""}`, {
            method: "GET",
            headers: {
                "X-Requested-With": "XMLHttpRequest",
                "Content-Type": "application/json"
            },
            credentials: "include"
        });
        if (!res.ok) throw new Error();
        comments = await res.json();
    } catch {
        return;
    }

    const commentHTML = (Array.isArray(comments) ? comments : []).map(commentTemplate).join("");

    if (offset) {
        event.target.setAttribute("onclick", `loadComments(event, ${offset + comments.length})`);
        event.target.insertAdjacentHTML("beforebegin", commentHTML);
        return;
    }

    const comment_area = `
<details class="comment-container" open>
    <summary>Click to see comments</summary>
    <div class="comment-input-area">
        <textarea class="comment-textarea" placeholder="Write your comment here..." rows="3"></textarea>
        <button id="submit-comment" type="submit">Post</button>
    </div>
    <div class="comments">${commentHTML}</div>
</details>`;
    event.target.outerHTML = comment_area;
};

window.loadPosts = async function ({ offset = 0, mode = "replace", target = null } = {}) {
    try {
        const response = await fetch(`/api/v1/get/posts?offset=${offset}`, {
            method: "GET",
            headers: {
                "X-Requested-With": "XMLHttpRequest",
                "Content-Type": "application/json"
            },
            credentials: "include",
        });

        if (!response.ok)
            throw new Error(`${response.status} ${response.statusText}`);

        const posts = await response.json();
        if (!Array.isArray(posts) || posts.length === 0) return;

        const app = document.getElementById("app");
        let html = posts.map(postTemplate).join("");

        switch (mode) {
            case "replace":
                app.insertAdjacentHTML("afterbegin", html);
                break;
            case "append":
                if (target) {
                    target.insertAdjacentHTML("beforebegin", html);
                } else {
                    app.insertAdjacentHTML("beforeend", html);
                }
                break;
            case "prepend":
                app.insertAdjacentHTML("afterbegin", html);
                break;
        }
    } catch (err) {
        console.error("Failed to load posts:", err);
    }
};

const postTemplate = (post) => `
  <div class="post" id="${post.pid}">
    <h3 class="post-title">${escapeHTML(post.title)}</h3>
    <span class="post-category">
      Categor${post.categories.length > 1 ? "ies" : "y"}: ${post.categories.join(" | ")}
    </span>
    
    <p class="post-content">${escapeHTML(post.content)}</p>

    <div class="post-info">
      <span class="post-author">Posted by <strong>${post.author}</strong></span>
      <span class="post-date">${new Date(post.creation_time).toLocaleString()}</span>
    </div>

    <button onclick="loadComments(event)" class="view-comments" data-post-id="${post.pid}">
      View Comments
    </button>
  </div>
`;





