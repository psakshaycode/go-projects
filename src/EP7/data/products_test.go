package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Ps",
		Price: 1,
		SKU:   "abs-abc-abc",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
