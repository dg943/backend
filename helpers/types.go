package helpers

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims = jwt.MapClaims

type Empty = struct{}

type Token struct {
	Bearer string `json:"bearer"`
}

func NewClaims(claims map[string]interface{}) Claims {
	return Claims(claims)
}

//func RunMigrations(sqlDB *sql.DB) {
//// goose will look for this file for configuration
//os.Setenv("GOOSE_CONFIG", "/Users/dinanathgupta/go/github.com/dg943/MyProject/backend/configs/dbconf.yml")
//if err := goose.Up(sqlDB, "../migrations"); err != nil {
//log.Fatalf("failed to migrate database: %v\n", err)
//}
//}
