package repository

import (
	"github.com/jmoiron/sqlx"

	"crm-backend/internal/rybakcrm/app/domain/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetUserByCredentials(username, password string) (models.User, error) {
	var user models.User
	err := u.db.Get(&user, "SELECT id, username, email FROM users WHERE username = $1 AND password = $2", username, password)

	return user, err
}
