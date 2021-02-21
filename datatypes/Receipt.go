package datatypes

// Receipt the actual receipt of the Sales
type Receipt struct {
	ReceiptLines []Item
	GrandTotal   float32
}
