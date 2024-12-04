package d4

func SearchForXmas2(input string) int {
	matrix := toIntMatrix(input)
	res := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == startChar {
				res += checkMas(matrix, i, j)
			}
		}
	}
	return res
}

const (
	startChar = 'A'
)

func checkMas(matrix [][]int32, i int, j int) int {
	res := 0
	if i == 0 || j == 0 {
		return 0
	}
	if i+1 == len(matrix) || j+1 == len(matrix[i]) {
		return 0
	}
	firstAxys := (valDir(matrix, i, j, NW) == m && valDir(matrix, i, j, SE) == s) || (valDir(matrix, i, j, NW) == s && valDir(matrix, i, j, SE) == m)
	if !firstAxys {
		return 0
	}
	secondAx := (valDir(matrix, i, j, NE) == m && valDir(matrix, i, j, SW) == s) || (valDir(matrix, i, j, NE) == s && valDir(matrix, i, j, SW) == m)
	if secondAx {
		return 1
	}
	return res
}

func valDir(matrix [][]int32, i int, j int, d direction) int32 {
	return matrix[i+d.deltaY][j+d.deltaX]
}
