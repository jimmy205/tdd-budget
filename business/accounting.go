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

		// 判斷結束時間有沒有在預算最後一天之後
		if end.After(budget.LastDay()) {
			return 0
		}

		// 判斷開始時間有沒有在預算第一天之後
		if start.Before(budget.FirstDay()) {
			return 0
		}

		return a.periodDiff(start, end) * budget.DailyAmount()
	}

	return 0
}

func (a *Accounting) periodDiff(start, end time.Time) float64 {
	return end.Sub(start).Hours()/24 + 1
}
