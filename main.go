package main

import (
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

type Employee interface {
	salaryCalculator()
}

type FullTime struct {
	salary float32
}

type Contractor struct {
	salary float32
}

type Freelancer struct {
	salary float32
	hours  int
}

func (ft FullTime) salaryCalculator() float32 {
	return ft.salary * 28
}

func (con Contractor) salaryCalculator() float32 {
	return con.salary * 28
}

func (fl Freelancer) salaryCalculator() float32 {
	return fl.salary * float32(fl.hours)
}

type Node struct {
	val         string
	left, right *Node
}

func (node *Node) preOrder() {
	if node == nil {
		return
	}
	fmt.Print(node.val)
	node.left.preOrder()
	node.right.preOrder()
}

func (node *Node) postOrder() {
	if node == nil {
		return
	}
	node.left.postOrder()
	node.right.postOrder()
	fmt.Print(node.val)
}

func Problem1() {
	var matrixA Matrix
	matrixA.rows = 2
	matrixA.columns = 3
	matrixA.matrix = [][]int{{0, 0, 0}, {0, 0, 0}}
	fmt.Println("matrixA:", matrixA)

	matrixA.addCellValue(0, 0, 1)
	matrixA.addCellValue(0, 1, 2)
	matrixA.addCellValue(0, 2, 3)
	matrixA.addCellValue(1, 0, 4)
	matrixA.addCellValue(1, 1, 5)
	matrixA.addCellValue(1, 2, 6)
	fmt.Println("matrixA:", matrixA)

	matrixB := Matrix{2, 3, [][]int{{7, 8, 9}, {10, 11, 12}}}
	fmt.Println("matrixB:", matrixB)

	matrixC := matrixA.addMatrix(matrixB)
	fmt.Println("matrixA + matrixB =", matrixC)
}

func Problem2() {
	emp1 := FullTime{500}
	emp2 := Contractor{100}
	hoursByFreeLancer := 20
	emp3 := Freelancer{10, hoursByFreeLancer}

	fmt.Printf("Full-time employees get paid %f per month\n", emp1.salaryCalculator())
	fmt.Printf("The contractor gets paid %f per month\n", emp2.salaryCalculator())
	fmt.Printf("Freelancer gets paid %f (if freelancer works %d hours)\n", emp3.salaryCalculator(), hoursByFreeLancer)
}

func Problem3() {
	var node2, node4, node5 Node
	node2.val, node4.val, node5.val = "a", "b", "c"
	node3 := Node{"-", &node4, &node5}
	node1 := Node{"+", &node2, &node3}

	node1.preOrder()
	fmt.Println()
	node1.postOrder()
}
func main() {
	Problem1()
	Problem2()
	Problem3()
}
