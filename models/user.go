package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"username"`
	Password string `json:"password"`
}

func CreateUserTable(db *gorm.DB) {
	fmt.Println(db.Config)
	if has := db.Migrator().HasTable(&User{}); !has {
		if err := db.Migrator().CreateTable(&User{}); err != nil {
			log.Fatal("Error in creating the user table")
		}
	}
}

func (u *User) Print() string {
	return fmt.Sprintf("Id : %d\nUsername : %s\nPassword : %s\n", u.ID, u.UserName, u.Password)
}
