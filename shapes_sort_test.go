package main

import "testing"

func TestSortShapesByArea(t *testing.T) {
	shapes := []Shape{
		Rectangle{4, 5}, // 20
		Triangle{4, 3},  // 6
		Circle{1},       // 3.14
	}

	SortShapesByArea(shapes)

	expected := []float64{
		3.141592653589793,
		6,
		20,
	}

	for i := range shapes {
		if shapes[i].Area() != expected[i] {
			t.Errorf("at index %d got %f want %f",
				i, shapes[i].Area(), expected[i])
		}
	}
}