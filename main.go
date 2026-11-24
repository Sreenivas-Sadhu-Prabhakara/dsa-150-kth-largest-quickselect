package main

import (
	"fmt"
	"math/rand"
)

// KthLargest returns the k-th largest element via quickselect: partition
// around a random pivot and recurse into the single side that contains the
// target index. ok is false for invalid k. Average O(n), no heap, no sort.
func KthLargest(nums []int, k int) (int, bool) {
	if k < 1 || k > len(nums) {
		return 0, false
	}
	a := append([]int{}, nums...)
	target := len(a) - k // k-th largest = this index in ascending order
	lo, hi := 0, len(a)-1
	for {
		if lo == hi {
			return a[lo], true
		}
		p := partition(a, lo, hi)
		switch {
		case p == target:
			return a[p], true
		case p < target:
			lo = p + 1
		default:
			hi = p - 1
		}
	}
}

// partition places a random pivot into its final sorted position within
// a[lo..hi] and returns that position (Lomuto scheme).
func partition(a []int, lo, hi int) int {
	r := lo + rand.Intn(hi-lo+1)
	a[r], a[hi] = a[hi], a[r]
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}

func main() {
	v, _ := KthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	fmt.Println(v)
}
