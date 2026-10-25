package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Score int
}

func Grade(score int) string {
	switch {
	case score < 0 || score > 100:
		return "?"
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func ParseLine(line string) (Student, error) {
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return Student{}, fmt.Errorf("expected name,score")
	}

	name := strings.TrimSpace(parts[0])
	if name == "" {
		return Student{}, fmt.Errorf("name is empty")
	}

	score, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return Student{}, fmt.Errorf("score must be an integer")
	}
	if score < 0 || score > 100 {
		return Student{}, fmt.Errorf("score must be between 0 and 100")
	}

	return Student{Name: name, Score: score}, nil
}

func ReadStudents(r io.Reader, report io.Writer) ([]Student, error) {
	var students []Student
	scanner := bufio.NewScanner(r)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			fmt.Fprintf(report, "Line %d: blank line skipped\n", lineNumber)
			continue
		}

		student, err := ParseLine(line)
		if err != nil {
			fmt.Fprintf(report, "Line %d: %v\n", lineNumber, err)
			continue
		}

		students = append(students, student)
	}

	if err := scanner.Err(); err != nil {
		return students, err
	}

	return students, nil
}

func PrintReport(w io.Writer, students []Student) {
	fmt.Fprintf(w, "%-20s %5s %6s\n", "Name", "Score", "Grade")
	fmt.Fprintf(w, "%-20s %5s %6s\n",
		strings.Repeat("-", 20),
		strings.Repeat("-", 5),
		strings.Repeat("-", 6),
	)

	if len(students) == 0 {
		fmt.Fprintln(w, "No valid scores to summarize.")
		return
	}

	total := 0
	highest := students[0].Score
	lowest := students[0].Score

	for _, student := range students {
		fmt.Fprintf(w, "%-20s %5d %6s\n",
			student.Name, student.Score, Grade(student.Score))

		total += student.Score
		if student.Score > highest {
			highest = student.Score
		}
		if student.Score < lowest {
			lowest = student.Score
		}
	}

	average := float64(total) / float64(len(students))
	fmt.Fprintf(w, "\nClass average: %.2f\nHighest: %d\nLowest: %d\n",
		average, highest, lowest)
}

func run(args []string, out, errOut io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(errOut, "Usage: grading <file>")
		return 2
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(errOut, "Could not open %q: %v\n", args[0], err)
		return 1
	}
	defer file.Close()

	students, err := ReadStudents(file, errOut)
	if err != nil {
		fmt.Fprintf(errOut, "Could not read %q: %v\n", args[0], err)
		return 1
	}

	PrintReport(out, students)
	return 0
}

func runChecks(out io.Writer) int {
	passed := 0
	failed := 0

	check := func(name string, ok bool) {
		if ok {
			fmt.Fprintf(out, "PASS: %s\n", name)
			passed++
		} else {
			fmt.Fprintf(out, "FAIL: %s\n", name)
			failed++
		}
	}

	check("Grade boundaries",
		Grade(0) == "F" &&
			Grade(60) == "D" &&
			Grade(70) == "C" &&
			Grade(80) == "B" &&
			Grade(90) == "A" &&
			Grade(100) == "A",
	)

	student, err := ParseLine("Ada Lovelace,95")
	check("Parse valid line",
		err == nil && student == (Student{Name: "Ada Lovelace", Score: 95}),
	)

	_, err = ParseLine("Grace,101")
	check("Reject score above 100", err != nil)

	_, err = ParseLine("missing comma")
	check("Reject malformed line", err != nil)

	var messages bytes.Buffer
	students, err := ReadStudents(
		strings.NewReader("Ava,88\nBad,101\nBo,72\n"),
		&messages,
	)
	check("Report bad line and continue",
		err == nil &&
			len(students) == 2 &&
			students[1].Name == "Bo" &&
			strings.Contains(messages.String(), "Line 2:"),
	)

	fmt.Fprintf(out, "\n%d passed, %d failed\n", passed, failed)
	if failed > 0 {
		return 1
	}
	return 0
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--test" {
		os.Exit(runChecks(os.Stdout))
	}

	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
