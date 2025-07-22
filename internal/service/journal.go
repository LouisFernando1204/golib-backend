package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/LouisFernando1204/golang-restapi.git/domain"
	"github.com/LouisFernando1204/golang-restapi.git/dto"
	"github.com/google/uuid"
)

type journalService struct {
	journalRepository   domain.JournalRepository
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
	customerRepository  domain.CustomerRepository
	chargeRepository    domain.ChargeRepository
}

func NewJournal(journalRepository domain.JournalRepository, bookRepository domain.BookRepository, bookStockRepository domain.BookStockRepository, customerRepository domain.CustomerRepository, chargeRepository domain.ChargeRepository) domain.JournalService {
	return &journalService{
		journalRepository:   journalRepository,
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
		customerRepository:  customerRepository,
		chargeRepository:    chargeRepository,
	}
}

// Index implements domain.JournalService.
func (j *journalService) Index(ctx context.Context, se domain.JournalSearch) ([]dto.JournalData, error) {
	journals, err := j.journalRepository.Find(ctx, se)
	if err != nil {
		return []dto.JournalData{}, err
	}

	customerIds := make([]string, 0)
	bookIds := make([]string, 0)
	for _, v := range journals {
		customerIds = append(customerIds, v.CustomerId)
		bookIds = append(bookIds, v.BookId)
	}

	customers := make(map[string]domain.Customer)
	if len(customerIds) > 0 {
		customersDb, _ := j.customerRepository.FindByIds(ctx, customerIds)
		for _, v := range customersDb {
			customers[v.ID] = v
		}
	}
	books := make(map[string]domain.Book)
	if len(bookIds) > 0 {
		booksDB, _ := j.bookRepository.FindByIds(ctx, bookIds)
		for _, v := range booksDB {
			books[v.Id] = v
		}
	}

	results := make([]dto.JournalData, 0)
	for _, v := range journals {
		book := dto.BookData{}
		if v2, e := books[v.BookId]; e {
			book = dto.BookData{
				Id:          v2.Id,
				Isbn:        v2.Isbn,
				Title:       v2.Title,
				Description: v2.Description,
			}
		}
		customer := dto.CustomerData{}
		if v2, e := customers[v.CustomerId]; e {
			customer = dto.CustomerData{
				ID:   v2.ID,
				Code: v2.Code,
				Name: v2.Name,
			}
		}

		results = append(results, dto.JournalData{
			Id:         v.Id,
			BookStock:  v.StockCode,
			Book:       book,
			Customer:   customer,
			Status:     v.Status,
			BorrowedAt: v.BorrowedAt.Time,
			ReturnedAt: v.ReturnedAt.Time,
		})
	}
	return results, nil
}

// Create implements domain.JournalService.
func (j *journalService) Create(ctx context.Context, req dto.CreateJournalRequest) error {
	book, err := j.bookRepository.FindById(ctx, req.BookId)
	if err != nil {
		return err
	}
	if book.Id == "" {
		return domain.ErrBookNotFound
	}

	stock, err := j.bookStockRepository.FindByBookAndCode(ctx, req.BookId, req.BookStock)
	if err != nil {
		return err
	}
	if stock.Code == "" {
		return domain.ErrBookNotFound
	}
	if stock.Status != domain.BookStockStatusAvailable {
		return errors.New("book stock have already been borrowed before")
	}

	journal := domain.Journal{
		Id:         uuid.NewString(),
		BookId:     req.BookId,
		StockCode:  req.BookStock,
		CustomerId: req.CustomerId,
		Status:     domain.JournalStatusInProgress,
		DueAt:      sql.NullTime{Valid: true, Time: time.Now().Add(7 * 24 * time.Hour)},
		BorrowedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	err = j.journalRepository.Save(ctx, &journal)
	if err != nil {
		return err
	}

	stock.Status = domain.BookStockStatusBorrowed
	stock.BorrowerAt = journal.BorrowedAt
	stock.BorrowerId = sql.NullString{Valid: true, String: journal.CustomerId}

	return j.bookStockRepository.Update(ctx, &stock)
}

// Return implements domain.JournalService.
func (j *journalService) Return(ctx context.Context, req dto.ReturnJournalRequest) error {
	journal, err := j.journalRepository.FindById(ctx, req.JournalId)
	if err != nil {
		return err
	}
	if journal.Id == "" {
		return domain.ErrJournalNotFound
	}

	stock, err := j.bookStockRepository.FindByBookAndCode(ctx, journal.BookId, journal.StockCode)
	if err != nil {
		return err
	}

	if stock.Code != "" {
		stock.Status = domain.BookStockStatusAvailable
		stock.BorrowerId = sql.NullString{Valid: false}
		stock.BorrowerAt = sql.NullTime{Valid: false}
		err = j.bookStockRepository.Update(ctx, &stock)
		if err != nil {
			return err
		}
	}

	journal.Status = domain.JournalStatusCompleted
	journal.ReturnedAt = sql.NullTime{Valid: true, Time: time.Now()}
	err = j.journalRepository.Update(ctx, &journal)
	if err != nil {
		return err
	}

	hoursLate := time.Since(journal.DueAt.Time).Hours()
	if hoursLate >= 24 {
		daysLate := int(hoursLate / 24)
		charge := domain.Charge{
			Id:           uuid.NewString(),
			JournalId:    journal.Id,
			DaysLate:     daysLate,
			DailyLateFee: 5000,
			Total:        5000 * daysLate,
			UserId:       req.UserId,
			CreatedAt:    sql.NullTime{Valid: true, Time: time.Now()},
		}
		err = j.chargeRepository.Save(ctx, &charge)
	}
	return err
}
