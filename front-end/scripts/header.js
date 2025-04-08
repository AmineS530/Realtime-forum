export function updateNavbar(auth) {
    const navList = document.querySelector(".nav");
    if (!navList) return;
    const formattedUsername = "admin";

    if (auth) {
        // Create "Create Post" item
        const createpost = document.createElement("li");
        createpost.innerHTML = `<a href="/post-creation"</a>`;
        navList.appendChild(createpost);

        // Create user dropdown
        const usernameItem = document.createElement("li");
        usernameItem.classList.add("dropdown");
        usernameItem.innerHTML = `
                <a href="#" class="dropdown-button">${formattedUsername}</a>
                <div class="dropdown-content">
                    <a href="/profile">Profile</a>
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
