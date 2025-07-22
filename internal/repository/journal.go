package repository

import (
	"context"
	"database/sql"

	"github.com/LouisFernando1204/golang-restapi.git/domain"
	"github.com/doug-martin/goqu/v9"
)

type journalRepository struct {
	db *goqu.Database
}

func NewJournal(con *sql.DB) domain.JournalRepository {
	return &journalRepository{
		db: goqu.New("default", con),
	}
}

// Find implements domain.JournalRepository.
func (j *journalRepository) Find(ctx context.Context, se domain.JournalSearch) (result []domain.Journal, err error) {
	dataset := j.db.From("journals")
	if se.CustomerId != "" {
		dataset = dataset.Where(goqu.C("customer_id").Eq(se.CustomerId))
	}
	if se.Status != "" {
		dataset = dataset.Where(goqu.C("status").Eq(se.Status))
	}
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

// FindById implements domain.JournalRepository.
func (j *journalRepository) FindById(ctx context.Context, id string) (result domain.Journal, err error) {
	dataset := j.db.From("journals").Where(goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// Save implements domain.JournalRepository.
func (j *journalRepository) Save(ctx context.Context, journal *domain.Journal) error {
	executor := j.db.Insert("journals").Rows(journal).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

// Update implements domain.JournalRepository.
func (j *journalRepository) Update(ctx context.Context, journal *domain.Journal) error {
	executor := j.db.Update("journals").Where(goqu.C("id").Eq(journal.Id)).Set(journal).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
