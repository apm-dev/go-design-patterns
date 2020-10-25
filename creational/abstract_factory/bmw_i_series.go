package abstract_factory

type BmwISeries struct {
	model string
}

func (*BmwISeries) GetManufacturer() string {
	return "BMW"
}

func (*BmwISeries) GetSeries() string {
	return "i"
}

func (i *BmwISeries) GetModel() string {
	return i.model
}
