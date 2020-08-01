package sort

func QuickSort(nums []int, begin, end int) []int {
	if end <= begin {
		return nil
	}
	pivot := partition(nums, begin, end)
	QuickSort(nums, begin, pivot-1)
	QuickSort(nums, pivot+1, end)

	return nums
}

func partition(nums []int, begin, end int) int {
	pivot, counter := end, begin
	for i := begin; i < end; i++ {
		if nums[i] < nums[pivot] {
			temp := nums[counter]
			nums[counter] = nums[i]
			nums[i] = temp
			counter++
		}
	}

	temp := nums[pivot]
	nums[pivot] = nums[counter]
	nums[counter] = temp

	return counter
}
