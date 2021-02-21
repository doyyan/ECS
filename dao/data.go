package dao

import (
	"github.com/doyyan/ECS/datatypes"
)

var (
	bakedBeansBuy2Get1Free datatypes.Offer
	sardinesQuarterOff     datatypes.Offer
	shampooGroupOffer      datatypes.Offer
	tooManyOffersOffer     datatypes.Offer
	testOffer              datatypes.Offer
	bakedBeans             datatypes.Product
	biscuits               datatypes.Product
	sardines               datatypes.Product
	shampooSmall           datatypes.Product
	shampooMedium          datatypes.Product
	shampooLarge           datatypes.Product
)

type datastruct struct {
	// Name of the product
	Name string
	// ID - a unique identifier for the product
	IssueDescription string
}

func init() {

	bakedBeansBuy2Get1Free = datatypes.Offer{ProductID: 1, ProductName: "Baked Beans", OfferName: "BakedBeansBuy2Get1Free", OfferID: 1, MultiBuyDiscount: datatypes.BuyXgetYFree{MinNoOfItems: 2, DiscountPercentage: 50}, GroupPromotion: datatypes.PromotionAsPartOfList{}}

	sardinesQuarterOff = datatypes.Offer{ProductID: 3, ProductName: "Sardines", OfferName: "SardinesQuarterOff", OfferID: 2, MultiBuyDiscount: datatypes.BuyXgetYFree{MinNoOfItems: 1, DiscountPercentage: 25}, GroupPromotion: datatypes.PromotionAsPartOfList{}}

	soapOfferItems := []datatypes.Item{{Product: datatypes.Product{
		Name: "Shampoo (Small)", ID: 4, BasicPrice: 2, Offer: datatypes.Offer{},
	}, NumberOfItems: 2}, {Product: datatypes.Product{
		Name: "Shampoo (Medium)", ID: 5, BasicPrice: 2.5, Offer: datatypes.Offer{},
	}, NumberOfItems: 1}, {Product: datatypes.Product{
		Name: "Shampoo (Large)", ID: 6, BasicPrice: 3.5, Offer: datatypes.Offer{},
	}, NumberOfItems: 3}}

	shampooGroupOffer = datatypes.Offer{ProductID: 4, ProductName: "Shampoo (Small)", OfferName: "ShampooGroupOffer", OfferID: 2, MultiBuyDiscount: datatypes.BuyXgetYFree{}, GroupPromotion: datatypes.PromotionAsPartOfList{ListOfProductToBePartOf: soapOfferItems, DiscountPercentage: 100}}
	tooManyOffersOffer = datatypes.Offer{ProductID: 3, ProductName: "Sardines", OfferName: "TooManyOffersOffer", OfferID: 2, MultiBuyDiscount: datatypes.BuyXgetYFree{MinNoOfItems: 1, DiscountPercentage: 25}, GroupPromotion: datatypes.PromotionAsPartOfList{ListOfProductToBePartOf: soapOfferItems, DiscountPercentage: 100}}

	bakedBeans = datatypes.Product{
		Name: "Baked Beans", ID: 1, BasicPrice: 0.99, Offer: datatypes.Offer{},
	}
	biscuits = datatypes.Product{
		Name: "Biscuits", ID: 2, BasicPrice: 1.20, Offer: datatypes.Offer{},
	}
	sardines = datatypes.Product{
		Name: "Sardines", ID: 3, BasicPrice: .89, Offer: datatypes.Offer{},
	}
	shampooSmall = datatypes.Product{
		Name: "Shampoo (Small)", ID: 4, BasicPrice: 2, Offer: datatypes.Offer{},
	}
	shampooMedium = datatypes.Product{
		Name: "Shampoo (Medium)", ID: 5, BasicPrice: 2.5, Offer: datatypes.Offer{},
	}
	shampooLarge = datatypes.Product{
		Name: "Shampoo (Large)", ID: 6, BasicPrice: 3.5, Offer: datatypes.Offer{},
	}

}

// GetOffers gets offers out to the consumer pretending to get it from persistant storage
func GetOffers() map[string]datatypes.Offer {
	return map[string]datatypes.Offer{bakedBeansBuy2Get1Free.OfferName: bakedBeansBuy2Get1Free,
		sardinesQuarterOff.OfferName: sardinesQuarterOff,
		shampooGroupOffer.OfferName:  shampooGroupOffer,
		tooManyOffersOffer.OfferName: tooManyOffersOffer}
}

// GetGoodOffers -- without the tooManyOffers offer that fails loading!!
func GetGoodOffers() map[string]datatypes.Offer {
	return map[string]datatypes.Offer{bakedBeansBuy2Get1Free.OfferName: bakedBeansBuy2Get1Free,
		sardinesQuarterOff.OfferName: sardinesQuarterOff,
		shampooGroupOffer.OfferName:  shampooGroupOffer}
}

// GetProducts gets products to the consumer pretending to be from persistant storage.
func GetProducts() map[string]datatypes.Product {
	var m = map[string]datatypes.Product{
		bakedBeans.Name:    bakedBeans,
		biscuits.Name:      biscuits,
		sardines.Name:      sardines,
		shampooSmall.Name:  shampooSmall,
		shampooMedium.Name: shampooMedium,
		shampooLarge.Name:  shampooLarge,
	}
	return m
}
