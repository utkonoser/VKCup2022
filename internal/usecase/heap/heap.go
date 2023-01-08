package heap

import (
	"bufio"
	"goElimination/internal/usecase/sort"
	"log"
	"os"
	"strconv"
	"sync"
)

var WgH sync.WaitGroup
var HeapCh = make(chan int, 100_000_000)
var HeapVar Heap

// HeapInterface - интерфейс кучи
type HeapInterface interface {
	Swap(index1, index2 int)
	Insert(value int)
	buildHeap(index int)
}

// RunHeap - функция, отвечающая за корректную работу кейса с кучей
func RunHeap(path string) {
	RunReadAllHeap(path)
	WgH.Add(1)
	go func() {
		defer WgH.Done()
		for data := range HeapCh {
			HeapVar.Insert(data)
		}
	}()
	WgH.Wait()
	sort.CreateTxt(HeapVar.Items, path)
}

// RunReadAllHeap - функция, которая читает данные из нужной папки и передает их через канал в другую горутину
func RunReadAllHeap(path string) {
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
				wg.Add(1)
				go func() {
					HeapCh <- num
					defer wg.Done()
				}()
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}()
	}
	wg.Wait()
	close(HeapCh)
}

// Heap - реализация кучи
type Heap struct {
	Items []int
}

// Swap - метод Heap, который меняет местами два элемента кучи
func (h *Heap) Swap(index1, index2 int) {
	h.Items[index1], h.Items[index2] = h.Items[index2], h.Items[index1]
}

// Insert - метод Heap, который вставляет новый элемент в кучу
func (h *Heap) Insert(value int) {
	h.Items = append(h.Items, value)
	h.buildHeap(len(h.Items) - 1)
}

// buildHeap - отвечает за корректную сборку кучи
func (h *Heap) buildHeap(index int) {
	var parent int
	if index > 0 {
		parent = (index - 1) / 2
		if h.Items[index] > h.Items[parent] {
			h.Swap(index, parent)
		}
		h.buildHeap(parent)
	}
}
