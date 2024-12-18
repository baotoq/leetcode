package sort

type Sorter struct {
}

func (s *Sorter) selectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minPosition := i

		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minPosition] {
				minPosition = j
			}
		}

		nums[i], nums[minPosition] = nums[minPosition], nums[i]
	}

	return nums
}
