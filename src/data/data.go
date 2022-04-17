package data

import (
	"fmt"
	"github.com/fatihbasol/GoGinExample/src/data/migrate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	appName  = "Go Gin Example"
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "go-gin-example-db"
)

var DB *gorm.DB

func ConnectToDatabase() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable application_name='%s'", host, port, user, password, dbname, appName)
	database, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		panic("db error")
	}
	DB = database

	migrate.MigrateModels(DB)
}
