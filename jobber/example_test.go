package jobber_test

import (
	"fmt"
	"time"

	"github.com/Akagi201/utils-go/jobber"
)

func Example_simple() {
	timeStr := "*/2 * * * * *"

	res, err := jobber.ParseFullTimeSpec(timeStr)
	if err != nil {
		fmt.Errorf("Parse time string failed, err: %v\n", err)
		return
	}

	fmt.Printf("Time string: %v\n", res.String())

	now := time.Now()
	fmt.Printf("now: %v, next: %v\n", now, res.Next(now))
}
