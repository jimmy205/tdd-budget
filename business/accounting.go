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

	budgets := GetBudgets()
	for _, budget := range budgets {
		return (end.Sub(start).Hours()/24 + 1) * budget.Amount / 30
	}

	return 0
}
