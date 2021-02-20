package product

/* Package product defines Types, Functions and Methods to create and maintain products available in a Superstore(for instance)

 */
import (
	"fmt"
	"math"

	"github.com/doyyan/ECS/offer"
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
	OfferAvailable offer.Offer
}

// CreateProduct a function to create and return a Product
func CreateProduct(name string, ID int, BasicPrice float32, OfferAvailable offer.Offer) Product {
	// Rounding to a Ceiling value of the incoming Float and saving it onto Float32 to save memory (as opposed to Float 64)
	return Product{name, ID, float32(math.Ceil(float64(BasicPrice)*100) / 100), OfferAvailable}
}

// String - a Stringer Interface implementation function to describe the type created.
func (p Product) String() string {
	return fmt.Sprintf("%v and price %v and offers %v", p.Name, p.BasicPrice, p.OfferAvailable)
}
