package business

import (
	"log"
	"time"
)

// Budget 預算
type Budget struct {
	YearMonth string
	Amount    float64
}

// GetBudgets 取得預算
func GetBudgets() (budgets []Budget) {

	budgets = []Budget{}

	return
}

// DailyAmount 取得每日金額
func (b *Budget) DailyAmount() float64 {

	return b.Amount / 30
}

// FirstDay 預算第一天
func (b *Budget) FirstDay() (firstDay time.Time) {
	var err error
	firstDay, err = time.Parse("20060102", b.YearMonth+"01")
	if err != nil {
		log.Println("[ 判斷預算第一天轉型失敗 ] Err:", err)
		return
	}

	return
}

// LastDay 預算最後一天
func (b *Budget) LastDay() (lastDay time.Time) {

	var err error
	lastDay, err = time.Parse("20060102", b.YearMonth+"30")
	if err != nil {
		log.Println("[ 判斷預算最後一天轉型失敗 ] Err:", err)
		return
	}

	return
}
