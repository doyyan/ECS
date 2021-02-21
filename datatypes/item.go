package datatypes

/*

Package to define an Item, which is essently a product chosen for purchase and dropped into a Basket in a superstore!

*/

// Item is a product and a number of those products chosen for purchase and put in a Shopping Basket.
type Item struct {
	// Prodcut that is in the basket.
	Product Product
	// NumberOfItems is the number of items that have been purchased.
	NumberOfItems int
	// Discount the total discount applied to this item.
	Discount float32
	// Offer - the offer that was applied to this item
	// BasicPrice - the basic price of a product
	BasicPrice float32

	// Dscounted Price or Subtotal - the basic price of a product
	DiscountedBasicPrice float32
}

// CreateItem creates an instance of an item for a Product with a given number of items
func CreateItem(p Product, i int) Item {
	return Item{Product: p, NumberOfItems: i}

}
