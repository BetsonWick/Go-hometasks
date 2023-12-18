package fact

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type Input struct {
	NumsOfGoroutine int   // n - число горутин
	Numbers         []int // слайс чисел, которые необходимо факторизовать
}

type Factorization interface {
	Work(Input, io.Writer) error
}

type FactorizationImpl struct {
	sync.Mutex
	counter atomic.Uint32
}

func (f *FactorizationImpl) countFactorization(n int) []string {
	result := make([]string, 0)
	if n < 0 {
		result = append(result, strconv.Itoa(-1))
		n *= -1
	}
	if n == 1 {
		return append(result, strconv.Itoa(1))
	}
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			n /= i
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

func (f *FactorizationImpl) writeAsMultipliers(writer io.Writer, numbers []string, number int) error {
	f.Lock()
	defer f.Unlock()
	formatted := fmt.Sprintf(
		"line %d, %d = %s\n",
		f.counter.Add(1),
		number,
		strings.Join(numbers, " * "),
	)
	_, err := writer.Write([]byte(formatted))
	if err != nil {
		return err
	}
	return nil
}

func (f *FactorizationImpl) Work(input Input, writer io.Writer) error {
	g := new(errgroup.Group)
	g.SetLimit(input.NumsOfGoroutine)

	for _, number := range input.Numbers {
		number := number
		g.Go(func() error {
			result := f.countFactorization(number)
			err := f.writeAsMultipliers(writer, result, number)
			if err != nil {
				return err
			}
			return nil
		},
		)
	}
	err := g.Wait()
	if err == nil {
		return nil
	}
	return err
}

func NewFactorization() *FactorizationImpl {
	return &FactorizationImpl{}
}
