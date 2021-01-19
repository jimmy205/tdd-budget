package business

import (
	"log"
	"tddbudget/repository"
	"time"
)

// Budget 預算
type Budget struct {
	yearMonth string
	amount    float64
	first     time.Time
	last      time.Time
}

// getBudgets 取得預算
func getBudgets() (budgets []*Budget) {

	data := repository.GetBudgets()
	for yearMonth, amount := range data {

		// 取得每個月的第一天
		first, err := time.Parse("20060102", yearMonth+"01")
		if err != nil {
			log.Println("[ Err ] parse date error :", err)
			continue
		}

		// 取得每個月的最後一天
		date, err := time.Parse("200601", yearMonth)
		if err != nil {
			log.Println("[ Err ] parse date error :", err)
			continue
		}

		last := time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, time.UTC)

		budgets = append(budgets, &Budget{yearMonth, amount, first, last})
	}

	return
}

// dailyAmount 取得每日金額
func (b *Budget) dailyAmount() float64 {

	return b.amount / float64(b.last.Day())
}
