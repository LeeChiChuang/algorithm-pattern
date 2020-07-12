package backtrack

// 46. 全排列
// 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
// 示例:
//
// 输入: [1,2,3]
// 输出:
// [
// [1,2,3],
// [1,3,2],
// [2,1,3],
// [2,3,1],
// [3,1,2],
// [3,2,1]
// ]

// 回溯模板
//
// result := []int{}
// func backTrack(路径， 选择列表) {
// 	if 满足终止条件 {
// 		return
// 	}
//
// 	// 做选择
// 	将选择从选择列表中删除
// 	路径.add(选择)
// 	backTrack(路径， 选择列表)
// 	// 撤销选择
// 	路径.remove(选择)
// 	将该选择加入选择列表
// }

var result [][]int

func Permute(nums []int) [][]int {
	result = [][]int{}
	used := make(map[int]bool, len(nums))
	appendPath := make([]int, 0)
	backTrace(appendPath, nums, used)

	return result
}

func backTrace(appendPath []int, nums []int, used map[int]bool) {
	// 终止条件
	if len(appendPath) == len(nums) {
		tmp := make([]int, len(appendPath))
		copy(tmp, appendPath)
		result = append(result, tmp)
		return
	}

	for i:=0; i<len(nums); i++ {
		if used[i] {
			continue
		}
		// 做选择
		used[i] = true
		appendPath = append(appendPath, nums[i])
		backTrace(appendPath, nums, used)
		// 撤销选择
		appendPath = appendPath[:len(appendPath)-1]
		used[i] = false
	}
}