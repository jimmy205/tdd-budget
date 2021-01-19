package accounting

import (
	"log"
	"tddbudget/repository"
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

	mock := at.mockGetBudgets(map[string]float64{"202104": 30})
	defer mock.Unpatch()

	start, end := at.setStartEnd("2021-04-01", "2021-04-30")
	at.totalShouldBe(start, end, 30)
}

func (at *AccountingSuite) Test_period_inside_month() {

	mock := at.mockGetBudgets(map[string]float64{"202104": 30})
	defer mock.Unpatch()

	start, end := at.setStartEnd("2021-04-01", "2021-04-01")
	at.totalShouldBe(start, end, 1)

}

func (at *AccountingSuite) Test_period_no_overlapping_first_day() {

	mock := at.mockGetBudgets(map[string]float64{"202104": 30})
	defer mock.Unpatch()

	start, end := at.setStartEnd("2021-03-01", "2021-03-01")
	at.totalShouldBe(start, end, 0)
}

func (at *AccountingSuite) Test_period_no_overlapping_last_day() {

	mock := at.mockGetBudgets(map[string]float64{"202104": 30})
	defer mock.Unpatch()

	start, end := at.setStartEnd("2021-05-01", "2021-05-01")
	at.totalShouldBe(start, end, 0)
}

func (at *AccountingSuite) Test_invalid_period() {

	mock := at.mockGetBudgets(map[string]float64{"202104": 30})
	defer mock.Unpatch()

	start, end := at.setStartEnd("2021-05-30", "2021-05-01")
	at.totalShouldBe(start, end, 0)
}

func (at *AccountingSuite) Test_period_ovelapping_budget_first_day() {

	mock := at.mockGetBudgets(map[string]float64{"202104": 30})
	defer mock.Unpatch()

	start, end := at.setStartEnd("2021-03-31", "2021-04-02")
	at.totalShouldBe(start, end, 2)
}

func (at *AccountingSuite) Test_period_ovelapping_budget_last_day() {

	mock := at.mockGetBudgets(map[string]float64{"202104": 30})
	defer mock.Unpatch()

	start, end := at.setStartEnd("2021-04-28", "2021-05-02")
	at.totalShouldBe(start, end, 3)
}

func (at *AccountingSuite) Test_daily_budget_10() {

	mock := at.mockGetBudgets(map[string]float64{"202104": 300})
	defer mock.Unpatch()

	start, end := at.setStartEnd("2021-04-01", "2021-04-03")
	at.totalShouldBe(start, end, 30)
}

func (at *AccountingSuite) Test_cross_2_month() {

	mock := at.mockGetBudgets(map[string]float64{"202104": 30, "202105": 310})
	defer mock.Unpatch()

	start, end := at.setStartEnd("2021-04-01", "2021-05-02")
	at.totalShouldBe(start, end, 30+20)
}

func (at *AccountingSuite) Test_cross_3_month() {

	mock := at.mockGetBudgets(map[string]float64{
		"202104": 30,
		"202105": 310,
		"202106": 3000,
	})
	defer mock.Unpatch()

	start, end := at.setStartEnd("2021-04-01", "2021-06-02")
	at.totalShouldBe(start, end, 30+310+200)
}

func (at *AccountingSuite) totalShouldBe(start, end time.Time, expected float64) {
	at.Equal(expected, at.Accounting.GetTotal(start, end))
}

func (at *AccountingSuite) mockGetBudgets(mockData map[string]float64) *monkey.PatchGuard {

	mock := monkey.Patch(repository.GetBudgets, func() map[string]float64 {
		return mockData
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
