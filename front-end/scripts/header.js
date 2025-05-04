import svg from "./svg.js";
import templates from "./templates.js";

export async function updateNavbar(auth) {
    const navList = document.querySelector(".nav");
    if (!navList) return;

    let Username = "";
    try {
        const res = await fetch("/api/profile");
        if (!res.ok) throw new Error("Failed to fetch username");
        const data = await res.json();
        Username = data.username;
    } catch (err) {
        console.error("Could not fetch username:", err);
    }

    if (auth) {
        await delay(150);
        // Post creation icon
        const postCreation = document.createElement("li");
        postCreation.innerHTML = `<a class="logo" href="#" title="Create Post">${svg.svg_post_creation}</a>`;
        navList.appendChild(postCreation);
        postCreation.querySelector("a").addEventListener("click", (e) => {
            e.preventDefault();
            const postSection = document.getElementById("create-post-section");
            if (postSection) {
                postSection.style.display = "block";
                document.body.classList.add("dimmed");
            } else {
               showNotification("You need to be at home page to create a post", "error");
            }
        }); 

        // DMs icon
        const showbubbles = document.createElement("li");
        showbubbles.innerHTML = `<a class="logo" href="#" title="Messages">${svg.two_bubbles}</a>`;
        navList.appendChild(showbubbles);
        showbubbles.querySelector("a").addEventListener("click", (e) => {
            e.preventDefault();
            dms_ToggleShowSidebar(e);
        });

        // Create user dropdown
        const usernameItem = document.createElement("li");
        usernameItem.classList.add("dropdown");
        usernameItem.innerHTML = `
    <a class="logo" class="dropdown-button">${Username}</a>
    <div class="dropdown-content">
        <a id="profile" href="#" >Profile</a>
        <a href="#" id="logout">Log out</a>
    </div>`;
        navList.appendChild(usernameItem);
        const profileLink = usernameItem.querySelector("#profile");
        if (profileLink) {
            profileLink.addEventListener("click", (e) => {
                e.preventDefault();
                loadPage("profile");
            });
        }
        const logoutButton = document.getElementById("logout");
        if (logoutButton) {
            logoutButton.addEventListener("click", function (event) {
                event.preventDefault();
                fetch("/api/logout", {
                    method: "POST",
                    credentials: "include",
                })
                    .then(() => {
                        document.getElementById("app").innerHTML = "";
                        window.location.href = "/";
                    })
                    .catch((error) => console.error("Logout failed:", error));
            });
        }
    }
    // else {
    //     const loginItem = document.createElement("li");
    //     loginItem.innerHTML = `<a href="/login">Sign Up or Login</a>`;
    //     navList.appendChild(loginItem);
    // }
}
