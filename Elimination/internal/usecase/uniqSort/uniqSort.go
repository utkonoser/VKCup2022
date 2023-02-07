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
	fmt.Println("Run case 'uniq sort'...")
	startCaseUniqSort := time.Now()

	SetVar.RunReadAllUniq(path)
	endReadAllFunc := time.Since(startCaseUniqSort)
	fmt.Println("End of reading all files and insert in set: ", endReadAllFunc)

	items := SetVar.Items().ShowItems()

	startQSort := time.Now()
	sort.QSort(items)
	endQSort := time.Since(startQSort)
	fmt.Println("End of Quick Sort of set: ", endQSort)

	startCreateTxt := time.Now()
	sort.CreateTxt(items, path)
	endCreateTxt := time.Since(startCreateTxt)
	fmt.Println("End of creating res.txt: ", endCreateTxt)

	endCaseUniqSort := time.Since(startCaseUniqSort)
	fmt.Println("Elapsed time for case 'uniq sort': ", endCaseUniqSort)
}
