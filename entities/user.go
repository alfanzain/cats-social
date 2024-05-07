package entities

type User struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email" validate:"required,email,min=3,max=200"`
}

type UserRegisterInput struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}
