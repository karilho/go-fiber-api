package dtos

type UserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Age      int    `json:"age" validate:"required"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}
