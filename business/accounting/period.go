package accounting

import "time"

// Period 期間
type Period struct {
	start time.Time
	end   time.Time
}

// newPeriod 新的期間
func newPeriod(start, end time.Time) *Period {
	return &Period{
		start: start,
		end:   end,
	}
}

func (p *Period) overlappingDay(another *Period) float64 {

	if another.start.After(p.end) || another.end.Before(p.start) {
		return 0
	}

	start := p.start
	if another.start.After(p.start) {
		start = another.start
	}

	end := p.end
	if another.end.Before(p.end) {
		end = another.end
	}

	return end.Sub(start).Hours()/24 + 1
}
