package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/* Unit Test 1 for sorting */

func TestBubbleSort(t *testing.T) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	fmt.Println("Before sorting: ", elements)
	BubbleSort(elements)

	if elements[0] != 0 {
		t.Error("First index is o")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("Last index is 9")
	}
	fmt.Println("After sorting: ", elements)
}

func TestSort(t *testing.T) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	fmt.Println("Before sorting: ", elements)
	Sort(elements)

	if elements[0] != 0 {
		t.Error("First index is o")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("Last index is 9")
	}
	fmt.Println("After sorting: ", elements)
}

/* Benchmarks */

func BenchmarkBubbleSort(b *testing.B) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

/* Unit Test 2 for sorting */

func TestBubbleSortByAsserts(t *testing.T) {
	elements := GetElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
}

func TestAppendSlices(t *testing.T) {
	numbers := []int{}
	fmt.Println("Before appending:", numbers)
	numbers = AppendSlices(numbers, 1)
	numbers = AppendSlices(numbers, 2)
	numbers = AppendSlices(numbers, 3)
	numbers = AppendSlices(numbers, 4)

	if numbers[len(numbers)-1] != 4 {
		t.Error("Last index is 4")
	}
	fmt.Println("After appending:", numbers)
}
