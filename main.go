package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

var Result []int
var HeapCh = make(chan int, 100_000_000)
var heap Heap
var set Set
var wgH sync.WaitGroup

func main() {

	s := flag.Bool("sort", false, "boolean value")
	u := flag.Bool("uniq", false, "boolean value")
	h := flag.Bool("heap", false, "boolean value")

	flag.Parse()
	args := len(os.Args)

	switch {
	case *s && *u && args == 3:
		fmt.Println("case uniq sort")
		runUniqSort()
	case *s && args == 2:
		fmt.Println("case sort")
		runSort()
	case *u && args == 2:
		fmt.Println("case uniq")
		runUniq()
	case *h && args == 2:
		fmt.Println("case heap")
		runHeap()
	default:
		fmt.Println("please add subcommand")
	}
}

func runSort() {
	RunReadAllSort()
	QSort(Result)
	CreateTxtWithQuickSort(Result)
}

func runHeap() {
	RunReadAllHeap()
	wgH.Add(1)
	go func() {
		defer wgH.Done()
		for data := range HeapCh {
			heap.Insert(data)
		}
	}()
	wgH.Wait()
	CreateTxtWithHeap(heap.Items)
}

func runUniq() {
	RunReadAllUniq()
}

func runUniqSort() {
	RunReadAllUniqSort()
	items := set.Items()
	QSort(items)
	CreateTxtWithQuickSort(items)
}
