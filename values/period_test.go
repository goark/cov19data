package values

import (
	"fmt"
	"testing"
	"time"
)

func TestPeriod(t *testing.T) {
	testCases := []struct {
		start, end, dt Date
		res            bool
	}{
		{start: NewDate(2020, time.March, 1), end: NewDate(2020, time.May, 31), dt: NewDate(2020, time.March, 0), res: false},
		{start: NewDate(2020, time.March, 1), end: NewDate(2020, time.May, 31), dt: NewDate(2020, time.March, 1), res: true},
		{start: NewDate(2020, time.May, 31), end: NewDate(2020, time.March, 1), dt: NewDate(2020, time.May, 31), res: true},
		{start: NewDate(2020, time.May, 31), end: NewDate(2020, time.March, 1), dt: NewDate(2020, time.May, 32), res: false},
		{start: NewDate(2020, time.March, 1), end: NewDateTime(time.Time{}), dt: NewDate(2020, time.March, 0), res: false},
		{start: NewDate(2020, time.March, 1), end: NewDateTime(time.Time{}), dt: NewDate(2020, time.March, 1), res: true},
		{start: NewDateTime(time.Time{}), end: NewDate(2020, time.May, 31), dt: NewDate(2020, time.May, 31), res: true},
		{start: NewDateTime(time.Time{}), end: NewDate(2020, time.May, 31), dt: NewDate(2020, time.May, 32), res: false},
		{start: NewDateTime(time.Time{}), end: NewDateTime(time.Time{}), dt: NewDate(2020, time.January, 1), res: true},
		{start: NewDateTime(time.Time{}), end: NewDateTime(time.Time{}), dt: NewDateTime(time.Time{}), res: false},
	}

	for _, tc := range testCases {
		p := NewPeriod(tc.start, tc.end)
		if res := p.Contains(tc.dt); res != tc.res {
			t.Errorf("{%v}.Contains(%v) = %v, want %v.", p, tc.dt, res, tc.res)
		} else {
			fmt.Println(p, "in", tc.dt, "=>", p.Contains(tc.dt))
		}
	}
}
