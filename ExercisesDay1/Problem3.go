package ExercisesDay1

import "fmt"

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
