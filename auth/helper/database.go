package helper

import (
	"database/sql"
	"errors"
	"net/url"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*sql.DB, error) {
	connString := os.Getenv("DB_CONN_STRING")

	if connString == "" {
		return nil, errors.New("DB_CONN_STRING not defined")
	}

	timeZone := os.Getenv("DB_TIMEZONE")

	if timeZone != "" {
		connString += "&loc=" + url.QueryEscape(timeZone)
	}

	conn, err := gorm.Open(mysql.Open(connString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return conn.DB()
}
