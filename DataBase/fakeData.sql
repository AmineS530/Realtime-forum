-- Inserting fake users
INSERT INTO users (username, first_name, last_name, email, age, gender) VALUES
('john_doe', 'John', 'Doe', 'john.doe@example.com', 25, 'male'),
('jane_smith', 'Jane', 'Smith', 'jane.smith@example.com', 30, 'female'),
('helicopter_guy', 'Chris', 'Jetson', 'chris.jetson@example.com', 28, 'Attack helicopter');

-- Inserting fake sessions
INSERT INTO sessions (user_id, session_id, expires_at) VALUES
(1, 'session_1', '2025-12-31 23:59:59'),
(2, 'session_2', '2025-12-31 23:59:59'),
(3, 'session_3', '2025-12-31 23:59:59');

-- Inserting fake credentials (hashes are represented as fake blobs for simplicity)
INSERT INTO credentials (id, hash) VALUES
(1, X'1234567890abcdef'),
(2, X'abcdef1234567890'),
(3, X'0987654321abcdef');

-- Inserting fake posts
INSERT INTO posts (uid, title, content) VALUES
(1, 'Post 1', 'This is the content of the first post.'),
(2, 'Post 2', 'This is the content of the second post.'),
(3, 'Post 3', 'This is the content of the third post.');

-- Inserting fake categories
INSERT INTO categories (name) VALUES
('Technology'),
('Lifestyle'),
('Gaming'),
('Health');

-- Inserting fake postCategory relationships (assigning categories to posts)
INSERT INTO postCategory (post_id, category_id) VALUES
(1, 1), -- Post 1 is in the Technology category
(2, 2), -- Post 2 is in the Lifestyle category
(3, 3), -- Post 3 is in the Gaming category
(1, 4); -- Post 1 is also in the Health category

-- Inserting fake comments
INSERT INTO comments (post_id, uid, content) VALUES
(1, 2, 'This is a comment on Post 1.'),
(2, 3, 'This is a comment on Post 2.'),
(3, 1, 'This is a comment on Post 3.');

-- Inserting fake direct messages (DMs)
INSERT INTO dms (sender_id, recipient_id, message) VALUES
(1, 2, 'Hey, how are you?'),
(2, 3, 'I have a question about your post.'),
(3, 1, 'Thanks for your feedback!');
