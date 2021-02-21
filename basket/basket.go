package basket

import (
	"fmt"
	"io"

	"github.com/doyyan/ECS/calculator"
	"github.com/doyyan/ECS/catalogue"
	"github.com/doyyan/ECS/datatypes"
)

var found bool

// Basket is a sales basket which has a collection of items..
type Basket struct {
	// This is the items being purchased, there will be a Unique value for a product in each value of the item.
	Items []*datatypes.Item
	// Baskets have to be attached to a Catalogue to apply prices and discounts
	Catalouge catalogue.Catalogue
}

// Pricer is an interface that takes a Sale Basket from a customer, calculates the Totals
// Sub-Totals and Discounts and outputs that to the given Writer
type Pricer interface {
	Price(w io.Writer)
}

// NewBasket to create a new Basket and attach it to a Catalogue.
func NewBasket(c catalogue.Catalogue) Basket {

	return Basket{Catalouge: c}

}

// AddOrUpdateItem update the Items in a basket
func (b *Basket) AddOrUpdateItem(item *datatypes.Item) {
	found = false
	for _, val := range b.Items {
		if val.Product.ID == item.Product.ID {
			found = true
			val = item

		}

	}

	if !found {
		b.Items = append(b.Items, item)
	}

}

// Price - pricer implementation, returns a Sales Receipt and can output the data to the Writer that is being sent in.
func (b *Basket) Price(w io.Writer) datatypes.Receipt {

	grandTotal := calculator.Calculate(b.Items, b.Catalouge)

	r := datatypes.Receipt{}

	receipt := *r.CreateReceipt(b.Items, grandTotal)
	if w != nil {
		fmt.Println(w, receipt)
	}

	return receipt
}
