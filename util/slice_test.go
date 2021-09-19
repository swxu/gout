package util

import "testing"

func TestSliceElemEqual(t *testing.T) {
	a := []int64{1, 2, 3}
	b := []int64{3, 2, 1}
	t.Log(SliceElemEqual(a, b)) // true

	a = []int64{1, 2, 3}
	b = []int64{3, 1, 1}
	t.Log(SliceElemEqual(a, b)) // false

	a = []int64{1, 2, 3}
	b = []int64{3, 1, 2}
	t.Log(SliceElemEqual(interface{}(a), interface{}(b))) // true

	b = []int64{}
	t.Log(SliceElemEqual(nil, b)) // false

	t.Log(SliceElemEqual(nil, nil)) // true

	a = []int64{}
	b = []int64{}
	t.Log(SliceElemEqual(a, b)) // true

	a = []int64{}
	s := []string{}
	t.Log(SliceElemEqual(a, s)) // false

	a = []int64{1}
	s = []string{"1"}
	t.Log(SliceElemEqual(a, s)) // false
}

func TestSliceContains(t *testing.T) {
	s := []int{1, 2, 2, 3, 4}
	t.Log(SliceContains(s, 2))                           // true
	t.Log(SliceContains(s, 0))                           // false
	t.Log(SliceContains(s, 5))                           // false
	t.Log(SliceContains(s, "1"))                         // false
	t.Log(SliceContains(interface{}(s), interface{}(2))) // true
	t.Log(SliceContains(interface{}(s), nil))            // false
}
