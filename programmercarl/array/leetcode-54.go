package array

/**
	title：螺旋矩阵I
	solution: 模拟，注意边界条件
**/

func spiralOrder(matrix [][]int) []int {

	m, n := len(matrix), len(matrix[0])
	res := []int{}

	left, right := 0, n-1
	top, bottom := 0, m-1

	for left <= right && top <= bottom {
		// move right
		for i := left; left <= right && i < right; i++ {
			res = append(res, matrix[top][i])
		}

		// move down
		for i := top; top <= bottom && i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}

		if left < right && top < bottom {
			// move left
			for i := right - 1; left <= right && i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}

			// move up
			for i := bottom - 1; top <= bottom && i > top; i-- {
				res = append(res, matrix[i][left])
			}
		}

		left++
		top++
		right--
		bottom--
	}

	return res
}
