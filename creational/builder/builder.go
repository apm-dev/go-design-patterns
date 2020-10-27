package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// Manufacture Director
type ManufactureDirector struct {
	builder BuildProcess
}

func (m *ManufactureDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

func (m *ManufactureDirector) Construct() {
	m.builder.SetWheels().SetSeats().SetStructure()
}

// Vehicle
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// Car Builder
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// MotorBike Builder
type MotorBikeBuilder struct {
	v VehicleProduct
}

func (m *MotorBikeBuilder) SetWheels() BuildProcess {
	m.v.Wheels = 2
	return m
}

func (m *MotorBikeBuilder) SetSeats() BuildProcess {
	m.v.Seats = 2
	return m
}

func (m *MotorBikeBuilder) SetStructure() BuildProcess {
	m.v.Structure = "MotorBike"
	return m
}

func (m *MotorBikeBuilder) GetVehicle() VehicleProduct {
	return m.v
}
