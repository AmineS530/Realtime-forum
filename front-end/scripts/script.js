console.log("hello world");

window.addEventListener("hashchange", () => {
    switch (window.location.hash) {
        case "#home":
            // TODO : home
            console.log("Navigating to the home section!");
            // Load or display home content dynamically.
            break;
        case "#about":
            // TODO : about
            console.log("Navigating to the about section!");
            // Load or display about content dynamically.
            break;
        case "#login":
            showLoginForm();
            break;
        case "#registration":
            showSignUpForm();
            break;
        default:
            console.log("Unknown hash:", window.location.hash);
            // Handle any other hash cases or fallback logic.
            break;
    }
});

function fetching(e, target) {
    e.preventDefault();
    let data = {};
    for (const element of e.target) {
        if (element.value != "") {
            console.log(element.value, element.id, element);
            data[element.id] = element.value;
        }
    }
    console.log(e.target, e.this, data, target);
    let a = fetch(target, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    })
        .then((response) => {
            if (!response.ok) {
                throw new Error(
                    " Network error \n\tstatus code :" + response.status
                );
            }
            return response.json();
        })
        .catch((error) => {
            console.log("Error:", error);
            // Return error object with a message
            return { error: error.message };
        });
    a.then((result) => {
        console.log("Response as object:", result);
        for (let key in result) {
            if (result.hasOwnProperty(key)) {
                alert(`${key}: ${result[key]}`);
            }
        }
    });
}
