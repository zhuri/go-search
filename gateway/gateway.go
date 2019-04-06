package gateway

import (
	"fmt"
)

type Gateway struct {
}

// w http.ResponseWriter, r *http.Request
func Handle() {
	//router := mux.NewRouter()
	productRoutes := GetProductRoutes()
	fmt.Println(productRoutes)
}
