package service

import (
	"context"
	"database/sql"
	"path"
	"time"

	"github.com/LouisFernando1204/golang-restapi.git/domain"
	"github.com/LouisFernando1204/golang-restapi.git/dto"
	"github.com/LouisFernando1204/golang-restapi.git/internal/config"
	"github.com/google/uuid"
)

type bookService struct {
	cnf                 *config.Config
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
	mediaRepository     domain.MediaRepository
}

func NewBook(cnf *config.Config, bookRepository domain.BookRepository, bookStockRepository domain.BookStockRepository, mediaRepository domain.MediaRepository) domain.BookService {
	return &bookService{
		cnf:                 cnf,
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
		mediaRepository:     mediaRepository,
	}
}

// Index implements domain.BookService.
func (b *bookService) Index(ctx context.Context) ([]dto.BookData, error) {
	books, err := b.bookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	coverIds := make([]string, 0)
	for _, v := range books {
		if v.CoverId.Valid {
			coverIds = append(coverIds, v.CoverId.String)
		}
	}

	covers := make(map[string]string)
	if len(coverIds) > 0 {
		coversDB, _ := b.mediaRepository.FindByIds(ctx, coverIds)
		for _, v := range coversDB {
			covers[v.Id] = path.Join(b.cnf.Server.Asset, v.Path)
		}
	}

	var bookData []dto.BookData
	for _, v := range books {
		var coverUrl string
		if v, e := covers[v.CoverId.String]; e {
			coverUrl = v
		}

		bookData = append(bookData, dto.BookData{
			Id:          v.Id,
			Isbn:        v.Isbn,
			Title:       v.Title,
			CoverUrl:    coverUrl,
			Description: v.Description,
		})
	}

	return bookData, nil
}

// Show implements domain.BookService.
func (b *bookService) Show(ctx context.Context, id string) (dto.BookShowData, error) {
	data, err := b.bookRepository.FindById(ctx, id)
	if err != nil {
		return dto.BookShowData{}, err
	}

	if data.Id == "" {
		return dto.BookShowData{}, domain.ErrBookNotFound
	}

	stocks, err := b.bookStockRepository.FindByBookId(ctx, id)
	if err != nil {
		return dto.BookShowData{}, err
	}

	stocksData := make([]dto.BookStockData, 0)
	for _, v := range stocks {
		stocksData = append(stocksData, dto.BookStockData{
			Code:   v.Code,
			Status: v.Status,
		})
	}

	var coverUrl string
	if data.CoverId.Valid {
		cover, _ := b.mediaRepository.FindById(ctx, data.CoverId.String)
		if cover.Path != "" {
			coverUrl = path.Join(b.cnf.Server.Asset, cover.Path)
		}
	}

	return dto.BookShowData{
		BookData: dto.BookData{
			Id:          data.Id,
			Isbn:        data.Isbn,
			Title:       data.Title,
			CoverUrl:    coverUrl,
			Description: data.Description,
		},
		Stocks: stocksData,
	}, nil
}

// Create implements domain.BookService.
func (b *bookService) Create(ctx context.Context, req dto.CreateBookRequest) error {
	coverId := sql.NullString{Valid: false, String: req.CoverId}
	if req.CoverId != "" {
		coverId.Valid = true
	}

	book := domain.Book{
		Id:          uuid.NewString(),
		Isbn:        req.Isbn,
		Title:       req.Title,
		Description: req.Description,
		CoverId:     coverId,
		CreatedAt:   sql.NullTime{Valid: true, Time: time.Now()},
	}

	return b.bookRepository.Save(ctx, &book)
}

// Update implements domain.BookService.
func (b *bookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {
	coverId := sql.NullString{Valid: false, String: req.CoverId}
	if req.CoverId != "" {
		coverId.Valid = true
	}

	persisted, err := b.bookRepository.FindById(ctx, req.Id)
	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return domain.ErrBookNotFound
	}

	persisted.Isbn = req.Isbn
	persisted.Title = req.Title
	persisted.Description = req.Description
	persisted.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}
	persisted.CoverId = coverId

	return b.bookRepository.Update(ctx, &persisted)
}

// Delete implements domain.BookService.
func (b *bookService) Delete(ctx context.Context, id string) error {
	persisted, err := b.bookRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return domain.ErrBookNotFound
	}

	err = b.bookRepository.Delete(ctx, persisted.Id)
	if err != nil {
		return err
	}

	return b.bookStockRepository.DeleteByBookId(ctx, persisted.Id)
}
