package datatypes

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
	ListOfProductToBePartOf []Item
	// The discount percentage for this Product if it participates in this 0ffer.
	DiscountPercentage int
}

// Item is a product and a number of those products chosen for purchase and put in a Shopping Basket.
type Item struct {
	// Prodcut that is in the basket.
	Product Product
	// NumberOfItems is the number of items that have been purchased.
	NumberOfItems int
}
