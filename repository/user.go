package repository

import (
	"github.com/dg943/MyProject/backend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewUserRepository(db *gorm.DB, _logger logger.Interface) UserRepository {
	return &userRepository{db, _logger}
}

type userRepository struct {
	db     *gorm.DB
	logger logger.Interface
}

func (u *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := u.db.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) Update(user *models.User) error {
	if err := u.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Delete(id uint) error {
	if err := u.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Find(username, password string) (*models.User, error) {
	user := &models.User{}
	if err := u.db.Where("UserName = ? AND Password = ?", username, password).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) Create(user *models.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
