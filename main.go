package main

import "fmt"
import "github.com/micro-search/go-search/data"

func main() {
	conn := data.NewConnection()
	fmt.Println(conn)
	conn.Connect()
}
