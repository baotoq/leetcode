package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_selectionSort1(t *testing.T) {
	var testData1 = []int{5, 6, 7, -8, 2, 4, 5, -2, -5, 9, 9, 10, 11, 0, -3}
	var expected1 = []int{-8, -5, -3, -2, 0, 2, 4, 5, 5, 6, 7, 9, 9, 10, 11}

	sorter := &Sorter{}

	act := sorter.selectionSort(testData1)

	assert.Equal(t, expected1, act)
}

func Test_bubbleSort1(t *testing.T) {
	var testData1 = []int{5, 6, 7, -8, 2, 4, 5, -2, -5, 9, 9, 10, 11, 0, -3}
	var expected1 = []int{-8, -5, -3, -2, 0, 2, 4, 5, 5, 6, 7, 9, 9, 10, 11}

	sorter := &Sorter{}

	act := sorter.bubbleSort(testData1)

	assert.Equal(t, expected1, act)
}
