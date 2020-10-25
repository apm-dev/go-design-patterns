package abstract_factory

type BmwXSeries struct {
	model  string
}

func (*BmwXSeries) GetManufacturer() string {
	return "BMW"
}

func (x *BmwXSeries) GetSeries() string {
	return "x"
}

func (x *BmwXSeries) GetModel() string {
	return x.model
}