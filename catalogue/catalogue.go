package catalogue

/*
This package is to create and maintain a Catalogue of Products
*/

import (
	"fmt"

	"github.com/doyyan/ECS/datatypes"
)

// Catalogue - is a list of products available for sale in a Shop.
type Catalogue struct {
	// Name of the catalogue, possibly the Store name
	Name string
	// ID a unique id for the catalogue
	ID int
	// Products A slice of Pointers to products, essential they are pointers to apply Offers
	Products []*datatypes.Product
}

var matched bool = false

// NewCatalogue - a Method to set the Products of a catalogue
func NewCatalogue(name string, ID int, Products []*datatypes.Product) Catalogue {
	return Catalogue{Name: name, ID: ID, Products: Products}
}

// SetProducts - a Method to set the Products of a catalogue
func (c *Catalogue) SetProducts(Products []*datatypes.Product) {
	c.Products = Products
}

// SetOffers - a Method to set the Offers for some/all Products in a Catalouge..
// Browse through the products and try and find the Product in each of the offer
// If found just set the offerID in the Product for now
// If not found return an error code as the team setting the Offer may be different
// to the team managing the Catalogue and its good to send a meaningful message back!!
func (c *Catalogue) SetOffers(offers []datatypes.Offer) []datatypes.Offer {
	// Range through the offers
	for _, offerProduct := range offers {
		// This flag is set to check to see if the intended product exists in the catalogue, if not the Caller is notified it isn't to take further action!!
		matched = false
		for _, product := range c.Products {
			// product IDs match, time to save the Offer in the product
			if offerProduct.ProductID == product.ID {

				matched = true
				product.ApplyOffer(offerProduct)
			}

		}
		if !matched {
			offerProduct.OfferIssuesFound.IssueCode = "ProdctNotFound"
			offerProduct.OfferIssuesFound.IssueDescription = "The product was not found in the Catalogue to set the offer!!"
		}
	}

	return offers

}

// GetProducts - a Method to get a slice of Pointers to products in the Catalogue
func (c *Catalogue) GetProducts() []*datatypes.Product {
	return c.Products
}

// GetProduct - a Method to get a slice of Pointers to products in the Catalogue
func (c *Catalogue) GetProduct(p datatypes.Product) datatypes.Product {
	product := datatypes.Product{}
	for _, product := range c.Products {
		// product IDs match, time to save the Offer in the product
		if p.ID == product.ID {

			return *product
		}

	}
	return product
}

// String - a Stringer Interface implementation function to describe the type created.
func (c Catalogue) String() string {
	return fmt.Sprintf(" products in the catalogue %v", c.Products)
}
