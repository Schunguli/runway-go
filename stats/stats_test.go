package main

import "testing"

func TestStats(t *testing.T) {
    t.Run("stats of a slice of numbers", func(t *testing.T) {
        nums := []float64{1, 2, 3, 4, 5}
        min, maximum, mean := Stats(nums)

        if min != 1 || maximum != 5 || mean != 3 {
            t.Errorf("got min=%v maximum=%v mean=%v, want min=1 maximum=5 mean=3", min, maximum, mean)
        }
    })

    t.Run("empty slice returns zeros", func(t *testing.T) {
        nums := []float64{}
        min, maximum, mean := Stats(nums)

        if min != 0 || maximum != 0 || mean != 0 {
            t.Errorf("got min=%v maximum=%v mean=%v, want all zeros", min, maximum, mean)
        }
    })
}
