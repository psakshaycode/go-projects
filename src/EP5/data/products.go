package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProducts(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}
func UpdateProducts(id int, p *Product) error {
	pos, _, err := findProduct(id)
	if err != nil {
		return err
	}
	productList[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (int, *Product, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, p, nil
		}
	}
	return -1, nil, ErrProductNotFound
}
func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
