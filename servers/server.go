package servers

import (
	"net/http"

	"github.com/dg943/MyProject/backend/configs"
	"github.com/dg943/MyProject/backend/routers"
	"github.com/gorilla/handlers"
	"gorm.io/gorm"
)

func NewServer(db *gorm.DB) *http.Server {
	host := configs.GetString("app_settings.host_details.host")
	port := configs.GetString("app_settings.host_details.port")
	return &http.Server{
		Addr: host + ":" + port,
		// this is for cors
		Handler: handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(routers.NewRouter(db)),
	}
}
