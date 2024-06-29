package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type GenericDatabaseContext interface {
	Connect()
	GetConfiguredSqlDb() *sql.DB
}

type DatabaseContext struct {
	ConnectionString string
	SqlDb            *sql.DB
}

func (c *DatabaseContext) Connect() {

	db, err := sql.Open("mysql", c.ConnectionString)
	if err != nil {
		log.Fatal("Connection to database cannot be established :", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database at %s cannot be reached : %s", c.ConnectionString, err)
	}

	c.SqlDb = db

}

func (c *DatabaseContext) GetConfiguredSqlDb() *sql.DB {
	return c.SqlDb
}
