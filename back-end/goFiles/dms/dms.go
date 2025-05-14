package dms

import (
	"fmt"
	"strconv"
	"time"

	helpers "RTF/back-end"
)

type Message struct {
	Sender  string    `json:"sender"`
	Content string    `json:"message"`
	Time    time.Time `json:"time"`
}

func GetdmHistory(uname1, uname2, page string) ([]Message, error) {
	if page == "" {
		page = "0"
	}
	p, err := strconv.Atoi(page)
	if err != nil {
		helpers.ErrorLog.Println("Error converting pid to int: ", err)
		return nil, err
	}
	if p < 1 {
		p = 0
	} else {
		p *= 10
	}
	fmt.Println(uname1, uname2, page, p)
	rows, err := helpers.DataBase.Query(`
	SELECT * 
FROM (
    SELECT
		sender.username , d.message, d.created_at
	FROM
		dms d
	JOIN
		users sender ON d.sender_id = sender.id
	JOIN
		users recipient ON d.recipient_id = recipient.id
	WHERE
		(sender.username = ? AND recipient.username = ?)
   	OR
		(sender.username = ? AND recipient.username = ?)
	ORDER BY
		d.created_at DESC
	LIMIT 10
	OFFSET ?
) AS sub ORDER BY created_at ASC;
	`, uname1, uname2, uname2, uname1, p)
	if err != nil {
		fmt.Println("Error getting posts: ", err)
		return nil, err
	}
	defer rows.Close()
	var messages []Message
	for rows.Next() {
		var message Message
		err := rows.Scan(&message.Sender, &message.Content, &message.Time)
		if err != nil {
			fmt.Println("Error scanning posts: ", err)
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func AddDm(sUname, rUname, msg string) error {
	query := `INSERT INTO dms (sender_id, recipient_id, message)
		VALUES (
			(SELECT id FROM users WHERE username = ?),
			(SELECT id FROM users WHERE username = ?),
			?
		);`

	_, err := helpers.DataBase.Exec(query, sUname, rUname, msg)
	if err != nil {
		helpers.ErrorLog.Fatalln("Database insertion error:", err)
		return err
	}
	return nil
}

type User struct {
	Online   bool   `json:"online"`
	Username string `json:"username"` // Exported field
}

func GetUserNames(uid int) ([]User, error) {
	rows, err := helpers.DataBase.Query(`
	SELECT 
    	u.username
	FROM 
    	users u
	LEFT JOIN 
    	dms m 
	ON 
    	(u.id = m.sender_id OR u.id = m.recipient_id)
    AND
		(m.sender_id = ? OR m.recipient_id = ? )
	WHERE 
    	u.id != ?
	GROUP BY 
    	u.id, u.username
	ORDER BY 
    	CASE WHEN MAX(m.created_at) IS NOT NULL THEN 1 
        ELSE 2 
    END,
    	MAX(m.created_at) DESC,
    u.username ASC;`, uid, uid, uid)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %w", err)
	}
	defer rows.Close()

	var userNames []User

	for rows.Next() {
		var username User
		if err := rows.Scan(&username.Username); err != nil {
			return userNames, fmt.Errorf("could not scan row: %w", err)
		}
		if _, e := helpers.Sockets[username.Username]; e {
			username.Online = true
		}
		userNames = append(userNames, username)
	}

	if err := rows.Err(); err != nil {
		return userNames, fmt.Errorf("row iteration error: %w", err)
	}

	return userNames, nil
}
