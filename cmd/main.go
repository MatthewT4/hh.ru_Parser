package main

import (
	"fmt"
	"hh.ru_Parser/internal/http"
)

func main() {

	vacancies, _ := http.GetVacancies("ProductManager", "7", "1", "73")
	fmt.Println(vacancies)

	for i := 0; i < len(vacancies.Items); i++ {
		fmt.Println(http.GetVacancy(vacancies.Items[i].Id))
	}
}
