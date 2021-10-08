package web

type UserProfileCreateRequest struct {
	UserID    uint   `json:"user_id" gorm:"required,number"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Age       uint   `json:"age" binding:"number"`
}

type UserProfileUpdateRequest struct {
	ID        int    `json:"id" binding:"required,number"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Age       uint   `json:"age" binding:"number"`
}

type UserProfileResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Age       uint   `json:"age"`
}
