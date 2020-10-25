package abstract_factory

type FordSedanSeries struct {
	model string
}

func (*FordSedanSeries) GetManufacturer() string {
	return "Ford"
}

func (*FordSedanSeries) GetType() string {
	return "Sedan"
}

func (f *FordSedanSeries) GetModel() string {
	return f.model
}
