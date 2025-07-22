package repository

import (
	"context"
	"database/sql"

	"github.com/LouisFernando1204/golang-restapi.git/domain"
	"github.com/doug-martin/goqu/v9"
)

type userRepository struct {
	db *goqu.Database
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &userRepository{
		db: goqu.New("default", con),
	}
}

// FindByEmail implements domain.UserRepository.
func (u *userRepository) FindByEmail(ctx context.Context, email string) (usr domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.C("email").Eq(email))
	_, err = dataset.ScanStructContext(ctx, &usr)
	return usr, err
}
