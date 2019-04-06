package gateway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro-search/go-search/domain/entities"
)

type Gateway struct {
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "asdd")
	bodyByte, _ := ioutil.ReadAll(r.Body)
	// varbll := string(bodyByte)
	product := entities.Product{}
	json.Unmarshal(bodyByte, &product)
	catalog := entities.Catalog{}
	catalog.AddProduct(product.Name, product.Description, product.Qty)
	fmt.Fprintln(w, product)
	// w.Write(val)
}

//
func Handle() {
	router := mux.NewRouter()
	productRoutes := GetProductRoutes()

	for _, route := range productRoutes {
		// fmt.Println(route.Method)
		router.HandleFunc(route.Name, handler).Methods(route.Method)
	}

	log.Fatal(http.ListenAndServe("localhost:8081", router))
	// fmt.Println(productRoutes)
	// fmt.Println(productRoutes[0])
	// fmt.Println(productRoutes[1])
}
