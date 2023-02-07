package sort

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	// concurrentThreshold - константа, при которой сортировка перестает выполняться конкурентно
	concurrentThreshold = 5000
	// quickSortThreshold - константа, при которой быстрая сортировка заменяется сортировкой вставкой
	quickSortThreshold = 30
)

// CreateTxt - функция создает результирующий файл res.txt в папке data
func CreateTxt(r []int, path string) {

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

	var sliceS []string

	for _, n := range r {
		sliceS = append(sliceS, strconv.Itoa(n))
	}

	s := strings.Join(sliceS, "\n")

	_, err = resTxt.WriteString(s)
	if err != nil {
		log.Fatal(err)
	}
}

// RunSort - запуск всех функций для реализации сортировки данных
func RunSort(path string) {
	var r Result

	fmt.Println("Run case 'sort'...")
	startCaseSort := time.Now()

	r.RunReadAll(path)
	endRunReadAllFunc := time.Since(startCaseSort)
	fmt.Println("End of reading all files: ", endRunReadAllFunc)

	startQSort := time.Now()
	QSort(r.Res)
	endQSort := time.Since(startQSort)
	fmt.Println("End of Quick Sort: ", endQSort)

	startCreateTxt := time.Now()
	CreateTxt(r.Res, path)
	endCreateTxt := time.Since(startCreateTxt)
	fmt.Println("End of creating res.txt: ", endCreateTxt)

	endCaseSort := time.Since(startCaseSort)
	fmt.Println("Elapsed time for case 'sort': ", endCaseSort)
}

// SliceItems - интерфейс результирующего списка
type SliceItems interface {
	RunReadAll(path string)
	ShowItems() []int
}

// Result - структура результирующих значений
type Result struct {
	Res []int
	sync.Mutex
}

// ShowItems - метод, возвращающий слайс структуры
func (r *Result) ShowItems() []int {
	return r.Res
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

// RunReadAll - функция читает значения из файлов в нужной папке и записывает из в один слайс
func (r *Result) RunReadAll(path string) {
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
				r.Lock()
				r.Res = append(r.Res, num)
				r.Unlock()
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}()
	}
	wg.Wait()
}
