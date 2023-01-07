package sort

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

// IsSorted - проверка отсортированы ли данные
func IsSorted(data []int) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func TestRunSort(t *testing.T) {
	var testRes []int
	dataPath := "../../../data/"
	RunSort(dataPath)

	file, err := os.Open(dataPath + "res.txt")
	if err != nil {
		t.Error("file 'res.txt' doesn't open")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			t.Error(err)
		}
		testRes = append(testRes, num)
	}
	if err := scanner.Err(); err != nil {
		t.Error(err)
	}

	if !IsSorted(testRes) {
		t.Log("wrong result, slice must be sorted")
	}
}

func TestInsertSort(t *testing.T) {
	var testSlice = []int{9, 7, 6, 8, 3, 2, 4, 5, 1}

	InsertSort(testSlice)

	if !IsSorted(testSlice) {
		t.Log("wrong result, slice must be sorted")
	}

	testSlice = append(testSlice, -100)

	if IsSorted(testSlice) {
		t.Log("wrong result, slice must be not sorted")
	}
}

func TestQSort(t *testing.T) {
	sortedBool := false
	size := 50_000_000
	testBigData := make([]int, size)

	for i := 0; i < size; i++ {
		testBigData[i] = 100 * rand.Int()
	}
	sortedBool = IsSorted(testBigData)
	t.Log("Before testing QSort\nIs sorted: ", sortedBool)

	start := time.Now()
	QSort(testBigData)
	elapsed := time.Since(start)

	sortedBool = IsSorted(testBigData)
	if !sortedBool {
		t.Log("wrong result, slice must be sorted")
	}

	t.Log("Elapsed time for concurrent quicksort = ", elapsed)
	t.Log("After testing QSort\nIs sorted: ", sortedBool)
}
