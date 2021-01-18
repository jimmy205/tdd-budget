package business

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
