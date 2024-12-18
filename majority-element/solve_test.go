package majority_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	arr := []int{2, 2, 1, 1, 1, 2, 2}
	act := majorityElement(arr)

	assert.Equal(t, 2, act)
}

func Test2(t *testing.T) {
	arr := []int{3, 2, 3}
	act := majorityElement(arr)

	assert.Equal(t, 3, act)
}
