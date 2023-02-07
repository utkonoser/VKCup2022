package uniq

import (
	"bufio"
	"os"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

// TestRunUniq - полная проверка работоспособности кейса с уникальным множеством
func TestRunUniq(t *testing.T) {
	var testRes SetStruct
	dataPath := "../../../data/"
	RunUniq(dataPath)

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
		testRes.Insert(num)
	}
	if err := scanner.Err(); err != nil {
		t.Error(err)
	}

	items := testRes.Items().ShowItems()
	for i, num := range items {
		if isInSlice(items[i+1:], num) {
			t.Error("result must be a unique set")
			return
		}
	}

	t.Log("success: result is unique set")
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

// TestSet_Insert - тест, проверяющий правильную работу методов Insert и Items
func TestSet_InsertAndItems(t *testing.T) {
	testSlice := []int{9, 7, 6, 9, 7, 2}
	testSet := new(SetStruct)
	for _, num := range testSlice {
		testSet.Insert(num)
	}
	result := testSet.Items().ShowItems()
	sort.Ints(result)
	expected := []int{2, 6, 7, 9}

	resultBool := reflect.DeepEqual(expected, result)

	if !resultBool {
		t.Errorf("wrong result, %v != %v", expected, result)
	}
	t.Log("success: methods Insert and Items working correctly")

	t.Run("50_000_000 insertions in set", func(t *testing.T) {
		test2Set := new(SetStruct)
		for i := 0; i < 50_000_000; i++ {
			test2Set.Insert(1)
		}
		lengthSet := len(test2Set.Items().ShowItems())
		if lengthSet != 1 {
			t.Errorf("length test2Set must be 1, not %v", lengthSet)
		}
	})
}
