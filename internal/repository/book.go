package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/LouisFernando1204/golang-restapi.git/domain"
	"github.com/doug-martin/goqu/v9"
)

type bookRepository struct {
	db *goqu.Database
}

func NewBook(con *sql.DB) domain.BookRepository {
	return &bookRepository{
		db: goqu.New("default", con),
	}
}

// FindAll implements domain.BookRepository.
func (b *bookRepository) FindAll(ctx context.Context) (results []domain.Book, err error) {
	dataset := b.db.From("books").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &results)
	return results, err
}

// FindById implements domain.BookRepository.
func (b *bookRepository) FindById(ctx context.Context, id string) (result domain.Book, err error) {
	dataset := b.db.From("books").Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return result, err
}

// FindByIds implements domain.BookRepository.
func (b *bookRepository) FindByIds(ctx context.Context, ids []string) (results []domain.Book, err error) {
	dataset := b.db.From("books").Where(goqu.C("deleted_at").IsNull(), goqu.C("id").In(ids))
	err = dataset.ScanStructsContext(ctx, &results)
	return results, err
}

// Save implements domain.BookRepository.
func (b *bookRepository) Save(ctx context.Context, c *domain.Book) error {
	executor := b.db.Insert("books").Rows(c).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

// Update implements domain.BookRepository.
func (b *bookRepository) Update(ctx context.Context, c *domain.Book) error {
	executor := b.db.Update("books").Where(goqu.C("id").Eq(c.Id)).Set(c).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

// Delete implements domain.BookRepository.
func (b *bookRepository) Delete(ctx context.Context, id string) error {
	executor := b.db.Update("books").Where(goqu.C("id").Eq(id)).Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
