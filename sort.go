package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

const (
	concurrentThreshold = 5000
	quickSortThreshold  = 30
)

func RunReadAllSort() {
	if _, err := os.Stat("data/res.txt"); err == nil {
		err = os.Remove("data/res.txt")
		if err != nil {
			log.Fatal(err)
		}
	}

	var wg sync.WaitGroup
	files, err := os.ReadDir("data/")
	if err != nil {
		log.Fatal(err)
	}
	for _, txtFile := range files {
		txtFile := txtFile
		wg.Add(1)
		go func() {
			file, err := os.Open("data/" + txtFile.Name())
			if err != nil {
				log.Fatal(err)
			}
			defer func() {
				file.Close()
				wg.Done()
			}()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				num, err := strconv.Atoi(scanner.Text())
				if err != nil {
					log.Fatal(err)
				}
				Result = append(Result, num)
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}()
	}
	wg.Wait()
}

func CreateTxtWithQuickSort(result []int) {

	resTxt, err := os.Create("data/res.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer resTxt.Close()

	for _, i := range result {
		_, err = resTxt.WriteString(fmt.Sprintln(i))
		if err != nil {
			log.Fatal(err)
		}
	}
}

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

func IsSorted(data []int) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

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

func QSort(data []int) {
	var wg sync.WaitGroup
	ConcurrentQuickSort(data, &wg)
	wg.Wait()
}
