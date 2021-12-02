package day1

import (
	"sync"
)

type Measurement int

type Measurements []Measurement

func (m Measurements) HasIncreasedOver(a Measurements) bool {
	return Sum(m) > Sum(a)
}

type Counter struct {
	mu    sync.Mutex
	wg    *sync.WaitGroup
	count int
}

func NewCounter() *Counter {
	return &Counter{
		wg: &sync.WaitGroup{},
	}
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

type Handler func(input Measurements, counter *Counter)

func RunSolution(input Measurements, fn Handler) *Counter {
	counter := NewCounter()
	fn(input, counter)
	counter.wg.Wait()
	return counter
}

func Solution1(input Measurements, counter *Counter) {
	total := len(input) - 1
	counter.wg.Add(total)

	for i := 0; i < total; i++ {
		a := input[i]
		b := input[i+1]
		go CountSums(counter, Measurements{a}, Measurements{b})
	}
}

func Solution2(input Measurements, counter *Counter) {
	total := len(input) - 3
	counter.wg.Add(total)

	for i := 0; i < total; i++ {
		a := input[i : i+3]
		b := input[i+1 : i+4]
		go CountSums(counter, a, b)
	}
}

func CountSums(counter *Counter, a, b Measurements) {
	defer counter.wg.Done()
	if b.HasIncreasedOver(a) {
		counter.Increment()
	}
}

func Sum(measurements Measurements) (sum Measurement) {
	for _, m := range measurements {
		sum += m
	}
	return
}
