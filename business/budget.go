package business

import (
	"log"
	"strconv"
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

	return b.Amount / float64(b.totalDay())
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
	lastDay, err = time.Parse(
		"20060102", b.YearMonth+strconv.Itoa(b.totalDay()),
	)
	if err != nil {
		log.Println("[ 判斷預算最後一天轉型失敗 ] Err:", err)
		return
	}

	return
}

// totalDay 預算總共有幾天
func (b *Budget) totalDay() int {

	date, err := time.Parse("200601", b.YearMonth)
	if err != nil {
		return 0
	}

	d := time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, time.UTC)

	return d.Day()
}

func (b *Budget) overlappingAmount(start, end time.Time) float64 {

	// 判斷是不是非法時間
	if end.Before(start) {
		return 0
	}

	// 判斷結束時間有沒有在預算最後一天之後
	if end.After(b.LastDay()) || start.Before(b.FirstDay()) {
		return 0
	}

	return b.periodDiff(start, end) * b.DailyAmount()
}

func (b *Budget) periodDiff(start, end time.Time) float64 {
	return end.Sub(start).Hours()/24 + 1
}
