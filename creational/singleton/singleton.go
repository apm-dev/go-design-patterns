package singleton

var instance *singleton

type singleton struct {
	Counter int
}

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.Counter++
	return s.Counter
}
