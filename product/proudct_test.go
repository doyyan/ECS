package product_test

import (
	"testing"

	"github.com/doyyan/ECS/offer"
	"github.com/doyyan/ECS/product"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

var BasicTestProduct datastruct
var testProduct product.Product

type datastruct struct {
	// Name of the product
	Name string
	// ID - a unique identifier for the product
	ID int

	// BasicPrice - the basic price of a product
	BasicPrice float32
	// Offer - any offer available for this product
	OfferAvailable offer.Offer
}

func init() {
	BasicTestProduct = datastruct{
		"Baked Beans", 1, 0.99, offer.Offer{},
	}
}

// TestParallelize- paraellized tests (great for CI/CD builds) that run and test happy and sad paths!!

func TestParallelize(t *testing.T) {

	type tableTest struct {
		name       string
		datainput  datastruct
		dataoutput datastruct
	}

	tests := []tableTest{
		{"statusok", BasicTestProduct, datastruct{"Baked Beans", 1, 0.99, offer.Offer{}}},
		//{"statusnotok", "x", "10", errorValue, http.StatusInternalServerError},
	}

	t.Log("Given the need to test if Products can be created successfully.")
	{
		for _, test := range tests {
			t.Logf("When checking to see if a product has been successfully created for product with data %v %v %v", test.datainput.Name, test.datainput.ID, test.datainput.BasicPrice)

			tf := func(t *testing.T) {
				t.Parallel()

				{
					product := product.CreateProduct(test.datainput.Name, test.datainput.ID, test.datainput.BasicPrice, test.datainput.OfferAvailable)

					t.Log("A product Should be Successfully created ")
					if product.Name == test.datainput.Name {
						t.Logf(" Products match inputdata: %v and output: %v", test.datainput, product)
					} else {
						t.Errorf("Expected a product with data %v but got %v", test.datainput, product)
					}
				}
			}
			t.Run(test.name, tf)
		}
	}
}

// BenchmarkFib20 -- test a small nominal value for Benchmarks
func BenchmarkFCreateProduct(b *testing.B) {
	b.ResetTimer()

	test := BasicTestProduct
	for n := 0; n < b.N; n++ {
		product := product.CreateProduct(test.Name, test.ID, test.BasicPrice, test.OfferAvailable)
		testProduct = product

	}
}
