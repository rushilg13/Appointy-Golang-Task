package users

type User struct {
	Name     string `json:"name"`
	Id       string `json:"id"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}
