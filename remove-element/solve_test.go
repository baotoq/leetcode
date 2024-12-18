package remove_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	arr := []int{3, 2, 2, 3}
	act := removeElement(arr, 3)

	assert.Equal(t, 2, act)
	assert.Equal(t, []int{2, 2}, arr[:act])
}

func Test2(t *testing.T) {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	act := removeElement(arr, 2)

	assert.Equal(t, 5, act)
	assert.Equal(t, []int{0, 1, 3, 0, 4}, arr[:act])
}
