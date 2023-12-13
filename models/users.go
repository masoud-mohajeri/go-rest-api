package models

import (
	"database/sql/driver"
	"time"
)

// https://stackoverflow.com/a/68637612
type UserType string

const (
	Admin    UserType = "Admin"
	Customer UserType = "Customer"
)

func (ut *UserType) Scan(value interface{}) error {
	*ut = UserType(value.([]byte))
	return nil
}

func (ut UserType) Value() (driver.Value, error) {
	return string(ut), nil
}

type User struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string   `gorm:"varchar(255);required;not null" json:"name"`
	Username  string   `gorm:"varchar(255);unique" json:"username"`
	Password  string   `gorm:"varchar(255)" json:"password"`
	Role      UserType `gorm:"type:user_type" json:"role"`
	// TODO : add a contactInfo filed and make it nested
}

// TODO: Automate this query
//CREATE TYPE user_type AS ENUM (
//'Admin',
//'Customer');
