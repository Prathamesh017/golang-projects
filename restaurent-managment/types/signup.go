package types
type SignupRequest struct{
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type SignupResponse struct{
	ID int `json:"id"`
	Token string `json:"token"`
}