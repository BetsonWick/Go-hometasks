package rangeI

import (
	"math"
	"testing"
)

func Test_length_Large(t *testing.T) {
	TestCase_length{
		gotRangeInt: NewRangeInt(-math.MaxInt32-1, math.MaxInt32),
		wantLength:  math.MaxInt32<<1 + 2,
	}.Run(t)
}

func TestIntersectLargeAndLarge(t *testing.T) {
	TestCase_intersect{
		aRangeInt:    NewRangeInt(-math.MaxInt32-1, math.MaxInt16),
		bRangeInt:    NewRangeInt(1, math.MaxInt32),
		wantRangeInt: NewRangeInt(1, math.MaxInt16),
	}.Run(t)
}

func Test_union_LargeAndLarge(t *testing.T) {
	TestCase_union{
		aRangeInt:    NewRangeInt(-math.MaxInt32-1, 0),
		bRangeInt:    NewRangeInt(1, math.MaxInt32),
		wantRangeInt: NewRangeInt(-math.MaxInt32-1, math.MaxInt32),
		wantUnion:    true,
	}.Run(t)
}

func Test_empty_NotEmptyLarge(t *testing.T) {
	TestCase_empty{
		gotRangeInt: NewRangeInt(-math.MaxInt32-1, math.MaxInt32),
		wantEmpty:   false,
	}.Run(t)
}

func Test_containsInt_Large(t *testing.T) {
	TestCase_containsInt{
		rangeInt:     NewRangeInt(-math.MaxInt32-1, math.MaxInt32),
		integer:      math.MaxInt32,
		wantContains: true,
	}.Run(t)
}

func Test_containsRange_LargeAndLarge(t *testing.T) {
	TestCase_containsRange{
		aRangeInt:    NewRangeInt(-math.MaxInt32-1, math.MaxInt32),
		bRangeInt:    NewRangeInt(math.MaxInt8, math.MaxInt16),
		wantContains: true,
	}.Run(t)
}

func Test_isIntersect_LargeAndLarge(t *testing.T) {
	TestCase_isIntersect{
		aRangeInt:       NewRangeInt(-math.MaxInt32-1, math.MaxInt8),
		bRangeInt:       NewRangeInt(-math.MaxInt8, math.MaxInt32),
		wantIsIntersect: true,
	}.Run(t)
}

func Test_toSlice_Large(t *testing.T) {
	TestCase_toSlice{
		rangeInt:  NewRangeInt(-10, 10),
		wantSlice: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}.Run(t)
}

func Test_minimum_Large(t *testing.T) {
	TestCase_minimum{
		rangeInt:  NewRangeInt(-math.MaxInt32-1, math.MaxInt32),
		isWantMin: true,
		wantMin:   -math.MaxInt32 - 1,
	}.Run(t)
}

func Test_maximum_Large(t *testing.T) {
	TestCase_maximum{
		rangeInt:  NewRangeInt(-math.MaxInt32-1, math.MaxInt32),
		isWantMax: true,
		wantMax:   math.MaxInt32,
	}.Run(t)
}

func Test_toString_Large(t *testing.T) {
	TestCase_toString{
		rangeInt:   NewRangeInt(-math.MaxInt32-1, math.MaxInt32),
		wantString: "[-2147483648,2147483647]",
	}.Run(t)
}
