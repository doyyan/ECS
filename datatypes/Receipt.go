package datatypes

import (
	"fmt"
)

const (
	FLOATFORMAT = "%8.2f"
)

// Receipt the actual receipt of the Sales
type Receipt struct {
	ReceiptLines []*Item
	GrandTotal   float32
}

// String - a Stringer Interface implementation function to describe the type created.
func (r Receipt) String() string {

	data := "\n ===========  Receipt of Sales ================\n"

	total := r.GrandTotal

	if len(r.ReceiptLines) > 0 {
		for _, val := range r.ReceiptLines {
			data = data + val.Product.Name + " \t Original Price " + fmt.Sprintf(FLOATFORMAT, val.BasicPrice) + " \t Discount " + fmt.Sprintf(FLOATFORMAT, val.Discount) + " \t Discounted Price " + fmt.Sprintf(FLOATFORMAT, val.DiscountedBasicPrice) + "\n"
			if val.DiscountedBasicPrice != 0 {
				total = total + val.DiscountedBasicPrice
			} else {
				total = total + val.BasicPrice
			}

		}

	}
	data = data + "\n ========================================================================="
	data = data + fmt.Sprintf("\n\t\tGrand Total\t "+FLOATFORMAT, total)

	return data

}

func (r *Receipt) CreateReceipt(receiptLines []*Item, grandTotal float32) *Receipt {

	r.ReceiptLines = receiptLines
	r.GrandTotal = grandTotal

	return r

}
