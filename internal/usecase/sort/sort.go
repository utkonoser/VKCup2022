package sort

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

const (
	// concurrentThreshold - константа, при которой сортировка перестает выполняться конкурентно
	concurrentThreshold = 5000
	// quickSortThreshold - константа, при которой быстрая сортировка заменяется сортировкой вставкой
	quickSortThreshold = 30
)

// Result - слайс результирующих значений
var Result struct {
	Res []int
	sync.Mutex
}

// RunSort - запуск всех функций для реализации сортировки данных
func RunSort(path string) {
	RunReadAllSort(path)
	QSort(Result.Res)
	CreateTxtWithQuickSort(Result.Res, path)
}

// RunReadAllSort - функция читает значения из файлов в папке data и записывает из в один слайс
func RunReadAllSort(path string) {
	if _, err := os.Stat(path + "res.txt"); err == nil {
		err = os.Remove(path + "res.txt")
		if err != nil {
			log.Fatal(err)
		}
	}

	var wg sync.WaitGroup
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, txtFile := range files {
		txtFile := txtFile
		wg.Add(1)
		go func() {
			file, err := os.Open(path + txtFile.Name())
			if err != nil {
				log.Fatal(err)
			}
			defer func() {
				err = file.Close()
				if err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				num, err := strconv.Atoi(scanner.Text())
				if err != nil {
					log.Fatal(err)
				}
				Result.Lock()
				Result.Res = append(Result.Res, num)
				Result.Unlock()
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}()
	}
	wg.Wait()
}

// CreateTxtWithQuickSort - функция создает результирующий файл res.txt в папке data
func CreateTxtWithQuickSort(result []int, path string) {

	resTxt, err := os.Create(path + "res.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = resTxt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for _, i := range result {
		_, err = resTxt.WriteString(fmt.Sprintln(i))
		if err != nil {
			log.Fatal(err)
		}
	}
}

// InsertSort - функция сортировки вставкой, эффективна при малом количестве сортируемых элементов
func InsertSort(data []int) {
	i := 1
	for i < len(data) {
		h := data[i]
		j := i - 1
		for j >= 0 && h < data[j] {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = h
		i++
	}
}

// Partition - вспомогательная функция быстрой сортировки
func Partition(data []int) int {
	data[len(data)/2], data[0] = data[0], data[len(data)/2]
	pivot := data[0]
	mid := 0
	i := 1
	for i < len(data) {
		if data[i] < pivot {
			mid++
			data[i], data[mid] = data[mid], data[i]
		}
		i++
	}
	data[0], data[mid] = data[mid], data[0]
	return mid
}

// ConcurrentQuickSort - функция конкурентной быстрой сортировки данных
func ConcurrentQuickSort(data []int, wg *sync.WaitGroup) {
	for len(data) >= quickSortThreshold {
		mid := Partition(data)
		var portion []int
		if mid < len(data)/2 {
			portion = data[:mid]
			data = data[mid+1:]
		} else {
			portion = data[mid+1:]
			data = data[:mid]
		}
		if len(portion) > concurrentThreshold {
			wg.Add(1)
			go func(data []int) {
				defer wg.Done()
				ConcurrentQuickSort(data, wg)
			}(portion)
		} else {
			ConcurrentQuickSort(portion, wg)
		}
	}
	InsertSort(data)
}

// QSort - функция, запускающая сортировку
func QSort(data []int) {
	var wg sync.WaitGroup
	ConcurrentQuickSort(data, &wg)
	wg.Wait()
}
