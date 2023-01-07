package heap

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

var WgH sync.WaitGroup
var HeapCh = make(chan int, 100_000_000)
var HeapVar Heap

func RunHeap() {
	RunReadAllHeap()
	WgH.Add(1)
	go func() {
		defer WgH.Done()
		for data := range HeapCh {
			HeapVar.Insert(data)
		}
	}()
	WgH.Wait()
	CreateTxtWithHeap(HeapVar.Items)
}

func RunReadAllHeap() {
	if _, err := os.Stat("../../data/res.txt"); err == nil {
		err = os.Remove("../../data/res.txt")
		if err != nil {
			log.Fatal(err)
		}
	}

	var wg sync.WaitGroup
	files, err := os.ReadDir("../../data/")
	if err != nil {
		log.Fatal(err)
	}
	for _, txtFile := range files {
		txtFile := txtFile
		wg.Add(1)
		go func() {
			file, err := os.Open("../../data/" + txtFile.Name())
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

type Heap struct {
	Items []int
}

func (h *Heap) Swap(index1, index2 int) {
	h.Items[index1], h.Items[index2] = h.Items[index2], h.Items[index1]
}

func (h *Heap) Insert(value int) {
	h.Items = append(h.Items, value)
	h.buildHeap(len(h.Items) - 1)
}

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

func CreateTxtWithHeap(items []int) {

	resTxt, err := os.Create("../../data/res.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = resTxt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for _, i := range items {
		_, err = resTxt.WriteString(fmt.Sprintln(i))
		if err != nil {
			log.Fatal(err)
		}
	}
}
