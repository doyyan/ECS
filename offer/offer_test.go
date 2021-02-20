package offer_test

import (
	"testing"

	"github.com/doyyan/ECS/item"
	"github.com/doyyan/ECS/offer"
	"github.com/doyyan/ECS/product"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

var (
	BakedBeans             product.Product
	Biscuits               product.Product
	Sardines               product.Product
	ShampooSmall           product.Product
	ShampooMedium          product.Product
	ShampooLarge           product.Product
	testProduct            product.Product
	BakedBeansBuy2Get1Free offer.Offer
	SardinesQuarterOff     offer.Offer
	ShampooGroupOffer      offer.Offer
	TooManyOffersOffer     offer.Offer
)

type datastruct struct {
	// Name of the product
	Name string
	// ID - a unique identifier for the product
	IssueDescription string
}

func init() {

	BakedBeansBuy2Get1Free = offer.Offer{ProductID: 1, ProductName: "Baked Beans", OfferName: "BakedBeansBuy2Get1Free", OfferID: 1, MultiBuyDiscount: offer.BuyXgetYFree{MinNoOfItems: 2, DiscountPercentage: 50}, GroupPromotion: offer.PromotionAsPartOfList{}}

	SardinesQuarterOff = offer.Offer{ProductID: 2, ProductName: "Sardines", OfferName: "SardinesQuarterOff", OfferID: 2, MultiBuyDiscount: offer.BuyXgetYFree{MinNoOfItems: 1, DiscountPercentage: 25}, GroupPromotion: offer.PromotionAsPartOfList{}}

	soapOfferItems := []item.Item{{product.Product{
		"Shampoo (Small)", 4, 2, 0,
	}, 2}, {product.Product{
		"Shampoo (Medium)", 5, 2.5, 0,
	}, 1}, {product.Product{
		"Shampoo (Large)", 6, 3.5, 0,
	}, 3}}

	ShampooGroupOffer = offer.Offer{ProductID: 4, ProductName: "Shampoo (Small)", OfferName: "ShampooGroupOffer", OfferID: 2, MultiBuyDiscount: offer.BuyXgetYFree{}, GroupPromotion: offer.PromotionAsPartOfList{soapOfferItems, 100}}
	TooManyOffersOffer = offer.Offer{ProductID: 7, ProductName: "Sardines", OfferName: "TooManyOffersOffer", OfferID: 2, MultiBuyDiscount: offer.BuyXgetYFree{MinNoOfItems: 1, DiscountPercentage: 25}, GroupPromotion: offer.PromotionAsPartOfList{soapOfferItems, 100}}

}

// TestParallelize- paraellized tests (great for CI/CD builds) that run and test happy and sad paths!!

func TestParallelize(t *testing.T) {
	//CreateOffer(productID int, productName string, offerName string, offerID int, multiBuyDiscount BuyXgetYFree, groupPromotion PromotionAsPartOfList)

	type tableTest struct {
		name       string
		datainput  offer.Offer
		dataoutput datastruct
	}

	tests := []tableTest{
		//	{"BakedBeansBuy2Get1FreeOK", BakedBeansBuy2Get1Free, datastruct{"Baked Beans", ""}},
		//{"SardinesQuarterOffOK", SardinesQuarterOff, datastruct{"Sardines", ""}},
		//	{"ShampooGroupOfferOK", ShampooGroupOffer, datastruct{"Shampoo (Small)", ""}},
		{"TooManyOffersOfferNotOk", TooManyOffersOffer, datastruct{"TooManyOffersOffer", "TooManyPromotions"}},
	}

	t.Log("Given the need to test if Offers can be created successfully.")
	{
		for _, test := range tests {
			t.Logf("When checking to see if an offer has been successfully created for offer with data %v", test.datainput.OfferName)

			tf := func(t *testing.T) {
				t.Parallel()

				{
					offer := offer.CreateOffer(test.datainput.ProductID, test.datainput.ProductName, test.datainput.OfferName, test.datainput.OfferID, test.datainput.MultiBuyDiscount, test.datainput.GroupPromotion)

					t.Log("An offer Should be Successfully created ")
					if offer.OfferName == test.datainput.OfferName && offer.OfferIssuesFound.IssueCode == test.dataoutput.IssueDescription {
						t.Logf(" Offers match inputdata: %v and output: %v", test.datainput, offer)
					} else {
						t.Errorf("Expected a offer with data %v but got %v", offer.OfferIssuesFound.IssueCode, test.dataoutput.IssueDescription)
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
