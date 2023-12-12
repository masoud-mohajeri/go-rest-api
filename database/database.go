package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(dbUser string, dbPassword string, dbName string) (*gorm.DB, error) {
	var connectionString = fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/%s",
		dbUser, dbPassword, dbName,
	)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	return db, err
}
