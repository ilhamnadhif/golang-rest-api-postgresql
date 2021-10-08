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


type UserWithProfileResponse struct {
	ID          uint                `json:"id"`
	Email       string              `json:"email"`
	UserProfile UserProfileResponse `json:"user_profile"`
	Books       []BookResponse      `json:"books"`
}
