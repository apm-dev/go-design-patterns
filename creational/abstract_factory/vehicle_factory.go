package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	BmwFactoryType  = 1
	FordFactoryType = 2
)

type VehicleFactory interface {
	NewVehicle(t int) (Vehicle, error)
}

func GetVehicleFactory(brand int) (VehicleFactory, error) {
	switch brand {
	case BmwFactoryType:
		return new(BMWFactory), nil
	case FordFactoryType:
		return new(FordFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Brand %d factory not exists", brand))
	}
}
