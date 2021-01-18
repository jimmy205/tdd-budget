package business

import (
	"testing"
	"time"

	"bou.ke/monkey"

	"github.com/stretchr/testify/suite"
)

// AccountingSuite 計算測試組
type AccountingSuite struct {
	suite.Suite
	*Accounting
}

func TestSuiteInit(t *testing.T) {
	suite.Run(t, new(AccountingSuite))
}

func (at *AccountingSuite) SetupTest() {
	at.Accounting = NewAccounting()
}

func (at *AccountingSuite) Test_no_budget() {
	start, _ := time.Parse("2006-01-02", "2021-04-01")
	end, _ := time.Parse("2006-01-02", "2021-04-01")

	at.totalShouldBe(start, end, 0)
}

func (at *AccountingSuite) Test_period_inside_budget_month() {
	start, _ := time.Parse("2006-01-02", "2021-04-01")
	end, _ := time.Parse("2006-01-02", "2021-04-30")

	// mock
	monkey.Patch(GetBudgets, func() []Budget {
		budgets := []Budget{
			{"202004", 3000},
		}
		return budgets
	})

	at.totalShouldBe(start, end, 3000)
}

func (at *AccountingSuite) totalShouldBe(start, end time.Time, expected float64) {
	at.Equal(expected, at.Accounting.GetTotal(start, end))
}
