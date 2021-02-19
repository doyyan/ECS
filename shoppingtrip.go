package main

import (
	"fmt"

	"github.com/doyyan/ECS/catalogue"
)

func main() {
	fmt.Println(" Hello world...")
	catalog := catalogue.GetCatalogue()
	fmt.Println(catalog)

}
