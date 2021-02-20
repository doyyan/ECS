package item

/*

Package to define an Item, which is essently a product chosen for purchase and dropped into a Basket in a superstore!

*/
import "github.com/doyyan/ECS/product"

// Item is a product and a number of those products chosen for purchase and put in a Shopping Basket.
type Item struct {
	// Prodcut that is in the basket.
	Product product.Product
	// NumberOfItems is the number of items that have been purchased.
	NumberOfItems int
}
