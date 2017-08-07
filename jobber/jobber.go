// Package jobber time strings parser from jobber
package jobber

import (
	"fmt"
	"strconv"
	"strings"

	"time"

	"github.com/pkg/errors"
)

const (
	TimeWildcard = "*"
)

type TimeSpec interface {
	String() string
	Satisfied(int) bool
}

type FullTimeSpec struct {
	Sec  TimeSpec
	Min  TimeSpec
	Hour TimeSpec
	Mday TimeSpec
	Mon  TimeSpec
	Wday TimeSpec
}

func (self FullTimeSpec) String() string {
	return fmt.Sprintf("%v %v %v %v %v %v",
		self.Sec,
		self.Min,
		self.Hour,
		self.Mday,
		self.Mon,
		self.Wday)
}

func (self FullTimeSpec) Next(now time.Time) time.Time {
	/*
	 * We test every second from now till 2 years from now,
	 * looking for a time that satisfies the job's schedule
	 * criteria.
	 */

	var year time.Duration = time.Hour * 24 * 365
	var max time.Time = now.Add(2 * year)
	for next := now; next.Before(max); next = next.Add(time.Second) {
		a := self.Sec.Satisfied(next.Second()) &&
			self.Min.Satisfied(next.Minute()) &&
			self.Hour.Satisfied(next.Hour()) &&
			self.Wday.Satisfied(weekdayToInt(next.Weekday())) &&
			self.Mday.Satisfied(next.Day()) &&
			self.Mon.Satisfied(monthToInt(next.Month()))
		if a {
			return next
		}
	}

	return now
}

func weekdayToInt(d time.Weekday) int {
	switch d {
	case time.Sunday:
		return 0
	case time.Monday:
		return 1
	case time.Tuesday:
		return 2
	case time.Wednesday:
		return 3
	case time.Thursday:
		return 4
	case time.Friday:
		return 5
	default:
		return 6
	}
}

func monthToInt(m time.Month) int {
	switch m {
	case time.January:
		return 1
	case time.February:
		return 2
	case time.March:
		return 3
	case time.April:
		return 4
	case time.May:
		return 5
	case time.June:
		return 6
	case time.July:
		return 7
	case time.August:
		return 8
	case time.September:
		return 9
	case time.October:
		return 10
	case time.November:
		return 11
	default:
		return 12
	}
}

type WildcardTimeSpec struct {
}

func (s WildcardTimeSpec) String() string {
	return "*"
}

func (s WildcardTimeSpec) Satisfied(v int) bool {
	return true
}

type OneValTimeSpec struct {
	Val int
}

func (s OneValTimeSpec) String() string {
	return fmt.Sprintf("%v", s.Val)
}

func (s OneValTimeSpec) Satisfied(v int) bool {
	return s.Val == v
}

type SetTimeSpec struct {
	Desc string
	Vals []int
}

func (s SetTimeSpec) String() string {
	return s.Desc
}

func (s SetTimeSpec) Satisfied(v int) bool {
	for _, v2 := range s.Vals {
		if v == v2 {
			return true
		}
	}
	return false
}

func ParseFullTimeSpec(s string) (*FullTimeSpec, error) {
	var fullSpec FullTimeSpec
	fullSpec.Sec = WildcardTimeSpec{}
	fullSpec.Min = WildcardTimeSpec{}
	fullSpec.Hour = WildcardTimeSpec{}
	fullSpec.Mday = WildcardTimeSpec{}
	fullSpec.Mon = WildcardTimeSpec{}
	fullSpec.Wday = WildcardTimeSpec{}

	var timeParts []string = strings.Fields(s)

	// sec
	if len(timeParts) > 0 {
		spec, err := parseTimeSpec(timeParts[0], "sec", 0, 59)
		if err != nil {
			return nil, err
		}
		fullSpec.Sec = spec
	}

	// min
	if len(timeParts) > 1 {
		spec, err := parseTimeSpec(timeParts[1], "minute", 0, 59)
		if err != nil {
			return nil, err
		}
		fullSpec.Min = spec
	}

	// hour
	if len(timeParts) > 2 {
		spec, err := parseTimeSpec(timeParts[2], "hour", 0, 23)
		if err != nil {
			return nil, err
		}
		fullSpec.Hour = spec
	}

	// mday
	if len(timeParts) > 3 {
		spec, err := parseTimeSpec(timeParts[3], "month day", 1, 31)
		if err != nil {
			return nil, err
		}
		fullSpec.Mday = spec
	}

	// month
	if len(timeParts) > 4 {
		spec, err := parseTimeSpec(timeParts[4], "month", 1, 12)
		if err != nil {
			return nil, err
		}
		fullSpec.Mon = spec
	}

	// wday
	if len(timeParts) > 5 {
		spec, err := parseTimeSpec(timeParts[5], "weekday", 0, 6)
		if err != nil {
			return nil, err
		}
		fullSpec.Wday = spec
	}

	if len(timeParts) > 6 {
		return nil, errors.Wrap(nil, "Excess elements in 'time' field.")
	}

	return &fullSpec, nil
}

func parseTimeSpec(s string, fieldName string, min int, max int) (TimeSpec, error) {
	errMsg := fmt.Sprintf("Invalid '%v' value", fieldName)

	if s == TimeWildcard {
		return WildcardTimeSpec{}, nil
	} else if strings.HasPrefix(s, "*/") {
		// parse step
		stepStr := s[2:]
		step, err := strconv.Atoi(stepStr)
		if err != nil {
			return nil, errors.Wrap(err, errMsg)
		}

		// make set of valid values
		vals := make([]int, 0)
		for v := min; v <= max; v = v + step {
			vals = append(vals, v)
		}

		// make spec
		spec := SetTimeSpec{Vals: vals, Desc: s}
		return spec, nil

	} else if strings.Contains(s, ",") {
		// split step
		stepStrs := strings.Split(s, ",")

		// make set of valid values
		vals := make([]int, 0)
		for _, stepStr := range stepStrs {
			step, err := strconv.Atoi(stepStr)
			if err != nil {
				return nil, errors.Wrap(err, errMsg)
			}
			vals = append(vals, step)
		}

		// make spec
		spec := SetTimeSpec{Vals: vals, Desc: s}
		return spec, nil
	} else if strings.Contains(s, "-") {
		// get range extremes
		extremes := strings.Split(s, "-")
		begin, err := strconv.Atoi(extremes[0])

		if err != nil {
			return nil, errors.Wrap(err, errMsg)
		}

		end, err := strconv.Atoi(extremes[1])

		if err != nil {
			return nil, errors.Wrap(err, errMsg)
		}

		// make set of valid values
		vals := make([]int, 0)

		for v := begin; v <= end; v++ {
			vals = append(vals, v)
		}

		// make spec
		spec := SetTimeSpec{Vals: vals, Desc: s}
		return spec, nil
	} else {
		// convert to int
		val, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.Wrap(err, errMsg)
		}

		// make TimeSpec
		spec := OneValTimeSpec{val}

		// check range
		if val < min {
			errMsg := fmt.Sprintf("%s: cannot be less than %v.", errMsg, min)
			return nil, errors.Wrap(nil, errMsg)
		} else if val > max {
			errMsg := fmt.Sprintf("%s: cannot be greater than %v.", errMsg, max)
			return nil, errors.Wrap(nil, errMsg)
		}

		return spec, nil
	}
}
