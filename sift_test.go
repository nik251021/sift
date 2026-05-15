package sift

import (
	"testing"
)

func TestFinishers(t *testing.T) {
	data := []int{10, 20, 30}

	t.Run("Distinct: Remove duplicates", func(t *testing.T) {
		dups := []int{1, 2, 2, 3, 1, 4, 4, 5}
		expected := []int{1, 2, 3, 4, 5}

		res := From(dups).Distinct().ToSlice()

		if len(res) != len(expected) {
			t.Errorf("Distinct failed: expected length %d, got %d", len(expected), len(res))
		}

		for i, v := range res {
			if v != expected[i] {
				t.Errorf("Distinct failed at index %d: expected %d, got %d", i, expected[i], v)
			}
		}

		strDups := []string{"go", "rust", "go", "java"}
		strRes := From(strDups).Distinct().ToSlice()

		if len(strRes) != 3 {
			t.Errorf("Distinct (strings) failed: expected 3 elements, got %d", len(strRes))
		}
	})
	t.Run("Select: Transformation", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		qNum := From(numbers)

		doubleQuery := Select(qNum, func(n int) int {
			return n * 2
		})

		resNum := doubleQuery.ToSlice()
		if resNum[0] != 2 || resNum[2] != 6 {
			t.Errorf("Select (int->int) failed: got %v", resNum)
		}

		strQuery := Select(qNum, func(n int) string {
			if n == 1 {
				return "one"
			}
			return "other"
		})

		resStr := strQuery.ToSlice()
		if resStr[0] != "one" || resStr[1] != "other" {
			t.Errorf("Select (int->string) failed: got %v", resStr)
		}
	})

	t.Run("Select: Struct to Field", func(t *testing.T) {
		type User struct {
			ID   int
			Name string
		}

		users := []User{
			{ID: 1, Name: "Ivan"},
			{ID: 2, Name: "Oleg"},
		}

		q := From(users)

		namesQuery := Select(q, func(u User) string {
			return u.Name
		})

		names := namesQuery.ToSlice()
		if len(names) != 2 || names[0] != "Ivan" || names[1] != "Oleg" {
			t.Errorf("Select (Struct field) failed: got %v", names)
		}
	})
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
