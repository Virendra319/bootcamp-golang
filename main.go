package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Problem 1

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

// Problem 2

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

func Problem2() {
	var node2, node4, node5 Node
	node2.val, node4.val, node5.val = "a", "b", "c"
	node3 := Node{"-", &node4, &node5}
	node1 := Node{"+", &node2, &node3}

	node1.preOrder()
	fmt.Println()
	node1.postOrder()
}

// Problem 3

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

func Problem3() {
	emp1 := FullTime{500}
	emp2 := Contractor{100}
	hoursByFreeLancer := 20
	emp3 := Freelancer{10, hoursByFreeLancer}

	fmt.Printf("Full-time employees get paid %f per month\n", emp1.salaryCalculator())
	fmt.Printf("The contractor gets paid %f per month\n", emp2.salaryCalculator())
	fmt.Printf("Freelancer gets paid %f (if freelancer works %d hours)\n", emp3.salaryCalculator(), hoursByFreeLancer)
}

// problem 4

func count(word string, freq map[string]int) {
	for i := 0; i < len(word); i++ {
		freq[string(word[i])]++
	}
}
func Problem4() {
	words := []string{"quick", "brown", "fox", "lazy", "dog"}
	freq := make(map[string]int)

	for i := 0; i < len(words); i++ {
		count(words[i], freq)
	}
	result, err := json.MarshalIndent(freq, "", " ")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(result))
	}
}

// problem 5

func rate(id int, totalRating *int) {
	fmt.Println("Student -", id, "started rating")
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	*totalRating += rand.Intn(6)
	fmt.Println("Student -", id, "finished rating")
}

func Problem5() {
	totalStudents := 200
	var totalRating int
	var wg sync.WaitGroup

	for i := 1; i <= totalStudents; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			rate(i, &totalRating)
		}()
	}
	wg.Wait()
	fmt.Println("totalRating:", totalRating)
	averageRating := float32(totalRating) / float32(totalStudents)
	fmt.Println(averageRating)
}

// Problem 6

type BankAccounts struct {
	mu             sync.Mutex
	accountDetails map[string]int
}

func (ba *BankAccounts) deposit(accountNumber string, amount int) {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	ba.accountDetails[accountNumber] += amount
}
func (ba *BankAccounts) withdraw(accountNumber string, amount int) {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	if ba.accountDetails[accountNumber]-amount < 0 {
		//panic("can't withdraw an amount more than current balance")
		fmt.Println("Can't withdraw an amount more than current balance, ")
	} else {
		ba.accountDetails[accountNumber] -= amount
	}
}

func Problem6() {
	accounts := BankAccounts{
		accountDetails: map[string]int{"acc1": 500},
	}
	var wg sync.WaitGroup
	withdraw := func(accountNumber string, amount int) {
		accounts.withdraw(accountNumber, amount)
		wg.Done()
	}
	deposit := func(accountNumber string, amount int) {
		accounts.deposit(accountNumber, amount)
		wg.Done()
	}
	wg.Add(12)
	go withdraw("acc1", 150)
	go deposit("acc1", 50)
	go withdraw("acc1", 150)
	go deposit("acc1", 50)
	go withdraw("acc1", 150)
	go deposit("acc1", 50)
	go withdraw("acc1", 150)
	go deposit("acc1", 50)
	go withdraw("acc1", 150)
	go deposit("acc1", 50)
	go withdraw("acc1", 150)
	go deposit("acc1", 50)
	wg.Wait()

	fmt.Println(accounts.accountDetails["acc1"])
}

func main() {
	Problem1()
	Problem2()
	Problem3()
	Problem4()
	Problem5()
	Problem6()
}
