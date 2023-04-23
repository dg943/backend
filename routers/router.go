//Package router is used to instantiate the router
package routers

import (
	"github.com/dg943/MyProject/backend/configs"
	"github.com/dg943/MyProject/backend/controllers"
	"github.com/dg943/MyProject/backend/middlewares"
	"github.com/dg943/MyProject/backend/repository"
	"github.com/dg943/MyProject/backend/services"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewRouter is used to create new router
func NewRouter(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	r.Use(middlewares.AuthenticateMiddleWare)
	addRoutes(r, db)
	return r
}

var (
	_logger logger.Interface
)

func addRoutes(r *mux.Router, db *gorm.DB) {

	// get logger
	_logger = configs.GetLogger()

	// add user routes
	userRepo := repository.NewUserRepository(db, _logger)
	userService := services.NewUserService(userRepo, _logger)
	userController := controllers.NewUserController(userService, _logger)
	r.HandleFunc("/login", userController.Login).Methods("POST")
	r.HandleFunc("/signup", userController.Signup).Methods("POST")
	r.HandleFunc("/test", userController.Test).Methods("POST")
	//r.HandleFunc("/users", ListUsers).Methods("GET")
	//r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	//r.HandleFunc("/users", CreateUser).Methods("POST")
	//r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	//r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

}
