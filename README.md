# quicksort-golang

My attempt at an optimized implementation of quicksort. This is an adaptive algorith which resorts to insertionsort when a slice is under a specific length. Also gets the pivot from a median of three partition.

The project includes a test case which compares equality between 1000 arrays sorted using quicksort and insertion sort. Each array contains 50 integers ranging from 0-2500.

How-to use:
  1) Compile to independent executable (LINUX)
    - go build && ./optimized-quicksort.exe
  2) Run via go runtime
    - go run main.go
  3) Run tests
    - go test
