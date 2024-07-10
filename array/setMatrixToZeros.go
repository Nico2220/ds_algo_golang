package array

func SetMatrixToZeros(matrix [][]int) {
	rows := make([]int, len(matrix))
	cols := make([]int, len(matrix[0]))

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(cols); j++ {
			if matrix[i][j] == 0 {
				rows[i] = 1
				cols[j] = 1
			}
		}
	}

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(cols); j++ {
			if rows[i] == 1 || cols[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}

func SetMatrixToZeros_solution2(matrix [][]int) {
	// rows := make([]int, len(matrix))
	// cols := make([]int, len(matrix[0]))

	//cols := matrix[0][j]
	//rows := matrix[i][0]
	col0 := 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				if j != 0 {
					matrix[0][j] = 0
				} else {
					col0 = 0
				}
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	if col0 == 0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

}
