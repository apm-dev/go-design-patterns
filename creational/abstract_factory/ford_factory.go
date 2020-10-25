package abstract_factory

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	FordSedanType = 1
	FordSuvType   = 2
)

type FordFactory struct{}

func (*FordFactory) NewVehicle(t int) (Vehicle, error) {
	switch t {
	case FordSedanType:
		return &FordSedanSeries{
			model: []string{"Figo", "Fiesta", "Mondeo"}[rand.Intn(3)],
		}, nil
	case FordSuvType:
		return &FordSuvSeries{
			model: []string{"Edge", "Expeditoin", "Explorer"}[rand.Intn(3)],
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Ford type %d not exists", t))
	}
}
