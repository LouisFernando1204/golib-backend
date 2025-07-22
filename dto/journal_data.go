package dto

import "time"

type JournalData struct {
	Id         string       `json:"id"`
	BookStock  string       `json:"book_stock"`
	Book       BookData     `json:"book"`
	Customer   CustomerData `json:"customer"`
	Status     string       `json:"status"`
	BorrowedAt time.Time    `json:"borrowed_at"`
	ReturnedAt time.Time    `json:"returned_at"`
}

type CreateJournalRequest struct {
	BookId     string `json:"book_id" validate:"required"`
	BookStock  string `json:"book_stock" validate:"required"`
	CustomerId string `json:"customer_id" validate:"required"`
}

type ReturnJournalRequest struct {
	JournalId string `json:"-"`
	UserId    string `json:"-"`
}
