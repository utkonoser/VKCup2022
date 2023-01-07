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

func main() {

	s := flag.Bool("sort", false, "boolean value")
	u := flag.Bool("uniq", false, "boolean value")
	h := flag.Bool("heap", false, "boolean value")

	flag.Parse()
	args := len(os.Args)

	switch {
	case *s && *u && args == 3:
		fmt.Println("case uniq sort")
		uniqSort.RunUniqSort()
	case *s && args == 2:
		fmt.Println("case sort")
		sort.RunSort()
	case *u && args == 2:
		fmt.Println("case uniq")
		uniq.RunUniq()
	case *h && args == 2:
		fmt.Println("case heap")
		heap.RunHeap()
	default:
		fmt.Println("please add subcommand")
	}
}
