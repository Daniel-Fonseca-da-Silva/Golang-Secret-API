package request

type UserRequest struct {
	Age      int8   `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
