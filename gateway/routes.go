package gateway

type Router struct {
	Method string
	Name   string
}

func NewRouter(method, name string) *Router {
	return &Router{method, name}
}

func GetProductRoutes() []Router {
	routes := make([]Router, 0)
	addProductRoute := NewRouter("POST", "/products")
	searchProductRoute := NewRouter("GET", "/products/{name}/{description}")
	routes = append(routes, *addProductRoute)
	routes = append(routes, *searchProductRoute)
	return routes
}
