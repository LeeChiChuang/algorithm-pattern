package backtrack

// https://leetcode-cn.com/problems/subsets/
// 78. 子集
// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
// 说明：解集不能包含重复的子集。
//
// 示例:
//
// 输入: nums = [1,2,3]
// 输出:
// [
// [3],
// [1],
// [2],
// [1,2,3],
// [1,3],
// [2,3],
// [1,2],
// []
// ]

func Subsets(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	subsetBackTrack(nums, 0, list, &result)

	// [[],[1],[1,2],[1,2,3],[1,3],[2],[2,3],[3]]
	return result
}

func subsetBackTrack(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)

	for i:=pos; i<len(nums); i++ {
		list = append(list, nums[i])
		subsetBackTrack(nums, i+1, list, result)
		list = list[0:len(list)-1]
	}
}
