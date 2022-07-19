/*

Title: Optimized Quicksort
Language: Golang
Author: Cody Botte
Goal: I plan on creating a pretty optimized quicksort implementation. I plan on including actions such as a 3 or 5 int median to find an initial pivot point. Insertion sort on small arrays. And such minor improvements.

This is an adaptive algorithm.

*/

package main

import (
	"fmt"
	"math/rand"
)

var quantityNums int
var rangeNums int

func main() {
	//Credits
	fmt.Println("An optimized quicksort programmed in golang.\nBy: Cody Botte")
	fmt.Println()

	//Variable initializaions
	rand.Seed(5)
	quantityNums = 50
	rangeNums = 50
	var arr[]int = make([]int, quantityNums, quantityNums)

	//Init array
	fillArr(arr)

	//Output
	fmt.Print("Initial array: ")
	fmt.Println(arr)
	fmt.Println()

	//Sort
	quicksort(arr)

	//Output
	fmt.Print("Sorted array: ")
	fmt.Println(arr)
}

func fillArr(arr []int) {
	i := 0
	for ; i < cap(arr); i++ {
		arr[i] = rand.Intn(rangeNums+1)
	}
}

func partition(arr []int) int {
	midway := len(arr)/2
	left := midway-midway/2
	right := midway+midway/2
	tmp := []int{left, midway, right} //The easy way out.
	insertionSort(tmp)
	return tmp[1]
}

func insertionSort(arr []int) {
	insertionSort_recur(arr, 0)
}

func insertionSort_recur(arr []int, offset int) {
	//Terminate if sorted
	if (offset == len(arr)) {
		return
	}
	//Find lowest index
	minIndex := offset
	i := offset
	for ; i < len(arr); i++ {
		if arr[minIndex] > arr[i] {
			minIndex = i
		}
	}
	//Swap
	tmp := arr[minIndex]
	arr[minIndex] = arr[offset]
	arr[offset] = tmp
	//Recursive call
	insertionSort_recur(arr, offset+1)
}

func quicksort(arr []int) {
	if len(arr) <= 5 { //Optimization
		insertionSort(arr)
		return
	}

	pivot := partition(arr)
	l, r := 0, len(arr)-1

	arr[pivot], arr[r] = arr[r], arr[pivot]
	pivot = r
	r--
	
	//Loop
	for l < r {
		if arr[l] <= arr[pivot] {
			l++
		}
		if arr[r] >= arr[pivot] {
			r--
		}
		if arr[l] > arr[pivot] && arr[r] < arr[pivot] {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	
	if (arr[l] > arr[pivot]) {
		arr[pivot], arr[l] = arr[l], arr[pivot]
	}

	if (l == 0) { //Prevents out of bounds error where l is max the whole time.
		quicksort(arr)
		return
	}
	//Recursive call
	quicksort(arr[:l+1])
	quicksort(arr[l:])
}
