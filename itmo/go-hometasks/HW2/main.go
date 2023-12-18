package main

import (
	"math"
	"math/cmplx"
)

func getCharByIndex(str string, idx int) rune {
	return []rune(str)[idx]
}

func getStringBySliceOfIndexes(str string, indexes []int) string {
	var resultRunes = make([]rune, len(indexes))
	for idx, rn := range indexes {
		resultRunes[idx] = getCharByIndex(str, rn)
	}
	return string(resultRunes)
}

func addPointers(ptr1, ptr2 *int) *int {
	if ptr1 == nil {
		return ptr2
	}
	if ptr2 == nil {
		return ptr1
	}
	var result = *ptr1 + *ptr2
	return &result
}

func isComplexEqual(a, b complex128) bool {
	var delta = 1e-6
	return math.Abs(real(a)-real(b)) < delta && math.Abs(imag(a)-imag(b)) < delta
}

func getRootsOfQuadraticEquation(a, b, c float64) (complex128, complex128) {
	var complexA = complex(a, 0)
	var complexB = complex(b, 0)
	var complexC = complex(c, 0)
	var sqrtD = cmplx.Sqrt(complexB*complexB - 4*complexA*complexC)
	return (-complexB - sqrtD) / (2 * complexA), (-complexB + sqrtD) / (2 * complexA)
}

func mergeSort(s []int) []int {
	var length = len(s)
	if length <= 1 {
		return s
	}
	return merge(mergeSort(s[:length/2]), mergeSort(s[length/2:]))
}

func merge(a, b []int) []int {
	var result = make([]int, 0)
	var i int
	var j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		result = append(result, a[i])
	}
	for ; j < len(b); j++ {
		result = append(result, b[j])
	}
	return result
}

func reverseSliceOne(s []int) {
	var length = len(s)
	for idx := range s {
		if idx == length/2 {
			return
		}
		s[idx], s[length-idx-1] = s[length-idx-1], s[idx]
	}
}

func reverseSliceTwo(s []int) []int {
	var result = make([]int, len(s))
	copy(result, s)
	reverseSliceOne(result)
	return result
}

func swapPointers(a, b *int) {
	if a == nil || b == nil {
		panic("Pointers must address the existing value")
	}
	*a, *b = *b, *a
}

func isSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	for idx, value := range a {
		if value != b[idx] {
			return false
		}
	}
	return true
}

func deleteByIndex(s []int, idx int) []int {
	var result = make([]int, len(s))
	copy(result, s)
	return append(result[:idx], result[idx+1:]...)
}
