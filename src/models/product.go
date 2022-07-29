package models

type ProductModel struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Product string `json:"product"`
	Price   string `json:"price"`
	Qty     string `json:"qty"`
}
