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
	DataGetErr     = errors.New("Product can not be get")
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
	return err
	//todo
}

func (c *Catalog) GetAllProducts() ([]Product, error) {
	conn := data.NewConnection()
	db := conn.Connect()

	res, err := db.Query("Select * from product;")

	if err != nil {
		return nil, DataGetErr
	}

	product := Product{}
	for res.Next() {
		err := res.Scan(&product.Name, &product.Description, &product.Qty)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", product)
		c.products = append(c.products, product)
	}
	return c.products, nil
}

func (c *Catalog) SearchProducts(name, description string) ([]Product, error) {
	conn := data.NewConnection()
	db := conn.Connect()
	query := "Select * from product where name like '%" + name + "%' or description like '%" + description + "%';"
	res, err := db.Query(query)

	if err != nil {
		return nil, DataGetErr
	}

	product := Product{}
	for res.Next() {
		err := res.Scan(&product.Name, &product.Description, &product.Qty)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", product)
		c.products = append(c.products, product)
	}
	return c.products, nil
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
