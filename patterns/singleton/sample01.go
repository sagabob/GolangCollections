package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

func GetInstance() Singleton {
	return nil
}
func (s *singleton) AddOne() int {
	return 0
}
