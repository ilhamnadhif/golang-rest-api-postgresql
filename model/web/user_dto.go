package web

type UserRequest struct {
	Email    string `json:"email" binding:"required,gte=6"`
	Password string `json:"password" binding:"required,gte=6"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

/////////////////////////////////

type ProfileResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Age       uint   `json:"age"`
}

type UserWithProfileResponse struct {
	ID          uint            `json:"id"`
	Email       string          `json:"email"`
	UserProfile ProfileResponse `json:"user_profile"`
}
