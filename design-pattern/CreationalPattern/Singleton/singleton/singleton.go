package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	number int
}

var instance *singleton

func init() {
	instance = &singleton{
		number: 100,
	}
}

func GetInstance() Singleton {
	return instance
}

func (s *singleton) AddOne() int {
	s.number++
	return s.number
}
