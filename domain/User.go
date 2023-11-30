package domain

type User struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Confirm  string `form:"confirm" json:"confirm"`
}
