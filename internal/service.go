package internal

type CounterService struct {
	counter int
}

func NewCounterService() *CounterService {
	return &CounterService{}
}

func (cs *CounterService) Increment() {
	cs.counter++
}

func (cs *CounterService) Decrement() {
	cs.counter--
}

func (cs *CounterService) GetCounter() int {
	return cs.counter
}

var CounterServiceInstance *CounterService
