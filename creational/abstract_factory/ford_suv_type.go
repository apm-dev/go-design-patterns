package abstract_factory

type FordSuvSeries struct {
	model string
}

func (*FordSuvSeries) GetManufacturer() string {
	return "Ford"
}

func (*FordSuvSeries) GetType() string {
	return "SUV"
}

func (f *FordSuvSeries) GetModel() string {
	return f.model
}
