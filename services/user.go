package services

import (
	"time"

	"github.com/dg943/MyProject/backend/helpers"
	"github.com/dg943/MyProject/backend/models"
	"github.com/dg943/MyProject/backend/repository"
	"gorm.io/gorm/logger"
)

type userService struct {
	userRepository repository.UserRepository
	logger         logger.Interface
}

func NewUserService(r repository.UserRepository, logger logger.Interface) UserService {
	return &userService{r, logger}
}

func (s *userService) FindAll() ([]models.User, error) {
	return s.userRepository.FindAll()
}

func (s *userService) FindByID(id uint) (*models.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *userService) Update(user *models.User) error {
	return s.userRepository.Update(user)
}

func (s *userService) Delete(id uint) error {
	return s.userRepository.Delete(id)
}

func (s *userService) Login(username, password string) (string, error) {
	// Todo := add logic after getting the user
	_, err := s.userRepository.Find(username, password)
	if err != nil {
		return "", err
	}
	claims := getClaims()
	jwtTokenString, err := helpers.EncodeJWT(&claims)
	if err != nil {
		return "", err
	}
	return jwtTokenString, nil
}

func (s *userService) Signup(user *models.User) (string, error) {
	err := s.userRepository.Create(user)
	if err != nil {
		return "", err
	}
	claims := getClaims()
	jwtTokenString, err := helpers.EncodeJWT(&claims)
	if err != nil {
		return "", err
	}
	return jwtTokenString, nil
}

func getClaims() helpers.Claims {
	return helpers.NewClaims(map[string]interface{}{
		"exp": time.Now().Add(time.Minute * 60).Unix(),
		"iss": "dina",
		"sub": "practice",
		"aud": "external",
		"iat": time.Now().Unix(),
	})
}
