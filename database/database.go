package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	//var connectionString = fmt.Sprintf(
	//	"postgres://%s:%s@localhost:5432/%s",
	//	dbUser, dbPassword, dbName,
	//)

	//db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	return gorm.Open(postgres.Open(dbURL), &gorm.Config{})
}
