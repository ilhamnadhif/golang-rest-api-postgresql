package web

type BookCreateRequest struct {
	Title       string  `json:"title" binding:"required,gte=6"`
	Description string  `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
	Rating      float32 `json:"rating" binding:"required,number"`
}

type BookUpdateRequest struct {
	ID          int    `json:"id" binding:"required,number"`
	Title       string  `json:"title" binding:"required;gte=6"`
	Description string  `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
	Rating      float32 `json:"rating" binding:"required,number"`
}
