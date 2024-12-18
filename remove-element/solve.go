package remove_element

// https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150

func removeElement(nums []int, val int) int {
	arr := []int{}
	for _, n := range nums {
		if n != val {
			arr = append(arr, n)
		}
	}

	for i, n := range arr {
		nums[i] = n
	}

	return len(arr)
}
