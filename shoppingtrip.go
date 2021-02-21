package main

import (
	"fmt"

	"github.com/doyyan/ECS/basket"
	"github.com/doyyan/ECS/catalogue"
	"github.com/doyyan/ECS/dao"
	"github.com/doyyan/ECS/datatypes"
)

var (
	testCatalogue catalogue.Catalogue
	productsdata  map[string]*datatypes.Product
	offersdata    map[string]datatypes.Offer
	produce       datatypes.Product
	produce2      datatypes.Product
)

func init() {
	productsdata = dao.GetProductsData()
	offersdata = dao.GetGoodOffers()

}

func main() {

	products := []*datatypes.Product{}

	for _, value := range productsdata {
		products = append(products, value)
		if value.Name == "Sardines" {
			produce = *value
		}
	}
	offers := []datatypes.Offer{}
	for _, offerValue := range offersdata {
		offers = append(offers, offerValue)
	}
	catalogue := catalogue.NewCatalogue("MegaStore", 1, products)
	catalogue.SetOffers(offers)

	basket := basket.NewBasket(catalogue)

	for _, value := range products {

		if value.Name == "Baked Beans" {
			produce = *value
		}
		if value.Name == "Biscuits" {
			produce2 = *value
		}
	}

	// Baked Beans have a Buy 3 Get 1 Free offer. Change the value below here to test."
	item := datatypes.CreateItem(produce, 5)
	// Buscuits are priced at £1.20 with No Discounts. Change the value below here to test.")
	item2 := datatypes.CreateItem(produce2, 14)

	basket.AddOrUpdateItem(&item2)

	basket.AddOrUpdateItem(&item)
	receipt := basket.Price(nil)
	fmt.Println(receipt)

}
