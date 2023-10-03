package request

type UpdateUserRequest struct {
	Id        int    `validate:"required" json:"id"`
	Name      string `validate:"required,min=1,max=10" json:"name"`
	Age       uint    `validate:"required" json:"age"`
}
