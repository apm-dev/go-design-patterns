package abstract_factory

import (
	"fmt"
	"testing"
)

func TestBmwFactory(t *testing.T) {
	bmwFactory, err := GetVehicleFactory(BmwFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	bmw1, err := bmwFactory.NewVehicle(BmwISeriesType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Vehicle is from %s company", bmw1.GetManufacturer())
	bmwISeries, ok := bmw1.(*BmwISeries)
	if !ok {
		t.Fatal("struct assertion failed")
	}
	t.Logf("BMW car is %s%s", bmwISeries.GetModel(), bmwISeries.GetSeries())

	bmw2, err := bmwFactory.NewVehicle(BmwXSeriesType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Vehicle is from %s company", bmw2.GetManufacturer())
	bmwXSeries, ok := bmw2.(*BmwXSeries)
	if !ok {
		t.Fatal("struct assertion failed")
	}
	t.Logf("BMW car is %s%s", bmwXSeries.GetSeries(), bmwXSeries.GetModel())

}

func TestFordFactory(t *testing.T) {
	fordFactory, err := GetVehicleFactory(FordFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	ford1, err := fordFactory.NewVehicle(FordSedanType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("Vehicle is from %s company", ford1.GetManufacturer()))
	fordSedanSeries, ok := ford1.(*FordSedanSeries)
	if !ok {
		t.Fatal("struct assertion failed")
	}
	t.Log(fmt.Sprintf("Ford car is %s %s", fordSedanSeries.GetType(), fordSedanSeries.GetModel()))

	ford2, err := fordFactory.NewVehicle(FordSuvType)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("Vehicle is from %s company", ford2.GetManufacturer()))
	fordSuvSeries, ok := ford2.(*FordSuvSeries)
	if !ok {
		t.Fatal("struct assertion failed")
	}
	t.Log(fmt.Sprintf("BMW car is %s %s", fordSuvSeries.GetType(), fordSuvSeries.GetModel()))

}
