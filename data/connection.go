package data

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	Username string
	Password string
	Database string
}

func NewConnection() *Connection {
	return &Connection{
		"endrit",
		"endrit",
		"catalog",
	}
}

func (c *Connection) Connect() *sql.DB {
	credentials := c.Username + ":" + c.Password + "@/" + c.Database
	db, err := sql.Open("mysql", credentials)

	if err != nil {
		fmt.Println("err", err)
	}

	return db
}
