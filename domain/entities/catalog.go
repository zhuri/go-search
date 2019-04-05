package entities

import (
	"errors"
	"fmt"
	"log"

	"github.com/micro-search/go-search/data"
)

type Catalog struct {
	ID       int32
	Name     string
	products []Product
}

var (
	InvalidNameErr = errors.New("Name is not valid")
	InvalidQtyErr  = errors.New("Invalid price")
	DataInsertErr  = errors.New("Product can not be inserted")
)

func NewCatalog() *Catalog {
	return &Catalog{
		products: make([]Product, 0),
	}
}

func (c *Catalog) AddProduct(name, description string, qty int) error {

	err := productDataConstraints(name, description, qty)
	if err != nil {
		return err
	}

	conn := data.NewConnection()
	db := conn.Connect()

	prep, err := db.Prepare("Insert into product(name, description, qty) values(?,?,?)")
	res, err := prep.Exec(name, description, qty)

	if err != nil {
		return DataInsertErr
	}
	fmt.Println(res)

	return nil
}

func (c *Catalog) UpdateProduct(name, description string, qty int) error {

	err := productDataConstraints(name, description, qty)
	if err != nil {
		return err
	}

	return nil
}

func (c *Catalog) GetAllProducts() {
	conn := data.NewConnection()
	db := conn.Connect()

	res, err := db.Query("Select * from product;")

	if err != nil {
		fmt.Println("Error happened", err)
	}

	product := Product{}
	for res.Next() {
		err := res.Scan(&product.ID, &product.Name, &product.Description, &product.Qty)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", product)
	}
}

func productDataConstraints(name, description string, qty int) error {

	if name == "" {
		return InvalidNameErr
	}

	if len(name) > 10 {
		return InvalidNameErr
	}

	return nil
}
