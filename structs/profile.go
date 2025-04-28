package structs

type User struct {
	Age       int    `json:"age"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
}
