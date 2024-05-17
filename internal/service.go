// package internal

// type CounterService struct {
// 	counter int
// }

// func NewCounterService() *CounterService {
// 	return &CounterService{}
// } //  Bu ifade, CounterService adında bir nesne oluşturur ve bu nesnenin bellek adresini işaret eden bir işaretçi döndürür.

// func (cs *CounterService) Increment() {
// 	cs.counter++
// } // Bu ifade, CounterService içindeki "counter" ifadesinin değerini 1 arttırır.

// func (cs *CounterService) Decrement() {
// 	cs.counter--
// } // Bu ifade, CounterService içindeki "counter" ifadesinin değerini 1 azaltır.

// func (cs *CounterService) GetCounter() int {
// 	return cs.counter
// } // Bu ifade, CounterService içindeki "counter" ifadesini döndürür.

// // "{}" ifadesi nesne oluşturmak için kulanılır."
// // İşaretçiler(Pointers): Bir değişkenin bellekteki yerini ifade eden yapılardır. İşaretçiler, aynı zamanda verileri paylaşmak için kullanılır.
// // "*"ın işlevi, bir ifadenin başına konulduğunda o ifadenin bir işaretçi olduğunu bize gösterir. Aynı zamanda bu ifadenin bellekteki konumunu da bildiriyor.
// // *CounterService ifadesindeki "*"ın işlevi, CounterService ifadesinin bir işaretçi olduğunu bize gösterir.

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
