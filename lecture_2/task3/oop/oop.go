package oop

type Account interface {
	WithDraw(num float64) bool
	Balance() float64
	Put(num float64)
	MonthIncomeWithTax(num float64)
}

type BaseAccount struct {
	balance float64
}

func (a *BaseAccount) WithDraw(num float64) bool {
	if a.balance >= num {
		a.balance -= num
		return true
	}
	return false
}

func (a *BaseAccount) Balance() float64 {
	return a.balance
}

func (a *BaseAccount) Put(num float64) {
	a.balance += num
}

func (a *BaseAccount) MonthIncomeWithTax(num float64) {
	num *= 0.9
	a.Put(num)
}

type BusinessAccount struct {
	BaseAccount
}

func (a *BusinessAccount) MonthIncomeWithTax(num float64) {
	num *= 0.8
	a.Put(num)
}

type StudentAccount struct {
	BaseAccount
	haveScholarship bool
}

func (a *StudentAccount) MonthIncomeWithTax(num float64) {
	if a.haveScholarship {
		a.Put(num)
	}
}

func (a *StudentAccount) SetScholarship(ok bool) {
	a.haveScholarship = ok
}
