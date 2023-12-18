package fact

import "testing"

func Test_TestCase_delimiters_largeInput_OneGoroutine(t *testing.T) {
	TestCase_delimitersOrderCorrectness{
		numGoroutine: 1,
		input:        []int{-1, -2, -3, -4, -5, 1, 2, 3, 4, 5, 100, 122, 4000},
	}.Run(t)
}

func Test_TestCase_delimiters_largeInput_MoreThanNeededGoroutines(t *testing.T) {
	TestCase_delimitersOrderCorrectness{
		numGoroutine: 50,
		input:        []int{1, 2, -1, 4},
	}.Run(t)
}

func Test_TestCase_delimiters_largeGoroutinesAmount(t *testing.T) {
	TestCase_delimitersOrderCorrectness{
		numGoroutine: 100000,
		input:        []int{-10, -225, -100, 250, 250, 9, 7, 5, 3, -117, -1},
	}.Run(t)
}

func Test_TestCase_factorization_largeInput_OneGoroutine(t *testing.T) {
	TestCase_factorizationCorrectness{
		numGoroutine: 1,
		input:        []int{-1, -2, -3, -4, -5, 1, 2, 3, 4, 5, 100, 122, 4000},
	}.Run(t)
}

func Test_TestCase_factorization_largeInput_MoreThanNeededGoroutines(t *testing.T) {
	TestCase_factorizationCorrectness{
		numGoroutine: 50,
		input:        []int{1, 2, -1, 4},
	}.Run(t)
}

func Test_TestCase_factorization_largeGoroutinesAmount(t *testing.T) {
	TestCase_factorizationCorrectness{
		numGoroutine: 100000,
		input:        []int{-10, -225, -100, 250, 250, 9, 7, 5, 3, -117, -1},
	}.Run(t)
}

func Test_TestCase_lineNumber_largeInput_OneGoroutine(t *testing.T) {
	TestCase_lineNumberCorrectness{
		numGoroutine: 1,
		input:        []int{-1, -2, -3, -4, -5, 1, 2, 3, 4, 5, 100, 122, 4000},
	}.Run(t)
}

func Test_TestCase_lineNumber_largeInput_MoreThanNeededGoroutines(t *testing.T) {
	TestCase_lineNumberCorrectness{
		numGoroutine: 50,
		input:        []int{1, 2, -1, 4},
	}.Run(t)
}

func Test_TestCase_lineNumber_largeGoroutinesAmount(t *testing.T) {
	TestCase_lineNumberCorrectness{
		numGoroutine: 100000,
		input:        []int{-10, -225, -100, 250, 250, 9, 7, 5, 3, -117, -1},
	}.Run(t)
}

func Test_TestCase_setNumbers_largeInput_OneGoroutine(t *testing.T) {
	TestCase_setNumbersCorrectness{
		numGoroutine: 1,
		input:        []int{-1, -2, -3, -4, -5, 1, 2, 3, 4, 5, 100, 122, 4000},
	}.Run(t)
}

func Test_TestCase_setNumbers_largeInput_MoreThanNeededGoroutines(t *testing.T) {
	TestCase_setNumbersCorrectness{
		numGoroutine: 50,
		input:        []int{1, 2, -1, 4},
	}.Run(t)
}

func Test_TestCase_setNumbers_largeGoroutinesAmount(t *testing.T) {
	TestCase_setNumbersCorrectness{
		numGoroutine: 100000,
		input:        []int{-10, -225, -100, 250, 250, 9, 7, 5, 3, -117, -1},
	}.Run(t)
}
