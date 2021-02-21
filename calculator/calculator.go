package calculator

import (
	"github.com/doyyan/ECS/catalogue"
	"github.com/doyyan/ECS/datatypes"
)

// Calculate the price of all the items in a Basket and returns a Grand Total value
func Calculate(Items []*datatypes.Item, Catalouge catalogue.Catalogue) float32 {

	// Range through the whole list of Items in the basket to calculate the total.
	for _, val := range Items {

		// First up check if the product exists in the first place!!
		product := Catalouge.GetProduct(val.Product)

		if product.ID != 0 {

			CalculateOffer(product, val)

		} else {
			val.BasicPrice = 0.0
			val.Discount = 0.0
			val.DiscountedBasicPrice = 0.0

		}

	}

	return 0.0

}

// CalculateOffer calculates any offer this product has
func CalculateOffer(p datatypes.Product, item *datatypes.Item) {

	if p.Offer.OfferID != 0 {

		if p.Offer.MultiBuyDiscount.DiscountPercentage != 0 {

			MultiBuyDiscountCalculator(p.Offer, item)

		}

	}

}

// MultiBuyDiscountCalculator -- this is to calculate offers like Buy 3 get 1 free..
func MultiBuyDiscountCalculator(offer datatypes.Offer, item *datatypes.Item) {
	// Only applies if they have purchased a Minimum number of items, for instance in a Buy 3 get 1 free, they must have purchased 3 or More!!
	if item.NumberOfItems >= offer.MultiBuyDiscount.MinNoOfItems {
		// This is the type of Offer like 25% off, that is EVERY item qualifies for the offer...
		if offer.MultiBuyDiscount.MinNoOfItems == 1 {
			item.BasicPrice = float32(item.NumberOfItems) * item.Product.BasicPrice

			percentageDiscount := float32(float32(offer.MultiBuyDiscount.DiscountPercentage) / float32(100.00))

			item.Discount = item.BasicPrice * float32(percentageDiscount)

			item.DiscountedBasicPrice = item.BasicPrice - float32(item.Discount)

		}

	}

}
