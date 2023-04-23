package services

import (
	"github.com/dg943/MyProject/backend/models"
)

type UserService interface {
	FindAll() ([]models.User, error)
	FindByID(uint) (*models.User, error)
	Update(*models.User) error
	Delete(uint) error
	Login(string, string) (string, error)
	Signup(*models.User) (string, error)
}
