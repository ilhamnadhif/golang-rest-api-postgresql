package web

type BookCreateRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       uint    `json:"price" binding:"required,number"`
	Rating      float32 `json:"rating" binding:"required,number"`
}

type BookUpdateRequest struct {
	ID          uint    `json:"id" binding:"required,number"`
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       uint    `json:"price" binding:"required,number"`
	Rating      float32 `json:"rating" binding:"required,number"`
}
