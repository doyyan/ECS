package catalogue_test

import (
	"testing"

	"github.com/doyyan/ECS/catalogue"
	"github.com/doyyan/ECS/dao"
	"github.com/doyyan/ECS/datatypes"
)

var (
	testCatalogue catalogue.Catalogue
	productsdata  map[string]datatypes.Product
	offersdata    map[string]datatypes.Offer
	products      []datatypes.Product
	offers        []datatypes.Offer
)

func init() {
	productsdata = dao.GetProducts()
	offersdata = dao.GetGoodOffers()

}

type datastruct struct {
	// Name of the product
	Name string
	// ID - a unique identifier for the product
	ID int
}

// TestParallelize- paraellized tests (great for CI/CD builds) that run and test happy and sad paths!!

func TestCreateCatalogue(t *testing.T) {

	type tableTest struct {
		name       string
		dataoutput datastruct
	}

	tests := []tableTest{
		{"statusok", datastruct{Name: "MegaStore", ID: 1}},

		//{"statusnotok", "x", "10", errorValue, http.StatusInternalServerError},
	}

	t.Log("Given the need to test if Catalogues can be created successfully.")
	{
		for _, test := range tests {
			products := []*datatypes.Product{}
			for _, value := range productsdata {
				products = append(products, &value)
			}
			offers := []datatypes.Offer{}
			for _, offerValue := range offersdata {
				offers = append(offers, offerValue)
			}
			t.Logf("When checking to see if a catalogue has been successfully created  with name %v  and number of products %v ", "Megastore", len(products))

			tf := func(t *testing.T) {
				t.Parallel()

				{

					catalogue := catalogue.NewCatalogue("MegaStore", 1, products)

					catalogue.SetOffers(offers)

					t.Log("A product Should be Successfully created ")
					if catalogue.Name == test.dataoutput.Name {
						t.Logf(" Catalogue data match inputdata: %v and output: %v number of products expected was %v and got %v ", "Megastore", catalogue.Name, len(products), len(catalogue.GetProducts()))
					} else {
						t.Errorf("Expected a catalogue with data %v but got %v and number of products expected was %v but got %v ", "Megastore", catalogue.Name, len(products), len(catalogue.GetProducts()))
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

	//for n := 0; n < b.N; n++ {
	//		catalogue := catalogue.NewCatalogue("MegaStore", 1, &products)
	//	catalogue.SetOffers(offers)
	//	testCatalogue = catalogue

	//	}
}
