package prototype

import (
	"errors"
	"fmt"
)

const (
	White  = 1
	Black  = 2
	Yellow = 3
)

var (
	whitePrototype  *Shirt = &Shirt{Price: 15.00, SKU: "empty", Color: White}
	blackPrototype  *Shirt = &Shirt{Price: 16.00, SKU: "empty", Color: Black}
	yellowPrototype *Shirt = &Shirt{Price: 17.00, SKU: "empty", Color: Yellow}
)

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtCloner interface {
	GetClone(color int) (ItemInfoGetter, error)
}

var shirtCache *ShirtCache

func GetShirtsCloner() ShirtCloner {
	if shirtCache == nil {
		shirtCache = &ShirtCache{}
	}
	return shirtCache
}

type ShirtCache struct {
}

func (*ShirtCache) GetClone(color int) (ItemInfoGetter, error) {
	switch color {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Yellow:
		newItem := *yellowPrototype
		return &newItem, nil
	default:
		return nil, errors.New("not implemented yet")
	}
}

type ShirtColor byte
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and color id %d costs %0.2f", s.SKU, s.Color, s.Price)
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}
