package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"sort"
)

func main() {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	elements = remove(elements, 8)
	indext := Find(elements, 4)
	fmt.Println(indext)
	elements = cryptographic(elements)
	elements = DeleteAll(elements)

}

/* Sorting */
func BubbleSort(elements []int) {
	keepWorking := true
	for keepWorking {
		keepWorking = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				keepWorking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}

func Sort(elements []int) {
	sort.Ints(elements)
}

func GetElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}

/* Append slices */
func AppendSlices(numbers []int, elem int) []int {
	numbers = append(numbers, elem)
	return numbers
}

func printSlice(x []int) {
	fmt.Printf("slice=%v\n", x)
}

func remove(s []int, i int) []int {
	fmt.Println("Deleting item at index 3:")
	fmt.Println(append(s[:i], s[i+1:]...))
	s = append(s[:i], s[i+1:]...)
	fmt.Println(s)
	fmt.Println("--------------")
	return s
}

func Find(a []int, x int) int {
	fmt.Println("Indexes:")

	for i, n := range a {
		if x == n {
			fmt.Println("Index of", x, "in the slice is: ", i)
			fmt.Println("--------------")
			return i
		}
	}
	fmt.Println("Index in the slice is: ", len(a))
	fmt.Println("--------------")
	return len(a)

}

func DeleteAll(a []int) []int {
	fmt.Println("Deleting all:")
	a = nil
	fmt.Println("Cleared array is:", a)
	fmt.Println("--------------")
	return a

}

func cryptographic(a []int) []int {
	fmt.Println("Encrypting the slice elements:")

	for i := 0; i < len(a); i++ {

		hash, _ := bcrypt.GenerateFromPassword([]byte("a[i]"), 10)
		fmt.Println(hash)

	}
	fmt.Println("--------------")
	return nil
}
