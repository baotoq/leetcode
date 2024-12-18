package majority_element

import (
	"maps"
	"slices"
	"sort"
)

// https://leetcode.com/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150

func majorityElement(nums []int) int {
	appearCount := make(map[int]int)

	for _, num := range nums {
		appearCount[num]++
	}

	keys := slices.Collect(maps.Keys(appearCount))
	sort.Slice(keys, func(i, j int) bool {
		return appearCount[keys[i]] > appearCount[keys[j]]
	})

	return keys[0]
}
