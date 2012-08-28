package sleepsort

import (
	"math/rand"
	"sort"
	"testing"
	"github.com/bmizerany/assert"
)

func TestAsc(t *testing.T) {
	nums := []int{ 1, 2, 3, 3, 4, 5 }
	ref := copySlice(nums)
	sort.Ints(ref)
	SortSlice(nums)
	assertAsc(t, nums)
	assertEqualSlice(t, ref, nums)
}

func TestDesc(t *testing.T) {
	nums := []int{ 5, 4, 4, 3, 2, 1 }
	ref := copySlice(nums)
	sort.Ints(ref)
	SortSlice(nums)
	assertAsc(t, nums)
	assertEqualSlice(t, ref, nums)
}

func TestJitter(t *testing.T) {
	for trial := 0; trial < 3; trial++ {
		nums := []int{}
		for i := 0; i < 30; i++ {
			nums = append(nums, rand.Int() % 500)
		}
		ref := copySlice(nums)
		sort.Ints(ref)
		SortSlice(nums)
		assertAsc(t, nums)
		assertEqualSlice(t, ref, nums)
	}
}

func copySlice(p []int) []int {
	q := []int{}
	for _, n := range p {
		q = append(q, n)
	}
	return q
}

func assertAsc(t *testing.T, nums []int) {
	for i, n := range nums {
		if i > 0 {
			assert.Tf(t, nums[i-1] <= n, "Not in ascending order: %v", nums)
		}
	}
}

func assertEqualSlice(t *testing.T, expect []int, actual []int) {
	assert.Equalf(t, len(expect), len(actual), "unequal lengths")
	for i := 0; i < len(expect); i++ {
		assert.Equalf(t, expect[i], actual[i], "at index %d", i)
	}
}
