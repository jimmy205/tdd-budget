package business

import "time"

// Period 期間
type Period struct {
	start time.Time
	end   time.Time
}

// NewPeriod 新的期間
func NewPeriod(start, end time.Time) *Period {
	return &Period{
		start: start,
		end:   end,
	}
}

// 期間相差幾天
func (p *Period) diff() float64 {
	return p.end.Sub(p.start).Hours()/24 + 1
}

func (p *Period) invalidPeriod() bool {
	return p.end.Before(p.start)
}
