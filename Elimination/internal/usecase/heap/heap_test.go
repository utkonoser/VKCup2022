package heap

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

// TestRunHeap - тест, проверяющий полную работоспособность кейса с кучей
func TestRunHeap(t *testing.T) {
	var testRes []int
	dataPath := "../../../data/"
	RunHeap(dataPath)

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

	node := 0
	if !IsHeap(testRes, node) {
		t.Log("wrong result, slice must be heap")
	}

}

// IsHeap - функция, проверяющая является ли слайс кучей
func IsHeap(items []int, idx int) bool {
	if 2*idx+2 > len(items) {
		return true
	}
	left := (items[idx] >= items[2*idx+1]) &&
		IsHeap(items, 2*idx+1)
	right := (2*idx+2 == len(items)) ||
		(items[idx] >= items[2*idx+2] && IsHeap(items, 2*idx+2))
	return left && right
}

// TestIsHeap - тест, тестирующий функцию для проверки кучи
func TestIsHeap(t *testing.T) {
	node := 0

	test1 := []int{0, 0, 0, 0, 0}
	if !IsHeap(test1, node) {
		t.Log("wrong result, slice must be heap")
	}

	test2 := []int{1, 2, 3, 4, 5}
	if IsHeap(test2, node) {
		t.Error("wrong result, slice is not heap")
	}

	test3 := []int{5, 2, 3, 1, 2}
	if !IsHeap(test3, node) {
		t.Log("wrong result, slice must be heap")
	}
}
