package repository

import (
	"github.com/dg943/MyProject/backend/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	Update(*models.User) error
	Delete(id uint) error
	Find(username, password string) (*models.User, error)
	Create(*models.User) error
}
