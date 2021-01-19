package business

import (
	"time"
)

// Accounting 計算組
type Accounting struct {
}

// NewAccounting 新的計算組
func NewAccounting() *Accounting {
	return &Accounting{}
}

// GetTotal 取得所有預算
func (a *Accounting) GetTotal(start, end time.Time) (total float64) {

	budgets := getBudgets()

	for _, budget := range budgets {
		period := newPeriod(start, end)
		another := newPeriod(budget.first, budget.last)

		total += period.overlappingDay(another) * budget.dailyAmount()
	}

	return total
}
