package main

import (
	"flag"
	"fmt"
	"goElimination/internal/usecase/heap"
	"goElimination/internal/usecase/sort"
	"goElimination/internal/usecase/uniq"
	"goElimination/internal/usecase/uniqSort"
	"os"
)

// dataPath - относительный путь к папке, где располагаются считываемые данные
const dataPath = "../../data/"

func main() {

	s := flag.Bool("sort", false, "boolean value")
	u := flag.Bool("uniq", false, "boolean value")
	h := flag.Bool("heap", false, "boolean value")

	flag.Parse()
	args := len(os.Args)

	switch {
	case *s && *u && args == 3:
		fmt.Println("Run case uniq sort...")
		uniqSort.RunUniqSort(dataPath)
	case *s && args == 2:
		fmt.Println("Run case sort...")
		sort.RunSort(dataPath)
	case *u && args == 2:
		fmt.Println("Run case uniq...")
		uniq.RunUniq(dataPath)
	case *h && args == 2:
		fmt.Println("Run case heap...")
		heap.RunHeap(dataPath)
	default:
		fmt.Println("Please add subcommand")
	}
}
