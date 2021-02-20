package catalogue

/*
This package is to create and maintain a Catalogue of Products
*/

import (
	"fmt"

	"github.com/doyyan/ECS/product"
)

// Catalogue - is a list of products available for sale in a Shop.
type Catalogue struct {
	Name     string
	Products []product.Product
}

var matched bool = false

// NewCatalogue - a Method to set the Products of a catalogue
func NewCatalogue(Products []product.Product) Catalogue {
	return Catalogue{Products: Products}
}

// SetProducts - a Method to set the Products of a catalogue
func (c *Catalogue) SetProducts(Products []product.Product) {
	c.Products = Products
}

// SetOffers- a Method to set the Offers for some/all Products in a Catalouge..
func (c *Catalogue) SetOffers(Products []product.Product) {

	for _, offerProducts := range Products {
		matched = false
		for _, product := range c.Products {

			if offerProducts.ID == product.ID {

				matched = true
			}

		}
	}

}

// String - a Stringer Interface implementation function to describe the type created.
func (c Catalogue) String() string {
	return fmt.Sprintf(" products in the catalogue ", c.Products)
}
