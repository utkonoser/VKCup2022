package uniqSort

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

// TestRunUniqSort - тест, где проверяется корректность сортировки уникальных чисел из заданных данных
func TestRunUniqSort(t *testing.T) {
	var testRes []int
	dataPath := "../../../data/"
	RunUniqSort(dataPath)

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

	t.Run("unique test", func(t *testing.T) {
		for i, num := range testRes {
			if isInSlice(testRes[i+1:], num) {
				t.Error("result must be a unique set")
				return
			}
		}
		t.Log("success: result is unique set")
	})

	t.Run("sorting test", func(t *testing.T) {
		if !IsSorted(testRes) {
			t.Error("wrong result, slice must be sorted")
			return
		}
		t.Log("success: result is sorted set")
	})
}

// IsSorted - проверка отсортированы ли данные
func IsSorted(data []int) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

// isInSlice - функция проверяет уникален ли элемент в слайсе
func isInSlice(slice []int, num int) bool {
	for i := range slice {
		if slice[i] == num {
			return true
		}
	}
	return false
}
