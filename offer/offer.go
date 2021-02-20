package offer

import (
	"fmt"

	"github.com/doyyan/ECS/item"
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

// BuyXgetYFree - to apply offers like Buy 1 get 1 free...
type BuyXgetYFree struct {
	// MinNoOfItems Minimum number of items in a batch to qualify for this discount..
	MinNoOfItems int
	// DiscountPercentage the percentage of discount.
	DiscountPercentage int
}

// Issue meaningful error codes/descriptions to be transported back..
type Issue struct {
	// IssueCodeThe Error code to be set
	IssueCode string
	// IssueDescriptiondescription of error.
	IssueDescription string
}

// PromotionAsPartOfList - if a group of items for which an offer applies..
type PromotionAsPartOfList struct {
	// ListOfProductToBePartOf  - the list of items, ALL of which have to be bought together to qualify for this offer..
	ListOfProductToBePartOf []item.Item
	// The discount percentage for this Product if it participates in this 0ffer.
	DiscountPercentage int
}

// CreateOffer a method to create an offer for a product..
func CreateOffer(productID int, productName string, offerName string, offerID int, multiBuyDiscount BuyXgetYFree, groupPromotion PromotionAsPartOfList) Offer {

	offer := Offer{}
	// Cannot be part of Two Promotions at the same time
	if (multiBuyDiscount != BuyXgetYFree{}) && groupPromotion.DiscountPercentage != 0 {
		offer.OfferName = offerName
		offer.OfferIssuesFound.IssueCode = "TooManyPromotions"
		offer.OfferIssuesFound.IssueDescription = "This promotion cannot be part of BuyXgetYFree and GroupPromotion"
		return offer

	}
	offer = Offer{ProductID: productID, ProductName: productName, OfferName: offerName, OfferID: offerID, MultiBuyDiscount: multiBuyDiscount, GroupPromotion: groupPromotion}
	return offer

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
