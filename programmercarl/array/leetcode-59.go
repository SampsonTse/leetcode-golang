package array

/**
	title：螺旋矩阵II
	solution: 模拟，注意边界条件
**/

func generateMatrix(n int) [][]int {

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	left, right := 0, n-1
	top, bottom := 0, n-1
	k := 1

	for k <= n*n {
		// 右移
		for i := left; i < right; i++ {
			matrix[top][i] = k
			k++
		}

		// 下移
		for i := top; i < bottom; i++ {
			matrix[i][right] = k
			k++
		}

		// 左移
		for i := right; i >= left; i-- {
			matrix[bottom][i] = k
			k++
		}

		//上移
		for i := bottom - 1; i > top; i-- {
			matrix[i][left] = k
			k++
		}

		left++
		top++
		bottom--
		right--
	}

	return matrix

}
