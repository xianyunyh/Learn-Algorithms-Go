package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	a := assert.New(t)
	nums2 := []int{0, 0, 0}
	a.Equal(removeDuplicates(nums), 5)
	a.Equal(removeDuplicates(nums2), 1)
}

func TestContainsDuplicate(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	a := assert.New(t)
	a.Equal(containsDuplicate(nums), true)
	nums1 := []int{0, 1, 2, 3}
	a.Equal(containsDuplicate(nums1), false)
}

func TestSingleNum(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	a := assert.New(t)
	a.Equal(singleNumber(nums), 4)
}

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	a := assert.New(t)
	rotate(nums, 2)
	a.Equal(nums, []int{4, 5, 1, 2, 3})
}

func TestIntersect(t *testing.T) {
	nums := []int{4, 9, 5}
	nums1 := []int{9, 4, 9, 8, 4}
	a := assert.New(t)
	a.Equal(Intersect(nums, nums1), []int{4, 9})
	nums2 := []int{1, 2, 2, 1}
	nums3 := []int{2, 2}
	a.Equal(Intersect(nums2, nums3), []int{2, 2})
	nums4 := []int{54, 93, 21, 73, 84, 60, 18, 62, 59, 89, 83, 89, 25, 39, 41, 55, 78, 27, 65, 82, 94, 61, 12, 38, 76, 5, 35, 6, 51, 48, 61, 0, 47, 60, 84, 9, 13, 28, 38, 21, 55, 37, 4, 67, 64, 86, 45, 33, 41}
	nums5 := []int{17, 17, 87, 98, 18, 53, 2, 69, 74, 73, 20, 85, 59, 89, 84, 91, 84, 34, 44, 48, 20, 42, 68, 84, 8, 54, 66, 62, 69, 52, 67, 27, 87, 49, 92, 14, 92, 53, 22, 90, 60, 14, 8, 71, 0, 61, 94, 1, 22, 84, 10, 55, 55, 60, 98, 76, 27, 35, 84, 28, 4, 2, 9, 44, 86, 12, 17, 89, 35, 68, 17, 41, 21, 65, 59, 86, 42, 53, 0, 33, 80, 20}
	a.Equal(Intersect(nums4, nums5), []int{54, 21, 73, 84, 60, 18, 62, 59, 89, 89, 41, 55, 27, 65, 94, 61, 12, 76, 35, 48, 0, 60, 84, 9, 28, 55, 4, 67, 86, 33})
}

func TestPlusOne(t *testing.T) {
	nums1 := []int{9, 4, 9, 8, 4}
	a := assert.New(t)
	a.Equal(PlusOne(nums1), []int{9, 4, 9, 8, 5})
	a.Equal(PlusOne([]int{0}), []int{1})
	a.Equal(PlusOne([]int{9, 9, 9, 9}), []int{1, 0, 0, 0, 0})

}
