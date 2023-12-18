package rangeI

import "fmt"

type RangeInt interface {
	Length() int
	Intersect(other RangeInt)
	Union(other RangeInt) bool
	IsEmpty() bool
	ContainsInt(i int) bool
	ContainsRange(other RangeInt) bool
	IsIntersect(other RangeInt) bool
	ToSlice() []int
	Minimum() (int, bool)
	Maximum() (int, bool)
	String() string
}

func NewRangeInt(a, b int) RangeInt {
	return &RangeIntImpl{left: a, right: b}
}

type RangeIntImpl struct {
	left  int
	right int
}

func (r *RangeIntImpl) Length() int {
	if r.IsEmpty() {
		return 0
	}
	return r.right - r.left + 1
}

func (r *RangeIntImpl) Intersect(other RangeInt) {
	if !r.IsIntersect(other) {
		r.left, r.right = 1, 0
		return
	}
	var otherLeft, _ = other.Minimum()
	var otherRight, _ = other.Maximum()
	r.left = max(r.left, otherLeft)
	r.right = min(r.right, otherRight)
}

func (r *RangeIntImpl) Union(other RangeInt) bool {
	var otherLeft, _ = other.Minimum()
	var otherRight, _ = other.Maximum()
	if r.IsEmpty() && !other.IsEmpty() {
		r.left = otherLeft
		r.right = otherRight
		return true
	}
	if r.IsIntersect(other) || r.left == otherRight+1 || r.right == otherLeft-1 {
		r.left = min(r.left, otherLeft)
		r.right = max(r.right, otherRight)
		return true
	}
	return false
}

func (r *RangeIntImpl) IsEmpty() bool {
	return r.left > r.right
}

func (r *RangeIntImpl) ContainsInt(i int) bool {
	return r.left <= i && r.right >= i
}

func (r *RangeIntImpl) ContainsRange(other RangeInt) bool {
	var otherLeft, _ = other.Minimum()
	var otherRight, _ = other.Maximum()
	return r.left <= otherLeft && r.right >= otherRight
}

func (r *RangeIntImpl) IsIntersect(other RangeInt) bool {
	if r.IsEmpty() || other.IsEmpty() {
		return false
	}
	var otherLeft, _ = other.Minimum()
	var otherRight, _ = other.Maximum()
	return r.right >= otherLeft && r.left <= otherRight
}

func (r *RangeIntImpl) ToSlice() []int {
	var result = make([]int, r.Length())
	for i := range result {
		result[i] = r.left + i
	}
	return result
}

func (r *RangeIntImpl) Minimum() (int, bool) {
	return r.left, !r.IsEmpty()
}
func (r *RangeIntImpl) Maximum() (int, bool) {
	return r.right, !r.IsEmpty()
}
func (r *RangeIntImpl) String() string {
	if r.IsEmpty() {
		return ""
	}
	return fmt.Sprintf("[%d,%d]", r.left, r.right)
}
