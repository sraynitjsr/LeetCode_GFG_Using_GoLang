package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	startTime int
	endTime   int
}

func main() {
	intervals := []Interval{
		{
			startTime: 10,
			endTime:   15,
		},
		{
			startTime: 17,
			endTime:   20,
		},
		{
			startTime: 25,
			endTime:   40,
		},
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].startTime < intervals[j].startTime
	})

	endTime := intervals[0].endTime

	for i := 1; i < len(intervals); i++ {
		currentStartTime := intervals[i].startTime
		currentEndTime := intervals[i].endTime
		if currentStartTime < endTime {
			fmt.Println("Invalid Intervals")
			return
		} else {
			endTime = currentEndTime
		}
	}
	fmt.Println("Valid Intervals")
}
