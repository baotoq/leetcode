package sort

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func Test1(t *testing.T) {
	sorter := &Sorter{}

	arr := []int{5, 6, 7, 8, 9, 9, 10, 11}
	act := sorter.selectionSort(arr)

	slices.Sort(arr)

	assert.Equal(t, arr, act)
}

func Test2(t *testing.T) {
	sorter := &Sorter{}

	arr := []int{3, 2, 3, -1, 99, 6, 4, 22}
	act := sorter.selectionSort(arr)

	slices.Sort(arr)

	assert.Equal(t, arr, act)
}
