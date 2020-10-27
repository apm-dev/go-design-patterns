package builder

import "testing"

func TestBuilder(t *testing.T) {
	manufactureComplex := &ManufactureDirector{}
	// Test Car Builder
	carBuilder := &CarBuilder{}
	manufactureComplex.SetBuilder(carBuilder)
	manufactureComplex.Construct()
	car := carBuilder.GetVehicle()
	if car.Wheels != 4 {
		t.Errorf("wheels on a car must be 4 but were %d", car.Wheels)
	}
	if car.Seats != 5 {
		t.Errorf("seats on a car must be 5 but were %d", car.Seats)
	}
	if car.Structure != "Car" {
		t.Errorf("structure on a car must be Car but were %s", car.Structure)
	}
	// Test MotorBike Builder
	bikeBuilder := &MotorBikeBuilder{}
	manufactureComplex.SetBuilder(bikeBuilder)
	manufactureComplex.Construct()
	bike := bikeBuilder.GetVehicle()
	if bike.Wheels != 2 {
		t.Errorf("wheels on a bike must be 2 but were %d", car.Wheels)
	}
	if bike.Structure != "MotorBike" {
		t.Errorf("structure on a bike must be MotorBike but were %s", car.Structure)
	}

}
