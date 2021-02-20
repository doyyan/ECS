package dao

import (
	"github.com/doyyan/ECS/item"
	"github.com/doyyan/ECS/offer"
	"github.com/doyyan/ECS/product"
)

var (
	BakedBeansBuy2Get1Free offer.Offer
	SardinesQuarterOff     offer.Offer
	ShampooGroupOffer      offer.Offer
	TooManyOffersOffer     offer.Offer
	testOffer              offer.Offer
	BakedBeans             product.Product
	Biscuits               product.Product
	Sardines               product.Product
	ShampooSmall           product.Product
	ShampooMedium          product.Product
	ShampooLarge           product.Product
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

func GetOffers() []offer.Offer {
	return []offer.Offer{BakedBeansBuy2Get1Free,
		SardinesQuarterOff,
		ShampooGroupOffer,
		TooManyOffersOffer}
}

func GetProducts() map[string]product.Product {
	var m = map[string]product.Product{
		Biscuits.Name:      Biscuits,
		Sardines.Name:      Sardines,
		ShampooSmall.Name:  ShampooSmall,
		ShampooMedium.Name: ShampooMedium,
		ShampooLarge.Name:  ShampooLarge,
	}
	return m
}
