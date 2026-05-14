package sift

import (
	"testing"
)

func TestFinishers(t *testing.T) {
	data := []int{10, 20, 30}

	t.Run("Any/All", func(t *testing.T) {
		q := From(data)
		if !q.Any(func(n int) bool { return n == 20 }) {
			t.Error("Any failed: should find 20")
		}
		if q.All(func(n int) bool { return n > 50 }) {
			t.Error("All failed: not all numbers are > 50")
		}
	})

	t.Run("First/Last", func(t *testing.T) {
		q := From(data)

		valF, okF := q.First()
		if !okF || valF != 10 {
			t.Errorf("First failed: got %v, %v", valF, okF)
		}

		valL, okL := q.Last()
		if !okL || valL != 30 {
			t.Errorf("Last failed: got %v, %v", valL, okL)
		}

		valEmpty, okEmpty := From([]int{}).First()
		if okEmpty {
			t.Error("First should return ok=false for empty slice")
		}
		if valEmpty != 0 {
			t.Error("First should return zero value for empty slice")
		}
	})

	t.Run("Find", func(t *testing.T) {
		val, ok := From(data).Find(func(n int) bool { return n > 15 })
		if !ok || val != 20 {
			t.Errorf("Find failed: expected 20, got %v", val)
		}
	})

	t.Run("At: Bounds check", func(t *testing.T) {
		q := From(data)

		val, ok := q.At(1)
		if !ok || val != 20 {
			t.Errorf("At(1) failed: got %v", val)
		}

		_, okHigh := q.At(3)
		if okHigh {
			t.Error("At(3) should be out of range")
		}

		_, okLow := q.At(-1)
		if okLow {
			t.Error("At(-1) should be out of range")
		}
	})
}
