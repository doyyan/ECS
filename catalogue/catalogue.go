package catalogue

type catalogue struct {
}

func GetCatalogue() catalogue {
	return catalogue{}
}

func (catalog catalogue) String() string {
	return "Catalogue"
}
