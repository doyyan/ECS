package datatypes_test

import (
	"testing"

	"github.com/doyyan/ECS/dao"
	"github.com/doyyan/ECS/datatypes"
	"github.com/doyyan/ECS/product"
)

var (
	testProduct datatypes.Product
	products    map[string]datatypes.Product
)

func init() {
	products = dao.GetProducts()
	offers = dao.GetOffers()
}

type datastruct struct {
	// Name of the product
	Name string
	// ID - a unique identifier for the product
	ID int

	// BasicPrice - the basic price of a product
	BasicPrice float32
	// OfferID - any offer available for this product
	Offer datatypes.Offer
}

// TestParallelize- paraellized tests (great for CI/CD builds) that run and test happy and sad paths!!

func TestCreateProducts(t *testing.T) {

	type tableTest struct {
		name       string
		datainput  datatypes.Product
		dataoutput datastruct
	}

	tests := []tableTest{
		{"statusok", products["Baked Beans"], datastruct{"Baked Beans", 1, 0.99, datatypes.Offer{}}},
		{"statusok", products["Biscuits"], datastruct{"Biscuits", 2, 1.20, datatypes.Offer{}}},
		{"statusok", products["Sardines"], datastruct{"Sardines", 3, 1.89, datatypes.Offer{}}},
		{"statusok", products["ShampooSmall"], datastruct{"Shampoo (Small)", 4, 2, datatypes.Offer{}}},
		{"statusok", products["ShampooMedium"], datastruct{"Shampoo (Medium)", 5, 2.5, datatypes.Offer{}}},
		{"statusok", products["ShampooLarge"], datastruct{"Shampoo (Large)", 6, 3.5, datatypes.Offer{}}},
		//{"statusnotok", "x", "10", errorValue, http.StatusInternalServerError},
	}

	t.Log("Given the need to test if Products can be created successfully.")
	{
		for _, test := range tests {
			t.Logf("When checking to see if a product has been successfully created for product with data %v %v %v", test.datainput.Name, test.datainput.ID, test.datainput.BasicPrice)

			tf := func(t *testing.T) {
				t.Parallel()

				{
					product := product.CreateProduct(test.datainput.Name, test.datainput.ID, test.datainput.BasicPrice, datatypes.Offer{})

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

func TestCreateProductsWithOffer(t *testing.T) {

	type tableTest struct {
		name       string
		datainput  datatypes.Product
		dataoutput datastruct
	}

	tests := []tableTest{
		{"statusok", products["Baked Beans"], datastruct{"Baked Beans", 1, 0.99, datatypes.Offer{}}},

		//{"statusnotok", "x", "10", errorValue, http.StatusInternalServerError},
	}

	t.Log("Given the need to test if Products can be created successfully.")
	{
		for _, test := range tests {
			t.Logf("When checking to see if a product has been successfully created for product with data %v %v %v", test.datainput.Name, test.datainput.ID, test.datainput.BasicPrice)

			tf := func(t *testing.T) {
				t.Parallel()

				{
					testOffer := offers["BakedBeansBuy2Get1Free"]
					offer := datatypes.CreateOffer(testOffer.ProductID, testOffer.ProductName, testOffer.OfferName, testOffer.OfferID, testOffer.MultiBuyDiscount, testOffer.GroupPromotion)
					product := product.CreateProduct(test.datainput.Name, test.datainput.ID, test.datainput.BasicPrice, offer)
					product.ApplyOffer(offer)

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

	test := products["Biscuits"]
	for n := 0; n < b.N; n++ {
		product := product.CreateProduct(test.Name, test.ID, test.BasicPrice, datatypes.Offer{})
		testProduct = product

	}
}
