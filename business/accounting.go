package business

import (
	"log"
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

		budgetFisrtDay, err := time.Parse("20060102", budget.YearMonth+"01")
		if err != nil {
			log.Println("[ 判斷預算第一天失敗 ] Err:", err)
			continue
		}

		// 判斷開始時間有沒有在預算第一天之後
		if start.Before(budgetFisrtDay) {
			return 0
		}

		return a.periodDiff(start, end) * budget.DailyAmount()
	}

	return 0
}

func (a *Accounting) periodDiff(start, end time.Time) float64 {
	return end.Sub(start).Hours()/24 + 1
}
