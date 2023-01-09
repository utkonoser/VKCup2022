package main

import (
	"flag"
	"fmt"
	"goElimination/internal/usecase/heap"
	"goElimination/internal/usecase/sort"
	"goElimination/internal/usecase/uniq"
	"goElimination/internal/usecase/uniqSort"
	"os"
	"runtime/debug"
	"time"
)

// dataPath - относительный путь к папке, где располагаются считываемые данные
const dataPath = "../../data/"

// memoryLimit - число в байтах, сколько памяти может максимально потреблять приложение
const memoryLimit = 524288000

func main() {
	debug.SetMemoryLimit(memoryLimit)

	s := flag.Bool("sort", false, "boolean value")
	u := flag.Bool("uniq", false, "boolean value")
	h := flag.Bool("heap", false, "boolean value")

	flag.Parse()
	args := len(os.Args)

	switch {
	case *s && *u && args == 3:
		fmt.Println("Run case uniq sort...")
		start := time.Now()
		uniqSort.RunUniqSort(dataPath)
		fin := time.Since(start)
		fmt.Println("Elapsed time for case uniq sort: ", fin)
	case *s && args == 2:
		fmt.Println("Run case sort...")
		start := time.Now()
		sort.RunSort(dataPath)
		fin := time.Since(start)
		fmt.Println("Elapsed time for case sort: ", fin)
	case *u && args == 2:
		fmt.Println("Run case uniq...")
		start := time.Now()
		uniq.RunUniq(dataPath)
		fin := time.Since(start)
		fmt.Println("Elapsed time for case uniq: ", fin)
	case *h && args == 2:
		fmt.Println("Run case heap...")
		start := time.Now()
		heap.RunHeap(dataPath)
		fin := time.Since(start)
		fmt.Println("Elapsed time for case heap: ", fin)
	default:
		fmt.Println("Please add subcommand")
	}
}
