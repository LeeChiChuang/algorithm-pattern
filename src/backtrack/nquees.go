package backtrack

// n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
//
// 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
// 示例:
//
// 输入: 4
// 输出: [
// [".Q..",  // 解法 1
// "...Q",
// "Q...",
// "..Q."],
//
// ["..Q.",  // 解法 2
// "Q...",
// "...Q",
// ".Q.."]
// ]
// 解释: 4 皇后问题存在两个不同的解法。
//  
//
// 提示：
//
// 皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一到七步，可进可退。
//
// 链接：https://leetcode-cn.com/problems/n-queens

// 全局变量,保存结果
var queenResult [][]string

func solveNQueens(n int) [][]string {
	// 清空变量
	queenResult = [][]string{}
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}
	backtrack(board, [][]byte{})
	return queenResult
}

func backtrack(board [][]bool, path [][]byte) {
	// 结束条件
	if len(path) == len(board) {
		t := make([]string, len(path))
		for k, bs := range path {
			t[k] = string(bs)
		}
		queenResult = append(queenResult, t)
	}

	for i := 0; i < len(board); i++ {
		// 不合法情况,剔除(剪枝)
		if !isValid(board, len(path), i) {
			continue
		}
		// 打印默认位置
		bs := printLine(len(board))
		// 放置皇后
		bs[i] = 'Q'
		// 放入棋盘
		board[len(path)][i] = true
		// 加入选择路径
		path = append(path, bs)
		// 进行下一次决策
		backtrack(board, path)
		// 撤销选择
		path = path[:len(path)-1]
		board[len(path)][i] = false
	}
}

// 是否能在 board[row][col] 位置放置皇后
// 皇后不可以上下左右对角线同时存在
func isValid(board [][]bool, row, col int) bool {
	// 检查行是否有皇后冲突
	for i := 0; i < row; i++ {
		if board[i][col] == true {
			return false
		}
	}
	// 检查列是否有皇后冲突
	for j := 0; j < len(board); j++ {
		if board[row][j] == true {
			return false
		}
	}
	// 检查对角线: "\"
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == true {
			return false
		}
	}
	// 检查对角线: "/"
	for i, j := row, col; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == true {
			return false
		}
	}

	return true
}

// 打印每行默认情况,都是'.'
func printLine(n int) []byte {
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = '.'
	}
	return bs
}

