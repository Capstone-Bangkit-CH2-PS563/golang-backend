package payload

type GetAllUserResponse struct {
	ID    uint   `json:"ID"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

type CreateUserRequest struct {
	Name            string `json:"name" form:"name" validate:"required,max=20"`
	Email           string `json:"email" form:"email" validate:"required,email"`
	Password        string `json:"password" form:"password" validate:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type UpdateUserRequest struct {
	Name string `json:"name" form:"name"`
	Role string `json:"role" form:"role"`
}

type CreateUserResponse struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
}

type UpdateUserResponse struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

type LoginUserRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type LoginUserResponse struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}
