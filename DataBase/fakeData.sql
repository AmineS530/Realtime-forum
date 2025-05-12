PRAGMA foreign_keys = ON;

-- USERS
INSERT INTO users (id, username, first_name, last_name, email, age, gender)
VALUES
(1, 'alice123', 'Alice', 'Smith', 'alice@example.com', 25, 'female'),
(2, 'bobBuilder', 'Bob', 'Builder', 'bob@example.com', 34, 'male'),
(3, 'heliPilot', 'Apache', 'Rotor', 'apache@flightmail.com', 41, 'Attack helicopter');

-- CREDENTIALS
INSERT INTO credentials (id, hash)
VALUES
(1, X'1234abcd'),
(2, X'abcd1234'),
(3, X'deadbeef');

-- SESSIONS
INSERT INTO sessions (user_id, session_id, expires_at)
VALUES
(1, 'sess-alice-xyz', DATETIME('now', '+1 day')),
(2, 'sess-bob-xyz', DATETIME('now', '+1 day')),
(3, 'sess-apache-xyz', DATETIME('now', '+1 day'));

-- POSTS
INSERT INTO posts (post_id, uid, title, content, categories)
VALUES
(1, 1, 'Hello Forum!', 'Glad to be here. Excited to connect!', 'introduction'),
(2, 2, 'DIY Shed Building', 'Today I built a shed from scratch using pallet wood!', 'DIY, construction'),
(3, 3, 'Rotor Tricks', 'Here’s how I hover perfectly at low altitude.', 'aviation'),
(4, 1, 'Forum Tips', 'Use the search feature before posting!', 'tips, general');

-- COMMENTS
INSERT INTO comments (comment_id, post_id, uid, content)
VALUES
(1, 1, 2, 'Welcome, Alice!'),
(2, 2, 1, 'That’s impressive, Bob!'),
(3, 3, 2, 'How do you manage stability mid-air?'),
(4, 4, 3, 'I’ll pretend I have a keyboard to reply to this.'),
(5, 1, 3, 'Rotor engaged. Comment deployed.');

-- DMS
INSERT INTO dms (message_id, sender_id, recipient_id, message)
VALUES
(1, 1, 2, 'Hi Bob! Great shed build. Want to collaborate?'),
(2, 2, 1, 'Thanks Alice! Sure, let’s plan something.'),
(3, 3, 1, 'Alice, I have rotor data you might enjoy.'),
(4, 2, 3, 'Do you ever land?'),
(5, 3, 2, 'No. I hover eternally.');
