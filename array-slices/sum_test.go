package main

import "testing"

func TestSumAll(t *testing.T) {
	got := SumAll([]int{4, 4}, []int{0, 4})
	want := []int{8, 4}
	if got != want {
		t.Errof("got %v want %v", got, want)
	}
}
