package handlers

import (
	"EP4/data"
	"net/http"
)

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProducts(&prod)
}
