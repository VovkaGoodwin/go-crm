package repository

import "crm-backend/internal/rybakcrm/app/domain/models"

type UserRepository interface {
	GetUserByCredentials(username, password string) (models.User, error)
}
