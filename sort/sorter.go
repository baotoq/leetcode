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

	return append([]int{}, nums...)
}

func (s *Sorter) bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums)-i; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return append([]int{}, nums...)
}
