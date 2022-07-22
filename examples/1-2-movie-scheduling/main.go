package main

import (
	"fmt"
	"sort"
	"time"
)

type movieJob struct {
	name      string
	startDate time.Time
	endDate   time.Time
}

func (job movieJob) overlaps(otherJob movieJob) bool {
	return job.startDate.After(otherJob.startDate) && job.startDate.Before(otherJob.endDate) || job.endDate.After(otherJob.startDate) && job.endDate.Before(otherJob.endDate)
}

/*
	Problem: Movie Scheduling Problem
	Input: A set I of n intervals on the line.
	Output: What is the largest subset of mutually non-overlapping intervals that can be selected from I?
*/

func main() {
	fmt.Println("Movie scheduling problem")
	location := time.UTC

	input := []movieJob{
		{name: "One", startDate: time.Date(2022, 1, 1, 1, 1, 1, 1, location), endDate: time.Date(2022, 2, 1, 1, 1, 1, 1, location)},
		{name: "Two", startDate: time.Date(2022, 1, 1, 1, 1, 1, 1, location), endDate: time.Date(2022, 3, 1, 1, 1, 1, 1, location)},
		{name: "Three", startDate: time.Date(2022, 2, 1, 1, 1, 1, 1, location), endDate: time.Date(2022, 4, 1, 1, 1, 1, 1, location)},
		{name: "Four", startDate: time.Date(2022, 2, 1, 1, 1, 1, 1, location), endDate: time.Date(2022, 6, 1, 1, 1, 1, 1, location)},
		{name: "Five", startDate: time.Date(2022, 2, 1, 1, 1, 1, 1, location), endDate: time.Date(2022, 3, 1, 1, 1, 1, 1, location)},
		{name: "Six", startDate: time.Date(2022, 2, 1, 1, 1, 1, 1, location), endDate: time.Date(2022, 4, 1, 1, 1, 1, 1, location)},
		{name: "Seven", startDate: time.Date(2022, 3, 1, 1, 1, 1, 1, location), endDate: time.Date(2022, 4, 1, 1, 1, 1, 1, location)},
		{name: "Eight", startDate: time.Date(2022, 2, 1, 1, 1, 1, 1, location), endDate: time.Date(2022, 5, 1, 1, 1, 1, 1, location)},
	}

	var result []movieJob
	s := input[:]

	sort.Slice(s, func(i, j int) bool {
		return s[i].endDate.Before(s[j].endDate)
	})

	for len(s) > 0 {
		movieDoneFirst := s[0]
		otherMovies := s[1:]
		buffer := []movieJob{}
		for i := range otherMovies {
			if !movieDoneFirst.overlaps(otherMovies[i]) {
				fmt.Printf("first movie: %+v compared to: %+v does not overlap\n", movieDoneFirst.name, otherMovies[i].name)
				buffer = append(buffer, otherMovies[i])
			}
		}
		s = buffer
		result = append(result, movieDoneFirst)
	}

	for i := range result {
		fmt.Printf("Result: %+v\n", result[i].name)
	}
}
