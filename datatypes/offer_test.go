package datatypes_test

import (
	"testing"

	"github.com/doyyan/ECS/dao"
	"github.com/doyyan/ECS/datatypes"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

var (
	offers    map[string]datatypes.Offer
	testOffer datatypes.Offer
)

type datastructOffer struct {
	// Name of the product
	Name string
	// ID - a unique identifier for the product
	IssueDescription string
}

func init() {

	offers = dao.GetOffers()

}

// TestParallelize- paraellized tests (great for CI/CD builds) that run and test happy and sad paths!!

func TestOfferParallelize(t *testing.T) {
	//CreateOffer(productID int, productName string, offerName string, offerID int, multiBuyDiscount BuyXgetYFree, groupPromotion PromotionAsPartOfList)

	type tableTest struct {
		name       string
		datainput  datatypes.Offer
		dataoutput datastructOffer
	}

	tests := []tableTest{
		{"BakedBeansBuy2Get1FreeOK", offers["BakedBeansBuy2Get1Free"], datastructOffer{"Baked Beans", ""}},
		{"SardinesQuarterOffOK", offers["SardinesQuarterOff"], datastructOffer{"Sardines", ""}},
		{"ShampooGroupOfferOK", offers["ShampooGroupOffer"], datastructOffer{"Shampoo (Small)", ""}},
		{"TooManyOffersOfferNotOk", offers["TooManyOffersOffer"], datastructOffer{"TooManyOffersOffer", "TooManyPromotions"}},
	}

	t.Log("Given the need to test if Offers can be created successfully.")
	{
		for _, test := range tests {
			t.Logf("When checking to see if an offer has been successfully created for offer with data %v", test.datainput.OfferName)

			tf := func(t *testing.T) {
				//t.Parallel()

				{
					o := datatypes.Offer{}
					offer := o.CreateOffer(test.datainput.ProductID, test.datainput.ProductName, test.datainput.OfferName, test.datainput.OfferID, test.datainput.MultiBuyDiscount, test.datainput.GroupPromotion)

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
func BenchmarkFCreateOffer(b *testing.B) {
	b.ResetTimer()

	BakedBeansBuy2Get1Free := offers["BakedBeansBuy2Get1Free"]
	o := datatypes.Offer{}

	for n := 0; n < b.N; n++ {
		offer := o.CreateOffer(BakedBeansBuy2Get1Free.ProductID, BakedBeansBuy2Get1Free.ProductName, BakedBeansBuy2Get1Free.OfferName, BakedBeansBuy2Get1Free.OfferID, BakedBeansBuy2Get1Free.MultiBuyDiscount, BakedBeansBuy2Get1Free.GroupPromotion)
		testOffer = offer

	}
}
