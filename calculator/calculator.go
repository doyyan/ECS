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
			// Its a valid product, lets calculate the price
			CalculateOffer(product, val)

		} else {

			// The ref data can't be found, lets Zero them all and let the Viewer find this to be odd and work out where the problem is!!
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
			//Calculate and return
			return

		}

	}
	// No applicable Offers, full price calculation
	item.BasicPrice = float32(item.NumberOfItems) * item.Product.BasicPrice

}

// MultiBuyDiscountCalculator -- this is to calculate offers like Buy 3 get 1 free..
func MultiBuyDiscountCalculator(offer datatypes.Offer, item *datatypes.Item) {
	// Only applies if they have purchased a Minimum number of items, for instance in a Buy 3 get 1 free, they must have purchased 3 or More!!
	if item.NumberOfItems >= offer.MultiBuyDiscount.MinNoOfItems {
		// This is the type of Offer like 25% off, that is EVERY item qualifies for the offer...
		if offer.MultiBuyDiscount.MinNoOfItems == 1 {

			// Now calcualte the Discounted Percentage for the items
			percentageDiscount := float32(float32(offer.MultiBuyDiscount.DiscountPercentage) / float32(100.00))

			// Now calculate what is the percentage taken Out by the discount
			item.Discount = item.BasicPrice * float32(percentageDiscount)
			// Finally the difference between what would have been the original price without discount and the discounted price.
			item.DiscountedBasicPrice = item.BasicPrice - float32(item.Discount)

		} else if offer.MultiBuyDiscount.MinNoOfItems > 1 {

			MultiBuyBuyXgetYFreeCalculator(offer, item)

		}

	} else {
		item.BasicPrice = float32(item.NumberOfItems) * item.Product.BasicPrice

	}

}

// MultiBuyBuyXgetYFreeCalculator this calculator is ONLY for the Buy 3 get 1 free type of calculations
// Lets say buy 3 get 1 free and I buy 11 items. The logic goes like this, do a 11 MOD 3, that gives as 2 as reminder
// Customer pays full price for those 2.
// All the Remaining items are discounted by 33.5% in this case, which has to be set in the DiscountPercentage!!
func MultiBuyBuyXgetYFreeCalculator(offer datatypes.Offer, item *datatypes.Item) {
	// Find the price without any Discount for this item
	item.BasicPrice = float32(item.NumberOfItems) * item.Product.BasicPrice
	// Find the MODULUS of the
	definiteFullPriceItems := item.NumberOfItems % offer.MultiBuyDiscount.MinNoOfItems
	// These are the items definitely paying full price!!
	fullPricedItems := float32(definiteFullPriceItems) * item.Product.BasicPrice

	// Now for the remaining items (if Any, for instance if a customer bought only 2 items in a Buy 3 get 1 free, however silly that might be)
	actualDiscountedItems := item.NumberOfItems - definiteFullPriceItems

	// Now calcualte the Discounted Percentage for the items
	percentageDiscount := float32((100 - float32(offer.MultiBuyDiscount.DiscountPercentage)) / float32(100.00))

	// Now calculated the discounted price of the discounted items
	totalDiscounts := float32(actualDiscountedItems) * float32(percentageDiscount) * item.Product.BasicPrice

	// Now calculate what is the percentage taken Out by the discount
	item.Discount = item.BasicPrice - (totalDiscounts + fullPricedItems)
	// Finally the difference between what would have been the original price without discount and the discounted price.
	item.DiscountedBasicPrice = totalDiscounts + fullPricedItems

}
