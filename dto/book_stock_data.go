package dto

type CreateBookStockRequest struct {
	BookId string   `json:"book_id" validate:"required"`
	Codes  []string `json:"codes" validate:"required,unique,min=1"`
}

type DeleteBookStockRequest struct {
	Codes []string
}
