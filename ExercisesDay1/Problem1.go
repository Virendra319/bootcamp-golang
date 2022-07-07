package ExercisesDay1

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	rows, columns int
	matrix        [][]int
}

func (m1 Matrix) rowCount() int {
	return m1.rows
}

func (m1 Matrix) columnCount() int {
	return m1.columns
}

func (m1 *Matrix) addCellValue(i, j, x int) {
	m1.matrix[i][j] = x
}

func (m1 Matrix) addMatrix(m2 Matrix) Matrix {
	var result Matrix
	r1, r2, c1, c2 := m1.rowCount(), m2.rowCount(), m1.columnCount(), m2.columnCount()
	if (r1 != r2) || (c1 != c2) {
		return result
	}
	result.rows, result.columns = r1, c1
	result.matrix = m1.matrix
	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			result.matrix[i][j] += m2.matrix[i][j]
		}
	}
	return result
}

func (m1 Matrix) printAsJson() {
	val, err := json.MarshalIndent(m1.matrix, "", " ")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(val))
	}
}

func Problem1() {
	var matrixA Matrix
	matrixA.rows = 2
	matrixA.columns = 3
	matrixA.matrix = [][]int{{0, 0, 0}, {0, 0, 0}}
	//fmt.Println("matrixA:", matrixA)

	matrixA.addCellValue(0, 0, 1)
	matrixA.addCellValue(0, 1, 2)
	matrixA.addCellValue(0, 2, 3)
	matrixA.addCellValue(1, 0, 4)
	matrixA.addCellValue(1, 1, 5)
	matrixA.addCellValue(1, 2, 6)
	//fmt.Println("matrixA:", matrixA,)
	matrixA.printAsJson()

	matrixB := Matrix{2, 3, [][]int{{7, 8, 9}, {10, 11, 12}}}
	// fmt.Println("matrixB:", matrixB)
	matrixB.printAsJson()

	matrixC := matrixA.addMatrix(matrixB)
	// fmt.Println("matrixA + matrixB =", matrixC)
	matrixC.printAsJson()
}
