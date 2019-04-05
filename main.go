package main

import (
	"fmt"
	"log"

	"github.com/micro-search/go-search/data"
	"github.com/micro-search/go-search/domain/entities"
)

func main() {
	conn := data.NewConnection()
	fmt.Println(conn)
	db := conn.Connect()

	res, err := db.Query("Select * from product;")

	if err != nil {
		fmt.Println("Error happened", err)
	}

	// fmt.Println("Result", res)

	product := entities.Product{}
	for res.Next() {
		err := res.Scan(&product.ID, &product.Name, &product.Description, &product.Qty)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", product)
	}
}
