package configs

import (
	"context"
	"sync"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	db_once sync.Once
	db      *gorm.DB
)

func GetDB() (*gorm.DB, error) {
	var err error
	_logger := GetLogger()
	db_once.Do(func() {
		conn_str := GetString("app_settings.db_details")
		_logger.Info(context.Background(), conn_str)
		db, err = gorm.Open(sqlserver.Open(conn_str), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: false,
				NoLowerCase:   true,
			},
			PrepareStmt: true,
			Logger:      _logger,
		})
		if err != nil {
			panic("Not able to connect to DB")
		}
	})
	sqlDB, err := db.DB()
	if err != nil {
		_logger.Error(context.Background(), err.Error())
	}
	if err := sqlDB.Ping(); err != nil {
		// database is not connected
		_logger.Error(context.Background(), err.Error())
	} else {
		// database is connected
		_logger.Info(context.Background(), "Database is not connected")
	}
	_logger.Info(context.Background(), "Db successfully connected")
	_logger.Info(context.Background(), db.Migrator().CurrentDatabase())
	return db, nil
}

//func RunMigrations(sqlDB *sql.DB) {
//// goose will look for this file for configuration
//os.Setenv("GOOSE_CONFIG", "/Users/dinanathgupta/go/github.com/dg943/MyProject/backend/configs/dbconf.yml")
//if err := goose.Up(sqlDB, "../migrations"); err != nil {
//log.Fatalf("failed to migrate database: %v\n", err)
//}
//}
