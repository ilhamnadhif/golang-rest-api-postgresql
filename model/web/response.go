package web

type BookResponse struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       uint    `json:"price"`
	Rating      float32 `json:"rating"`
}
