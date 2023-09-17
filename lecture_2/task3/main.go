package main

import (
	"fmt"
	"github.com/nanmenkaimak/justcode-go/lecture_2/task3/oop"
)

func main() {
	var business oop.BusinessAccount
	business.Put(6000)

	var student oop.StudentAccount
	student.Put(600)
	student.SetScholarship(true)

	business.MonthIncomeWithTax(300)
	fmt.Println(business.Balance())

	student.MonthIncomeWithTax(300)
	fmt.Println(student.Balance())

	business.WithDraw(300)
	fmt.Println(business.Balance())
}
