package request

type CreateUserRequest struct {
	Name      string `validate:"required,min=1,max=100" json:"name"`
	Age       uint    `validate:"required" json:"age"`
}
