package types

type LoginRequest struct{
	Email string `json:"email"`
	Password string `json:"password" binding:"required,min=6"`
}