package abstract_factory

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	BmwXSeriesType = 1
	BmwISeriesType = 2
)

type BMWFactory struct{}

func (*BMWFactory) NewVehicle(series int) (Vehicle, error) {
	switch series {
	case BmwISeriesType:
		return &BmwISeries{
			model: []string{"528", "630", "720"}[rand.Intn(3)],
		}, nil
	case BmwXSeriesType:
		return &BmwXSeries{
			model: []string{"3", "4", "6"}[rand.Intn(3)],
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Series %d not exists", series))
	}
}
