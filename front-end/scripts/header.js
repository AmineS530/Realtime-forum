import  svg  from "./svg.js";

export async function updateNavbar(auth) {
    const navList = document.querySelector(".nav");
    if (!navList) return;

    const formattedUsername = "placeholder";

    if (auth) {
        const showbubbles = document.createElement("li");
        showbubbles.innerHTML = `<a class="logo" href="#"onclick="dms_ToggleShowSidebar(event)"  title="Messages">
                ${svg.two_bubbles}`;
        navList.appendChild(showbubbles);
        
        // Create user dropdown
        const usernameItem = document.createElement("li");
        usernameItem.classList.add("dropdown");
        usernameItem.innerHTML = `
    <a href="#" class="logo" class="dropdown-button">${formattedUsername}</a>
    <div class="dropdown-content">
        <a id="profile" href="#" onclick="loadPage('profile')">Profile</a>
        <a href="#" id="logout">Log out</a>
    </div>
`;

    
        navList.appendChild(usernameItem);

        const logoutButton = document.getElementById("logout");
        if (logoutButton) {
            logoutButton.addEventListener("click", function (event) {
                event.preventDefault();
                fetch("/api/logout", {
                    method: "POST",
                    credentials: "include"
                })
                    .then(() => {
                        window.location.href = "/";
                    })
                    .catch((error) => console.error("Logout failed:", error));
            });
        }

    } else {
        const loginItem = document.createElement("li");
        loginItem.innerHTML = `<a href="/login">Sign Up or Login</a>`;
        navList.appendChild(loginItem);
    }
}