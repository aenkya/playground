package algo

import (
	"fmt"
	"sort"

	pio "enkya.org/playground/practice/io"
)

type Median2SortedArrays struct {
	description string
	examples    []pio.IO
	testData    []pio.IO
	versions    []func(nums1 []int, nums2 []int) float64
}

func (m *Median2SortedArrays) Median2SortedArraysV1(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	if len(nums1)%2 == 0 {
		return float64(nums1[len(nums1)/2-1]+nums1[len(nums1)/2]) / 2
	}

	return float64(nums1[len(nums1)/2])
}

func (m *Median2SortedArrays) Median2SortedArraysV2(nums1, nums2 []int) float64 {
	totalNums := len(nums1) + len(nums2)
	medianIndex := totalNums / 2

	if totalNums%2 == 0 {
		first := m.findKthSmallest(nums1, nums2, medianIndex)
		second := m.findKthSmallest(nums1, nums2, medianIndex+1)

		return float64(first+second) / 2
	}

	return float64(m.findKthSmallest(nums1, nums2, medianIndex+1))
}

func (m *Median2SortedArrays) findKthSmallest(nums1, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}

	if len(nums2) == 0 {
		return nums1[k-1]
	}

	if k == 1 {
		return min(nums1[0], nums2[0])
	}

	mid1 := min(len(nums1), k/2)
	mid2 := min(len(nums2), k/2)

	if nums1[mid1-1] < nums2[mid2-1] {
		return m.findKthSmallest(nums1[mid1:], nums2, k-mid1)
	}

	return m.findKthSmallest(nums1, nums2[mid2:], k-mid2)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func (m *Median2SortedArrays) RunAlgo() {
	for _, v := range m.versions {
		for _, d := range m.testData {
			arr1 := d.Input.([][]int)[0]
			arr2 := d.Input.([][]int)[1]
			fmt.Println(v(arr1, arr2))
		}
	}
}

func (m *Median2SortedArrays) LoadTestData() {
	m.testData = []pio.IO{
		{Input: [][]int{{1, 3}, {2}}, Output: 2.0},
		{Input: [][]int{{1, 2}, {3, 4}}, Output: 2.5},
		{Input: [][]int{{0, 0}, {0, 0}}, Output: 0.0},
		{Input: [][]int{{}, {1}}, Output: 1.0},
		{Input: [][]int{{2}, {}}, Output: 2.0},
	}
}

func (m *Median2SortedArrays) Describe() {
	fmt.Printf("\nDescription: %s\n", m.description)
	fmt.Println("Examples:")

	for _, e := range m.examples {
		fmt.Printf("\tInput: %v\n\tOutput: %v\n", e.Input, e.Output)
	}
}

func NewMedian2SortedArrays() *Median2SortedArrays {
	m := &Median2SortedArrays{
		description: "Median of Two Sorted Arrays",
		examples: []pio.IO{
			{Input: [][]int{{1, 3}, {2}}, Output: 2.0},
			{Input: [][]int{{1, 2}, {3, 4}}, Output: 2.5},
			{Input: [][]int{{0, 0}, {0, 0}}, Output: 0.0},
			{Input: [][]int{{}, {1}}, Output: 1.0},
			{Input: [][]int{{2}, {}}, Output: 2.0},
		},
	}

	m.versions = []func(nums1 []int, nums2 []int) float64{
		m.Median2SortedArraysV1,
		m.Median2SortedArraysV2,
	}

	m.LoadTestData()

	return m
}
