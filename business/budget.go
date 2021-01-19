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

// dailyAmount 取得每日金額
func (b *Budget) dailyAmount() float64 {

	return b.Amount / float64(b.totalDay())
}

// firstDay 預算第一天
func (b *Budget) firstDay() (firstDay time.Time) {

	var err error
	firstDay, err = time.Parse("20060102", b.YearMonth+"01")
	if err != nil {
		log.Println("[ 判斷預算第一天轉型失敗 ] Err:", err)
		return
	}

	return
}

// lastDay 預算最後一天
func (b *Budget) lastDay() (lastDay time.Time) {

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

func (b *Budget) overlappingAmount(period *Period) float64 {

	// 判斷是不是非法時間
	if period.invalidPeriod() {
		return 0
	}

	if b.firstDay().After(period.end) {
		return 0
	}

	if period.start.After(b.lastDay()) {
		return 0
	}

	start := period.start
	if b.firstDay().After(period.start) {
		start = b.firstDay()
	}

	return (period.end.Sub(start).Hours()/24 + 1) * b.dailyAmount()
}
