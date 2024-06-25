package driver

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// This file is to connect our application to the database

// DB holds the database connection pool
type DB struct {

	// driver for mysql database
	// different database may have different drivers
	SQL *sql.DB
}

var dbConn = &DB{}

// these constants define the nature of the connection pool
const maxOpenDbConn = 10
const maxIdelDbConn = 5
const maxDbLifetime = 5 * time.Minute

// ConnectSQL creates database connection pool for mysql
// dsn = database connection string
func ConnectSQL(dsn string) (*DB, error) {

	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdelDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}
	return dbConn, nil

}

// testDB tries to ping the database
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}

	return nil
}

// NewDatabase create new database for application
func NewDatabase(dsn string) (*sql.DB, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
