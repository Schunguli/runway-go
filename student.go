package main

import "fmt"

type Student struct {
	Name   string
	Scores []float64
}

func (s Student) Average() float64 {
	if len(s.Scores) == 0 {
		return 0
	}

	var sum float64
	for _, score := range s.Scores {
		sum += score
	}

	return sum / float64(len(s.Scores))
}

func (s Student) Passed() bool {
	return s.Average() >= 50
}

func main() {
	student := Student{
		Name:   "Sasha",
		Scores: []float64{60, 70, 80},
	}

	fmt.Println("Name:", student.Name)
	fmt.Println("Average:", student.Average())
	fmt.Println("Passed:", student.Passed())
}