package offer

import "github.com/doyyan/ECS/datatypes"

// CreateOffer a method to create an offer for a product..
func CreateOffer(productID int, productName string, offerName string, offerID int, multiBuyDiscount datatypes.BuyXgetYFree, groupPromotion datatypes.PromotionAsPartOfList) datatypes.Offer {

	offer := datatypes.Offer{}
	// Cannot be part of Two Promotions at the same time
	if (multiBuyDiscount != datatypes.BuyXgetYFree{}) && groupPromotion.DiscountPercentage != 0 {
		offer.OfferName = offerName
		offer.OfferIssuesFound.IssueCode = "TooManyPromotions"
		offer.OfferIssuesFound.IssueDescription = "This promotion cannot be part of BuyXgetYFree and GroupPromotion"
		return offer

	}
	offer = datatypes.Offer{ProductID: productID, ProductName: productName, OfferName: offerName, OfferID: offerID, MultiBuyDiscount: multiBuyDiscount, GroupPromotion: groupPromotion}
	return offer

}
