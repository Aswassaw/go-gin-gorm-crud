package response

type UserResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Age       uint    `json:"age"`
}
