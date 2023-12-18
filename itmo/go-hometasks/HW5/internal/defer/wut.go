package _defer

import (
	"fmt"
	"io"
)

type Counter struct {
	writer io.Writer
	value  int
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Printer() func() {
	value := c.value
	printer := func() {
		_, _ = fmt.Fprint(c.writer, value)
	}
	return printer
}

func PrintSequence1(writer io.Writer) {
	counter := Counter{writer, 0}
	counter.Increment()
	counter.Printer()()
	counter.Increment()
	defer counter.Printer()()
	counter.Increment()
	counter.Printer()()
	counter.Increment()
	defer counter.Printer()()
	counter.Increment()
	counter.Printer()()
	counter.Increment()
	counter.Printer()()
}

func PrintSequence2(writer io.Writer) {
	counter := Counter{writer, 0}
	defer counter.Printer()()
	counter.Increment() // 1
	counter.Printer()()
	counter.Increment() // 2
	func() {
		func() {
			defer counter.Printer()()
			counter.Increment() // 3
			counter.Printer()()
			counter.Increment() // 4
			counter.Printer()()
			counter.Increment() // 5
			counter.Printer()()
		}()
		counter.Increment() // 6
		defer counter.Printer()()
		func() {
			counter.Increment() // 7
			defer counter.Printer()()
			counter.Increment() // 8
			defer counter.Printer()()
		}()
		counter.Increment() // 9
		counter.Printer()()
	}()
}
