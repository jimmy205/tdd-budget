package business

import (
	"testing"
	"time"

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
func (at *AccountingSuite) Test_GetTotal() {
	start, _ := time.Parse("2006-01-02", "2021-04-01")
	end, _ := time.Parse("2006-01-02", "2021-04-01")
	at.Equal(float64(0), at.Accounting.GetTotal(start, end))
}
