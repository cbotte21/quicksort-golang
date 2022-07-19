/*
*	Test cases not yet implemented correctly.
*/

package main

import (
	"math/rand"
	"testing"
)

func TestQuicksort(t *testing.T) {
	quantityNums = 50
	rangeNums = 50
	var insertion[]int = make([]int, quantityNums, quantityNums)
	var quick[]int = make([]int, quantityNums, quantityNums)

	for i := 0; i < 100; i++ { //Test on 10 different arrays
		rand.Seed(int64(i))
		fillArr(insertion)
		copy(quick, insertion)
		insertionSort(insertion)
		quicksort(quick)
		
		for k := 0; k < quantityNums; k++ {
			if insertion[k] != quick[k] {
				t.Fatalf(`quicksort(arr) did not match insertionSort(arr)`)
			}
		}
	}
}
