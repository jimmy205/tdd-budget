package business

import "time"

// Accounting 計算組
type Accounting struct {
}

// NewAccount 新的計算組
func NewAccount() *Accounting {
	return &Accounting{}
}

// GetTotal 取得所有預算
func (a *Accounting) GetTotal(start, end time.Time) (total float64) {

	return 0
}
