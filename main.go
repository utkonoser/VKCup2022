package main

import (
	"sync"
)

var Result []int
var HeapCh = make(chan int, 100_000_000)
var heap Heap
var wgH sync.WaitGroup

func main() {
	//RunReadAllSort()
	//fmt.Println("before sorting: ",Result)
	//QSort(Result)
	//fmt.Println("after quicksort: ", Result)
	//CreateTxtWithQuickSort()

	//RunReadAllHeap()
	//wgH.Add(1)
	//go func() {
	//	defer wgH.Done()
	//	for data := range HeapCh {
	//		heap.Insert(data)
	//	}
	//}()
	//wgH.Wait()
	//fmt.Println(heap.Items)
	//CreateTxtWithHeap(heap.Items)
}
