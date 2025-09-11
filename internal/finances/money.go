package finances

import "fmt"

type Money struct {
	Cents int64
}

func NewMoneyFromCents(cents int64) Money {
	return Money{Cents: cents}
}

func NewMoney(euros int64, cents int64) Money {
	return Money{Cents: euros*100 + cents}
}

func (m Money) String() string {
	euros := m.Cents / 100
	cents := m.Cents % 100
	return fmt.Sprintf("%d.%02d", euros, cents)
}

func (m Money) Add(other Money) Money {
	return Money{Cents: m.Cents + other.Cents}
}

func (m Money) Sub(other Money) Money {
	return Money{Cents: m.Cents - other.Cents}
}
