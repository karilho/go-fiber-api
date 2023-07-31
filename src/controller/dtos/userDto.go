package dtos

type UserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Age      int8   `json:"age" validate:"required"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
}

type UserUpdateRequest struct {
	Name string `json:"name" validate:"omitempty,min=3,max=50"`
	Age  int8   `json:"age" validate:"omitempty"`
}
