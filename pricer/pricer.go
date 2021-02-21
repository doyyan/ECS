package pricer

import (
	"io"
)

// Pricer is an interface that takes a Sale Basket from a customer, calculates the Totals
// Sub-Totals and Discounts and outputs that to the given Writer
type Pricer interface {
	Price(w io.Writer)
}
