import svg from "./svg.js";

const header = `
    <header>
      <!-- Logo -->
      <div class="logo">
        <h1><a href="/" onclick="loadPage('home', event)">${WebAppName}</a></h1>
      </div>
      <!-- Navigation -->
      <nav id="nav">
        <ul class="nav"></ul>
      </nav>
    </header>
    <br />
`;

const auth = `
<link rel="stylesheet" href="/front-end/styles/style.css" />
<link rel="stylesheet" href="/front-end/styles/log-reg.css" />
<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined" /> 
<div id="auth">
    <div class="container">
        <!-- Buttons -->
        <div class="btn">
            <button id="login-btn" onclick="showLoginForm()">Login</button>
            <button id="register-btn" onclick="showSignUpForm()">Sign Up</button>
        </div>

        <!-- Slider -->
        <div class="slider"></div>
        <!-- Form Section -->
        <div class="form-section">
            <!-- Login Form -->
            <div class="login-box">
                <form onsubmit="fetching(event,'/api/login')" method="post">
                    <label for="name_or_email">Username or Email</label>
                    <input type="text" id="name_or_email" name="name_or_email" maxlength="50" required/>
                    <label for="password">Password</label>
                    <div class="input-wrapper">
                    <input type="password" id="logpassword" name="password" maxlength="40" required />
                    <i class="togglePwd" > <span class="icon material-symbols-outlined">visibility</span></i>
                    </div>
                    <button type="submit">Login</button>
                    <p>
                        Don't have an account?
                        <a style="cursor: pointer;" onclick="showSignUpForm()" >Sign Up</a>
                    </p>
                </form>
            </div>

            <!-- Sign Up Form -->
            <div class="register-box">
                <form onsubmit="fetching(event,'/api/register')" method="post">
                    <label for="username">Username</label>
                    <input type="text" id="username" name="Username" minlength="3" maxlength="18" required />
                    <label for="email">Email</label>
                    <input type="email" id="email" name="Email" maxlength="320" required />
                    <label for="password">Password</label>
                    <div class="input-wrapper">
                        <input type="password" id="regpassword" name="Password" maxlength="30" required />
                        <i class="togglePwd" > <span class="icon material-symbols-outlined">visibility</span></i>
                    </div>
                    <label for="password_confirmation">Confirm Password</label>
                    <div class="input-wrapper">
                        <input type="password" id="password_confirmation"name="Password_confirmation" maxlength="30" required />
                        <i class="togglePwd" > <span class="icon material-symbols-outlined">visibility</span></i>
                    </div>
                    <label for="age">Age</label>
                    <input type="number" id="age" name="Age" min="15" max="90" required />
                    <label for="gender">Gender </label>
                    <select id="gender" name="Gender" required>
                        <option value="">Select Gender</option>
                        <option value="male">Male</option>
                        <option value="female">Female</option>
                        <option value="Attack helicopter">Attack helicopter</option>
                    </select>
                    <label for="first-name">First Name</label>
                    <input type="text" id="first-name" name="First_Name" required />
                    <label for="last-name">Last Name</label>
                    <input type="text" id="last-name" name="Last_Name" required />
                    <button type="submit">Sign Up</button>
                    <p>
                        Already have an account? <a style="cursor: pointer;" onclick="showLoginForm()" >Login</a>
                    </p>
                </form>
            </div>
        </div>
    </div>
</div>
`;

const posts = `
    <link rel="stylesheet" href="/front-end/styles/header.css" />
    <link rel="stylesheet" href="/front-end/styles/style.css" />
    <button id="load-posts" onclick="viewPosts(event)">Load More Posts</button>
`;

const dms = `
<div id="backdrop" onclick="event.target.id === 'backdrop' ? dms_ToggleShowSidebar(event) : event.stopPropagation();">
    <div class="show" id="side-menu" aria-modal="true" role="dialog">
        <div class="side-menu-head">
            <h1 class="offcanvas-title" style="margin-bottom: 0; line-height: 5vh;">Messages</h1>
            <button onclick="dms_ToggleShowSidebar(event)" class="btn">
                ${svg.two_bubbles}
            </button>
        </div>

        <!-- Search input -->
        <input
            type="text"
            id="userSearch"
            placeholder="Search users..."
            class="search-bar"
            onkeyup="searchUser()"
        />

        <div class="chat-users"></div>

        <!-- Chat messages -->
          <div id="chat-box" class="chat-box" style="display: none">
          <div id="chat-header">
          <button class="close-chat" onclick="closeChat()">✖</button>
            <span id="chat-username">Chat</span>
            <div id="typing" class="hidden">
                <span></span>
                <span></span>
                <span></span>
            </div>
          </div>
          </div>
         <ul id="discussion" class="discussion"></ul>
         
        <!-- Message input form -->
        <form class="input-group" onsubmit="event.preventDefault();sendDm(event)" style="display: none">
            <input
                type="text"
                oninput="startTyping()"
                class="form-control"
                minlength="1"
                maxlength="250"
                placeholder="New Message..."
                aria-label="Message"
            />
            <button class="btn btn-primary" type="submit">
                ${svg.svg_send}
            </button>
        </form>
    </div>
</div>
`;

const postCreation = `
    <center>
        <div id="create-post-section" style="display: none">
            <h2>Create a new Post</h2>
            <br />
            <button id="close-post-creator">&times;</button>
            <form id="post-form">
                <input type="text" id="post-title" placeholder="Post Title" minlength="3" maxlength="30" required />
                <textarea id="post-content" placeholder="Write your post..." minlength="10" maxlength="1500" required></textarea>
                <input id="post-category" placeholder="Enter categories... (max 3)" minlength="3" maxlength="30" />
                <button type="submit">Submit Post</button>
            </form>
        </div>
    </center>
`;

export default { header, posts, auth, dms, postCreation };
