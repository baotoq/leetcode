package max_profit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	arr := []int{7, 1, 5, 3, 6, 4}
	act := maxProfit(arr)

	assert.Equal(t, 5, act)
}

func Test2(t *testing.T) {
	arr := []int{7, 6, 4, 3, 1}
	act := maxProfit(arr)

	assert.Equal(t, 0, act)
}
