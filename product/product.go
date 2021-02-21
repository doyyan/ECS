package product

/* Package product defines Types, Functions and Methods to create and maintain products available in a Superstore(for instance)

 */
import (
	"math"

	"github.com/doyyan/ECS/datatypes"
)

// CreateProduct a function to create and return a Product
func CreateProduct(name string, ID int, basicPrice float32, offer datatypes.Offer) datatypes.Product {
	// Rounding to a Ceiling value of the incoming Float and saving it onto Float32 to save memory (as opposed to Float 64)
	return datatypes.Product{Name: name, ID: ID, BasicPrice: float32(math.Ceil(float64(basicPrice)*100) / 100), Offer: offer}
}
