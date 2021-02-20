package product

/* Package product defines Types, Functions and Methods to create and maintain products available in a Superstore(for instance)

 */
import (
	"fmt"
	"math"
)

// Product is a generic struct to hold product info.
type Product struct {
	// Name of the product
	Name string
	// ID - a unique identifier for the product
	ID int

	// BasicPrice - the basic price of a product
	BasicPrice float32
	// Offer - any offer available for this product
	OfferID int
}

// CreateProduct a function to create and return a Product
func CreateProduct(name string, ID int, basicPrice float32, offerID int) Product {
	// Rounding to a Ceiling value of the incoming Float and saving it onto Float32 to save memory (as opposed to Float 64)
	return Product{name, ID, float32(math.Ceil(float64(basicPrice)*100) / 100), offerID}
}

// String - a Stringer Interface implementation function to describe the type created.
func (p Product) String() string {
	return fmt.Sprintf("%v and price %v and offers %v", p.Name, p.BasicPrice, 0)
}
