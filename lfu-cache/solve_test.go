package lfu_cache

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test1(t *testing.T) {
	capacity := 2
	input := [][]int{
		{1, 1},
		{2, 2},
		{1},
		{3, 3},
		{2},
		{3},
		{4, 4},
		{1},
		{3},
		{4},
	}

	ops := []string{
		"put", "put", "get", "put", "get", "get", "put", "get", "get", "get",
	}

	obj := Constructor(capacity)

	var result []*int
	for i, op := range ops {
		if op == "get" {
			v := obj.Get(input[i][0])
			result = append(result, &v)
		} else if op == "put" {
			obj.Put(input[i][0], input[i][1])
			result = append(result, nil)
		}
	}

	expected := []*int{nil, nil, getIntPointer(1), nil, getIntPointer(-1), getIntPointer(3), nil, getIntPointer(-1), getIntPointer(3), getIntPointer(4)}

	for i, e := range expected {
		if e == nil {
			assert.Nilf(t, result[i], "expected nil but actual %s at index %d", stringPtrToString(result[i]), i)
		} else {
			ok := assert.NotNilf(t, result[i], "expected NOT nil but actual %s at index %d", stringPtrToString(result[i]), i)
			if ok {
				assert.Equalf(t, *e, *result[i], "expected equal %d but actual %s at index %d", *e, stringPtrToString(result[i]), i)
			}
		}
	}
}

func Test2(t *testing.T) {
	capacity := 2
	input := [][]int{
		{2, 1},
		{1, 1},
		{2, 3},
		{4, 1},
		{1},
		{2},
	}

	ops := []string{
		"put", "put", "put", "put", "get", "get",
	}

	obj := Constructor(capacity)

	var result []*int
	for i, op := range ops {
		if op == "get" {
			v := obj.Get(input[i][0])
			result = append(result, &v)
		} else if op == "put" {
			obj.Put(input[i][0], input[i][1])
			result = append(result, nil)
		}
	}

	expected := []*int{nil, nil, nil, nil, getIntPointer(-1), getIntPointer(3)}

	for i, e := range expected {
		if e == nil {
			assert.Nilf(t, result[i], "expected nil but actual %s at index %d", stringPtrToString(result[i]), i)
		} else {
			ok := assert.NotNilf(t, result[i], "expected NOT nil but actual %s at index %d", stringPtrToString(result[i]), i)
			if ok {
				assert.Equalf(t, *e, *result[i], "expected equal %d but actual %s at index %d", *e, stringPtrToString(result[i]), i)
			}
		}
	}
}

func Test3(t *testing.T) {
	capacity := 2
	input := [][]int{
		{1, 1},
		{2, 2},
		{1},
		{3, 3},
		{2},
		{3},
		{4, 4},
		{1},
		{3},
		{4},
	}

	ops := []string{
		"put", "put", "get", "put", "get", "get", "put", "get", "get", "get",
	}

	obj := Constructor(capacity)

	var result []*int
	for i, op := range ops {
		if op == "get" {
			v := obj.Get(input[i][0])
			result = append(result, &v)
		} else if op == "put" {
			obj.Put(input[i][0], input[i][1])
			result = append(result, nil)
		}
	}

	expected := []*int{nil, nil, getIntPointer(1), nil, getIntPointer(-1), getIntPointer(3), nil, getIntPointer(-1), getIntPointer(3), getIntPointer(4)}

	for i, e := range expected {
		if e == nil {
			assert.Nilf(t, result[i], "expected nil but actual %s at index %d", stringPtrToString(result[i]), i)
		} else {
			ok := assert.NotNilf(t, result[i], "expected NOT nil but actual %s at index %d", stringPtrToString(result[i]), i)
			if ok {
				assert.Equalf(t, *e, *result[i], "expected equal %d but actual %s at index %d", *e, stringPtrToString(result[i]), i)
			}
		}
	}
}

func Test4(t *testing.T) {
	capacity := 2
	input := [][]int{
		{2, 1},
		{3, 2},
		{3},
		{2},
		{4, 3},
		{2},
		{3},
		{4},
	}

	ops := []string{
		"put", "put", "get", "get", "put", "get", "get", "get",
	}

	obj := Constructor(capacity)

	var result []*int
	for i, op := range ops {
		if op == "get" {
			v := obj.Get(input[i][0])
			result = append(result, &v)
		} else if op == "put" {
			obj.Put(input[i][0], input[i][1])
			result = append(result, nil)
		}
	}

	expected := []*int{nil, nil, getIntPointer(2), getIntPointer(1), nil, getIntPointer(1), getIntPointer(-1), getIntPointer(3)}

	for i, e := range expected {
		if e == nil {
			assert.Nilf(t, result[i], "expected nil but actual %s at index %d", stringPtrToString(result[i]), i)
		} else {
			ok := assert.NotNilf(t, result[i], "expected NOT nil but actual %s at index %d", stringPtrToString(result[i]), i)
			if ok {
				assert.Equalf(t, *e, *result[i], "expected equal %d but actual %s at index %d", *e, stringPtrToString(result[i]), i)
			}
		}
	}
}

func getIntPointer(val int) *int {
	return &val
}

func stringPtrToString(p *int) string {
	if p != nil {
		return strconv.Itoa(*p)
	}
	return "nil"
}
