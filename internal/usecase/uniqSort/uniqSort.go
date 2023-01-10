package uniqSort

import (
	"fmt"
	"goElimination/internal/usecase/sort"
	"goElimination/internal/usecase/uniq"
	"time"
)

// SetVar - множество уникальных значений
var SetVar uniq.SetStruct

// RunUniqSort - основная функция кейса с сортировкой уникальных значений
func RunUniqSort(path string) {
	start := time.Now()
	SetVar.RunReadAllUniq(path)
	finRead := time.Since(start)
	fmt.Println("End of reading all files and insert in set: ", finRead)
	items := SetVar.Items().ShowItems()
	start = time.Now()
	sort.QSort(items)
	finQSort := time.Since(start)
	fmt.Println("End of Quick Sort of set: ", finQSort)
	start = time.Now()
	sort.CreateTxt(items, path)
	finTxt := time.Since(start)
	fmt.Println("End of creating res.txt: ", finTxt)
}
