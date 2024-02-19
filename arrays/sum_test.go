package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Calculates the sum of numbers in the array", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertCorrectSum(t, got, want, numbers)
	})

	t.Run("Incorrect sum", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 1}
		got := Sum(numbers)
		want := 15
		assertCorrectSum(t, got, want, numbers)
	})
}

func TestSumOrdinaryForLoop(t *testing.T) {
	t.Run("Calculates the sum of numbers in the array", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := SumOrdinaryForLoop(numbers)
		want := 15
		assertCorrectSum(t, got, want, numbers)
	})

	t.Run("Incorrect sum", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 1}
		got := SumOrdinaryForLoop(numbers)
		want := 15
		assertCorrectSum(t, got, want, numbers)
	})
}

func assertCorrectSum(t testing.TB, got, want int, numbers [5]int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
