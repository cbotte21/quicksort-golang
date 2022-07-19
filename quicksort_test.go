/*
*	Test cases not yet implemented correctly.
*/

package main

import (
	"math/rand"
	"testing"
	"fmt"
)

func TestQuicksort(t *testing.T) {
	quantityNums = 50
	rangeNums = 50
	var arr[]int = make([]int, quantityNums, quantityNums)

	for i := 0; i < 10; i++ { //Test on 10 different arrays
		rand.Seed(int64(i))
		fillArr(arr)
		fmt.Println(arr)
		insertion := insertionSort(arr)
		quick := quicksort(arr)

		if insertion != quick {
			t.Fatalf("quicksort(arr) did not match insertionSort(arr)")
		}
	}
}
