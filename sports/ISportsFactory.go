package sports

import "fmt"

type ISportsFactory interface {
	MakeShirt() IShirt
	MakeShoe() IShoe
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}

	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed: %q", brand)
}
