package product_test

import (
	"testing"

	"github.com/doyyan/ECS/product"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

var (
	BakedBeans    product.Product
	Biscuits      product.Product
	Sardines      product.Product
	ShampooSmall  product.Product
	ShampooMedium product.Product
	ShampooLarge  product.Product
	testProduct   product.Product
)

type datastruct struct {
	// Name of the product
	Name string
	// ID - a unique identifier for the product
	ID int

	// BasicPrice - the basic price of a product
	BasicPrice float32
	// OfferID - any offer available for this product
	OfferID int
}

func init() {
	BakedBeans = product.Product{
		"Baked Beans", 1, 0.99, 0,
	}
	Biscuits = product.Product{
		"Biscuits", 2, 1.20, 0,
	}
	Sardines = product.Product{
		"Sardines", 3, 1.89, 0,
	}
	ShampooSmall = product.Product{
		"Shampoo (Small)", 4, 2, 0,
	}
	ShampooMedium = product.Product{
		"Shampoo (Medium)", 5, 2.5, 0,
	}
	ShampooLarge = product.Product{
		"Shampoo (Large)", 6, 3.5, 0,
	}
}

// TestParallelize- paraellized tests (great for CI/CD builds) that run and test happy and sad paths!!

func TestParallelize(t *testing.T) {

	type tableTest struct {
		name       string
		datainput  product.Product
		dataoutput datastruct
	}

	tests := []tableTest{
		{"statusok", BakedBeans, datastruct{"Baked Beans", 1, 0.99, 0}},
		{"statusok", Biscuits, datastruct{"Biscuits", 2, 1.20, 0}},
		{"statusok", Sardines, datastruct{"Sardines", 3, 1.89, 0}},
		{"statusok", ShampooSmall, datastruct{"Shampoo (Small)", 4, 2, 0}},
		{"statusok", ShampooMedium, datastruct{"Shampoo (Medium)", 5, 2.5, 0}},
		{"statusok", ShampooLarge, datastruct{"Shampoo (Large)", 6, 3.5, 0}},
		//{"statusnotok", "x", "10", errorValue, http.StatusInternalServerError},
	}

	t.Log("Given the need to test if Products can be created successfully.")
	{
		for _, test := range tests {
			t.Logf("When checking to see if a product has been successfully created for product with data %v %v %v", test.datainput.Name, test.datainput.ID, test.datainput.BasicPrice)

			tf := func(t *testing.T) {
				t.Parallel()

				{
					product := product.CreateProduct(test.datainput.Name, test.datainput.ID, test.datainput.BasicPrice, 0)

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

	test := Biscuits
	for n := 0; n < b.N; n++ {
		product := product.CreateProduct(test.Name, test.ID, test.BasicPrice, 0)
		testProduct = product

	}
}
