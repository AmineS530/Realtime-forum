PRAGMA foreign_keys = ON;
CREATE TABLE
    IF NOT EXISTS `users` (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE NOT NULL,
        first_name TEXT NOT NULL,
        last_name TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL,
        age INTEGER NOT NULL,
        gender CHAR NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CHECK (gender IN ("male", "female", "DFK"))
    );

CREATE TABLE
    IF NOT EXISTS `cerdentials` (
        id INTEGER PRIMARY KEY NOT NULL,
        hash BLOB NOT NULL,
        FOREIGN KEY (id) REFERENCES users (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
    );

CREATE TABLE
    IF NOT EXISTS `posts` (
        posts_id INTEGER PRIMARY KEY AUTOINCREMENT,
        uid INTEGER NOT NULL,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        categories TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (uid) REFERENCES users (id)
    );

CREATE TABLE
    IF NOT EXISTS `comments` (
        comments_id INTEGER PRIMARY KEY AUTOINCREMENT,
        posts_id INTEGER NOT NULL,
        uid INTEGER NOT NULL,
        content TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (posts_id) REFERENCES posts (posts_id),
        FOREIGN KEY (uid) REFERENCES users (id),
    );

CREATE TABLE
    IF NOT EXISTS `dms` (
        message_id INTEGER PRIMARY KEY AUTOINCREMENT,
        sender_id INTEGER NOT NULL,
        recipient_id INTEGER NOT NULL,
        message TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (sender_id) REFERENCES users (id),
        FOREIGN KEY (recipient_id) REFERENCES users (id),
        CHECK (sender_id <> recipient_id)
    );

-- selects
WITH
    vars AS (
        SELECT
            ? AS name1,
            ? AS name2,
            ? AS offset
    )
SELECT
    d.message,
    d.created_at,
    u1.username,
    u2.username
FROM
    dms d
    JOIN users u1 ON d.sender_id = u1.id
    JOIN users u2 ON d.recipient_id = u2.id
WHERE
    u1.username = vars.name1
    AND u2.username = vars.name2
    OR u1.username = vars.name2
    AND u2.username = vars.name1
ORDER BY
    d.created_at DESC
    LIMIT vars.offset, 10;


-- select latest posts
SELECT
    p.id,
    p.title,
    p.content,
    p.categories,
    p.creation_date,
    u.username
    FROM posts p
    JOIN users u ON p.uid = u.id
    LEFT JOIN posts_categories pc ON p.id = pc.post_id
    LEFT JOIN categories c ON pc.category_id = c.id
    GROUP BY p.id
    ORDER BY p.creation_date DESC
    LIMIT ?,10;
-- select comments by post id
SELECT
    c.id,
    c.comment_text,
    c.comment_date,
    u.username
FROM comments c
JOIN users u ON c.uid = u.id
WHERE c.post_id = ?
ORDER BY c.comment_date DESC
LIMIT ?,10;

-- select  hash
SELECT c.hash
FROM users u
JOIN cerdentials c ON u.id = c.uid
WHERE u.username = ? OR u.email = ?;

-- insert new post and it's categories
INSERT INTO posts (title, content, categories, id) VALUES (?, ?, ?);

-- insert new comment
INSERT INTO comments (comment_text, id, post_id) VALUES (?, ?, ?);

-- insert new dm
INSERT INTO dms (sender_id, recipient_id, message) VALUES (?, ?, ?);

-- insert new user
INSERT INTO users (username, email, first_name, last_name, age, gender) VALUES (?, ?, ?, ?, ?, ? );

-- insert new credentials
INSERT INTO credentials (uid, hash) VALUES ((SELECT TOP 1 id FROM users WHERE username = ? AND email= ?),? );