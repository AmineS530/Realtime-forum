let commentLimit = undefined;

function fetching(e, target) {
    e.preventDefault();
    let data = {};
    for (const element of e.target) {
        if (element.value != "") {
            //(element.value, element.id, element);
            data[element.id] = element.value;
        }
    }
    //console.log(e.target, e.this, data, target);
    let a = fetch(target, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    })
        .then((response) => {
            if (!response.ok) {
                throw new Error(" Network error \n\tstatus code :" + response.status);
            }
            return response.json();
        })
        .catch((error) => {
            //console.log("Error:", error);
            // Return error object with a message
            return { error: error.message };
        });
    a.then((result) => {
        //console.log("Response as object:", result);
        for (let key in result) {
            if (result.hasOwnProperty(key)) {
                alert(`${key}: ${result[key]}`);
            }
        }
    });
}

window.viewComments = async function viewComments(event, offset) {
    let parent = event.target.parentElement;
    // console.log(event, parent);
    let comments = [];
    try {
        let response = await fetch(
            `/api/v1/get/comments?pid=${parent.id}${commentLimit ? `&limit=${commentLimit}` : ""
            }${offset ? `&offset=${offset}` : ""}`
        );
        // console.log(response);
        if (!response.ok) {
            throw new Error(`Error: ${response.status} - ${response.statusText}`);
        }
        comments = await response.json();
        // console.log(comments);
    } catch (error) {
        // console.error("Error:", error);
        return;
    }
    // console.log("azer", comments);
    let commentall = "";
    for (const comment of comments) {
        commentall += `<div class="comment" id="${comment.id
            }" style="borderwidth: 5px;">
        <span class="comment-info">
            <span class="comment-author">published by ${comment.author}</span>
            <span class="comment-date">published ${new Date(comment.creation_time)}</span>
        </span>
        <p>${comment.content}<p>
        </div>`;
    }
    if (offset) {
        event.target.setAttribute(
            "onclick",
            `viewComments(event, ${offset + comments.length})`
        );
        event.target.insertAdjacentHTML("beforebegin", commentall);
    } else {
        let azer = `<details class="comment-container" open>
        <summary>Click to see comments</summary> ${commentall}<button onclick="viewComments(event, ${comments.length})">load more comments</button>
        </details>`;
        event.target.outerHTML = azer;
    }
};

window.viewPosts = async function viewPosts(event, offset = 0) {
    try {
        const response = await fetch(`/api/v1/get/posts?offset=${offset}`);
        if (!response.ok)
            throw new Error(`${response.status} ${response.statusText}`);
        const posts = await response.json();

        let html = "";
        for (const post of posts) {
            html += `
            <div class="post" id="post${post.pid}">
                <h3>${post.title}</h3>
                <p>${post.content}</p>
                <span class="post-info">
                    <span class="post-author">Posted by ${post.author}</span>
                    <span class="post-date">${new Date(
                post.creation_time
            ).toLocaleString()}</span>
                    <span class="post-category">${post.categories.join(
                " | "
            )}</span>
                </span>
                <button onclick="viewComments(event)" class="view-comments" data-post-id="${post.pid
                }">View Comments</button>
            </div>`;
        }

        event.target.insertAdjacentHTML("beforebegin", html);
    } catch (err) {
        console.error("Failed to load posts:", err);
    }
};
/*
document.addEventListener("DOMContentLoaded", () => {
    let qsdf = document.getElementById("postSeeMore");
    if (qsdf) {
        qsdf.click();
        setTimeout(() => {
            document.getElementById("post_id-1").lastElementChild.click();
            // console.log("5 seconds have passed!");
        }, 500);
    }
});*/

const postTemplate = (post) => `
<div class="post" id="post${post.pid}">
    <h3>${post.title}</h3>
    <p>${post.content}</p>
    <span class="post-info">
        <span class="post-author">Posted by ${post.author}</span>
        <span class="post-date">${new Date(
    post.creation_time
).toLocaleString()}</span>
        <span class="post-category">${post.categories.join(" | ")}</span>
    </span>
    <button onclick="viewComments(event)" class="view-comments" data-post-id="${post.pid
    }">View Comments</button>
</div>
`;
