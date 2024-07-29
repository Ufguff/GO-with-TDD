package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, nums - %v", got, want, numbers)
		}
	})
	/*
		t.Run("sum of any numbers", func(t *testing.T) {
				numbers := []int{1, 2, 3}

				got := Sum(numbers)
				want := 6

				if got != want {
					t.Errorf("got %d, want %d, nums - %v", got, want, numbers)
				}
			})
	*/
}
