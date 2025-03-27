// document.addEventListener("DOMContentLoaded", async () => {
//     const app = document.getElementById("app");

//     try {
//         const response = await fetch("/api/check-auth", { credentials: "include" });
//         const authData = await response.json();

//         if (authData.authenticated) {
//             // ✅ User is logged in → Load Forum UI
//             app.innerHTML = `
//                 <ul class="nav">
//                     <li>
//                         <a href="#" onclick="dms_ToggleShowSidebar(event)" title="Messages">
//                             📩 Messages
//                         </a>
//                     </li>
//                     <li>
//                         <a id="logout" href="#" onclick="logoutUser()" title="Logout">🚪 Logout</a>
//                     </li>
//                 </ul>
//                 <div id="forum-container">
//                     <button onclick="viewPosts(event)" class="view-comments" id="postSeeMore">See More</button>
//                 </div>
//                 <div id="comments-container"></div>
//                 <div id="private-messages-container"></div>
//             `;

//             loadComments();
//             loadPrivateMessages();
//         } else {
//             // ❌ Not logged in → Show login form
//             const authResponse = await fetch("/front-end/templates/auth.html");
//             app.innerHTML = await authResponse.text();
//         }
//     } catch (error) {
//         console.error("Error loading content:", error);
//         app.innerHTML = "<h2>Error loading the page. Please try again.</h2>";
//     }
// });

// async function logoutUser() {
//     await fetch("/api/logout", { method: "POST", credentials: "include" });
//     window.location.reload();
// }

// function loadComments() {
//     document.getElementById("comments-container").innerHTML = "Loading comments...";
//     fetch("/api/comments")
//         .then(res => res.text())
//         .then(html => {
//             document.getElementById("comments-container").innerHTML = html;
//         });
// }

// function loadPrivateMessages() {
//     document.getElementById("private-messages-container").innerHTML = "Loading messages...";
//     fetch("/api/messages")
//         .then(res => res.text())
//         .then(html => {
//             document.getElementById("private-messages-container").innerHTML = html;
//         });
// }
