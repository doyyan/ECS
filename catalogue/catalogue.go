package catalogue

/*
This package is to create and maintain a Catalogue of Products
*/

import (
	"fmt"

	"github.com/doyyan/ECS/offer"
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

// SetOffers - a Method to set the Offers for some/all Products in a Catalouge..
// Browse through the products and try and find the Product in each of the offer
// If found just set the offerID in the Product for now
// If not found return an error code as the team setting the Offer may be different
// to the team managing the Catalogue and its good to send a meaningful message back!!
func (c *Catalogue) SetOffers(offers []offer.Offer) []offer.Offer {

	for _, offerProduct := range offers {
		matched = false
		for _, product := range c.Products {

			if offerProduct.ProductID == product.ID {

				matched = true
				product.OfferID = offerProduct.OfferID
			}

		}
		if !matched {
			offerProduct.OfferIssuesFound.IssueCode = "ProdctNotFound"
			offerProduct.OfferIssuesFound.IssueDescription = "The product was not found in the Catalogue to set the offer!!"
		}
	}

	return offers

}

// String - a Stringer Interface implementation function to describe the type created.
func (c Catalogue) String() string {
	return fmt.Sprintf(" products in the catalogue ", c.Products)
}
