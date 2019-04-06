package entities

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Qty         int    `json:"qty"`
}
