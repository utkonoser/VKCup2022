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
)

// dataPath - относительный путь к папке, где располагаются считываемые данные
const dataPath = "../../data/"

// memoryLimit - число в байтах, сколько памяти может максимально потреблять приложение
const memoryLimit = 524288000

func main() {
	debug.SetMemoryLimit(memoryLimit)

	flagSort := flag.Bool("sort", false, "boolean value")
	flagUniq := flag.Bool("uniq", false, "boolean value")
	flagHeap := flag.Bool("heap", false, "boolean value")

	flag.Parse()
	args := len(os.Args)

	switch {
	case *flagSort && *flagUniq && args == 3:
		uniqSort.RunUniqSort(dataPath)

	case *flagSort && args == 2:
		sort.RunSort(dataPath)

	case *flagUniq && args == 2:
		uniq.RunUniq(dataPath)

	case *flagHeap && args == 2:
		heap.RunHeap(dataPath)

	default:
		fmt.Println("Please add subcommand!")
	}
}
