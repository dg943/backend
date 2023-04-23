package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dg943/MyProject/backend/helpers"
	"github.com/dg943/MyProject/backend/models"
	"github.com/dg943/MyProject/backend/services"
	"gorm.io/gorm/logger"
)

type UserController struct {
	userService services.UserService
	logger      logger.Interface
}

func NewUserController(s services.UserService, logger logger.Interface) UserController {
	return UserController{s, logger}
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error parsing user", http.StatusBadRequest)
		return
	}
	if user.Password == "" || user.UserName == "" {
		http.Error(w, "Username and Password can not be empty", http.StatusBadRequest)
		return
	}
	jwtToken, err := c.userService.Login(user.UserName, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token := helpers.Token{
		Bearer: jwtToken,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(token); err != nil {
		http.Error(w, "Error parsing user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *UserController) Signup(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error parsing user", http.StatusBadRequest)
		return
	}
	if user.Password == "" || user.UserName == "" {
		http.Error(w, "Username and Password can not be empty", http.StatusBadRequest)
		return
	}
	jwtToken, err := c.userService.Signup(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+jwtToken)
	// Redirect the user to the dashboard page
	//http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (c *UserController) Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "dindin")
}
