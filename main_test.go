package main

import (
	"math/rand"
	"sort"
	"testing"
)

func TestKthLargest(t *testing.T) {
	cases := []struct {
		in   []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{[]int{1}, 1, 1},
		{[]int{7, 7, 7}, 3, 7},
		{[]int{-1, -5, 0}, 1, 0},
	}
	for _, c := range cases {
		got, ok := KthLargest(c.in, c.k)
		if !ok || got != c.want {
			t.Errorf("KthLargest(%v,%d) = (%d,%v), want %d", c.in, c.k, got, ok, c.want)
		}
	}
	if _, ok := KthLargest([]int{1}, 2); ok {
		t.Error("k beyond length must fail")
	}
	if _, ok := KthLargest(nil, 1); ok {
		t.Error("empty input must fail")
	}
}

// Property: agree with sorting across random inputs and every k.
func TestAgainstSort(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	for trial := 0; trial < 50; trial++ {
		n := 1 + rng.Intn(30)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.Intn(20) - 10
		}
		sorted := append([]int{}, nums...)
		sort.Ints(sorted)
		for k := 1; k <= n; k++ {
			got, ok := KthLargest(nums, k)
			if !ok || got != sorted[n-k] {
				t.Fatalf("trial %d: KthLargest(%v,%d) = %d, want %d", trial, nums, k, got, sorted[n-k])
			}
		}
	}
}
