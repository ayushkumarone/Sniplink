package users

type User struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Linkscreated int    `json:"number-of-links"`
}
