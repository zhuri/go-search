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
		"pms",
	}
}

func (c *Connection) Connect() {
	credentials := c.Username + ":" + c.Password + "@/" + c.Database
	fmt.Println(credentials)
	db, err := sql.Open("mysql", credentials)

	if err != nil {
		fmt.Println("success", db)
	}

	fmt.Println("errr", err)
}
