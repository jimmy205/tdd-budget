package business

import (
	"log"
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

	start, end := at.setStartEnd("2021-04-01", "2021-04-01")

	at.totalShouldBe(start, end, 0)
}

func (at *AccountingSuite) Test_period_inside_budget_month() {

	// mock
	mock := at.mockGetBudgets(map[string]float64{"202104": 3000})
	defer mock.Unpatch()

	start, end := at.setStartEnd("2021-04-01", "2021-04-30")
	at.totalShouldBe(start, end, 3000)
}

func (at *AccountingSuite) Test_period_inside_month() {

	// mock
	mock := at.mockGetBudgets(map[string]float64{"202104": 30})
	defer mock.Unpatch()

	start, end := at.setStartEnd("2021-04-01", "2021-04-01")
	at.totalShouldBe(start, end, 1)

}

func (at *AccountingSuite) Test_period_no_overlapping_first_day() {
	// mock
	mock := at.mockGetBudgets(map[string]float64{"202104": 30})
	defer mock.Unpatch()

	start, end := at.setStartEnd("2021-03-01", "2021-03-01")
	at.totalShouldBe(start, end, 0)
}

func (at *AccountingSuite) totalShouldBe(start, end time.Time, expected float64) {
	at.Equal(expected, at.Accounting.GetTotal(start, end))
}

func (at *AccountingSuite) mockGetBudgets(mockData map[string]float64) *monkey.PatchGuard {

	mock := monkey.Patch(GetBudgets, func() (budgets []Budget) {
		for yearMonth, amount := range mockData {
			budgets = append(budgets, Budget{yearMonth, amount})
		}
		return budgets
	})

	return mock
}

func (at *AccountingSuite) setStartEnd(startStr, endStr string) (
	start, end time.Time,
) {

	var err error
	start, err = time.Parse("2006-01-02", startStr)
	if err != nil {
		log.Println("[ 設定初始化開始時間失敗 ] GOT: ", startStr)
		return
	}
	end, err = time.Parse("2006-01-02", endStr)
	if err != nil {
		log.Println("[ 設定初始化結束時間失敗 ] GOT: ", endStr)
		return
	}

	return
}
