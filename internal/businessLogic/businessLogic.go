package businessLogic

import (
	"fmt"
	"hh.ru_Parser/internal/excelExport"
	"hh.ru_Parser/internal/http"
	"sync"
)

func Parse() {
	vacancies, _ := http.GetVacancies("ProductManager", "7", "1", "73")

	resultChannel := make(chan http.Vacancy, len(vacancies.Items))
	wg := &sync.WaitGroup{}

	for i := 0; i < len(vacancies.Items); i++ {
		wg.Add(1)
		go http.AsyncGetVacancy(vacancies.Items[i].Id, resultChannel, wg)
	}

	wg.Wait()

	vacs := make([]http.Vacancy, 0, len(vacancies.Items))

	close(resultChannel)

	for i := range resultChannel {
		vacs = append(vacs, i)
	}

	fmt.Println(vacs)

	excelExport.WriteData(vacs)
}
