package datatypes

import (
	"fmt"
)

// Offer is the main struct for Offer teams to create Offers for a SINGLE product
type Offer struct {
	// ProductID - the ID of the product
	ProductID int
	// ProductName the name of the product
	ProductName string
	// OfferName name of the offer
	OfferName string
	// OfferID  ID of the offer
	OfferID int
	// MultiBuyDiscount - multibuy discount to apply
	MultiBuyDiscount BuyXgetYFree
	// OfferErrorsFound errors found during offer application
	OfferIssuesFound Issue
	// Group Promotion offer, if the current product is part of a group of products participating in a promotion, like 'buy 2 large soaps and get a small soap free...'
	GroupPromotion PromotionAsPartOfList
}

// CreateOffer a method to create an offer for a product..
func CreateOffer(productID int, productName string, offerName string, offerID int, multiBuyDiscount BuyXgetYFree, groupPromotion PromotionAsPartOfList) Offer {
	o := Offer{}

	// Cannot be part of Two Promotions at the same time
	if (multiBuyDiscount != BuyXgetYFree{}) && groupPromotion.DiscountPercentage != 0 {
		o.OfferName = offerName
		o.OfferIssuesFound.IssueCode = "TooManyPromotions"
		o.OfferIssuesFound.IssueDescription = "This promotion cannot be part of BuyXgetYFree and GroupPromotion"
		return o

	}
	o = Offer{ProductID: productID, ProductName: productName, OfferName: offerName, OfferID: offerID, MultiBuyDiscount: multiBuyDiscount, GroupPromotion: groupPromotion}
	return o

}

// String - a Stringer Interface implementation function to describe the type created.
func (o Offer) String() string {
	offerDesc := fmt.Sprintf("Product %v has an offer named %v", o.ProductName, o.OfferName)
	if (o.MultiBuyDiscount != BuyXgetYFree{}) {
		offerDesc = offerDesc + fmt.Sprintf(" It has a Multibuy discount with Minimum no of items to buy as %v and discount percentage %v", o.MultiBuyDiscount.MinNoOfItems, o.MultiBuyDiscount.DiscountPercentage)

	}
	if o.GroupPromotion.DiscountPercentage != 0 {
		offerDesc = offerDesc + fmt.Sprintf(" It has a Group Buy discount with items %v and a discount of %v", o.GroupPromotion.ListOfProductToBePartOf, o.GroupPromotion.DiscountPercentage)

	}

	return offerDesc
}
