let commentLimit = undefined;

window.loadComments = async function loadComments(event, offset) {
    let parent = event.target.parentElement;
    let comments = [];
    try {
        let response = await fetch(
            `/api/v1/get/comments?pid=${parent.id}${commentLimit ? `&limit=${commentLimit}` : ""
            }${offset ? `&offset=${offset}` : ""}`,
            {
                method: "GET",
                headers: {
                    "X-Requested-With": "XMLHttpRequest",
                    "Content-Type": "application/json",
                },
                credentials: "include"
            }
        );
        // console.log(response);
        if (!response.ok) {
            throw new Error(`Error: ${response.status} - ${response.statusText}`);
        }
        comments = await response.json();
        } catch (error) {
        return;
    }
    // console.log("azer", comments);
    let commentall = "";
    console.log("comments", comments);
    if (comments != null) {
        for (const comment of comments) {
            commentall += commentTemplate(comment);
        }
    }
    if (offset) {
        event.target.setAttribute(
            "onclick",
            `viewComments(event, ${offset + comments.length})`
        );
        event.target.insertAdjacentHTML("beforebegin", commentall);
    } else {
        let comment_area = `
    <details class="comment-container" open>
    <summary>Click to see comments</summary> 
    <div class="comment-input-area">
        <textarea 
            class="comment-textarea" 
            placeholder="Write your comment here..."
            rows="3"
        ></textarea>
        <button id="submit-comment" type="submit">Post</button>
    </div>`
        if (comments != null) {
            let allcomments = `  
        ${commentall}
        <button class="view-comments" onclick="viewComments(event, ${comments.length})">Load more comments</button>
        </details>
        `
            event.target.outerHTML = comment_area + allcomments;
        } else {
            event.target.outerHTML = comment_area;
        }
    }
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

    <button onclick="viewComments(event)" class="view-comments" data-post-id="${post.pid}">
      View Comments
    </button>
  </div>
`;


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
