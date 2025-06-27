package response

type UserResponse struct {
	ID       string `json:"id"`
	Age      int8   `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
