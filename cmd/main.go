package main

import "hh.ru_Parser/internal/businessLogic"

func main() {

	//vacancies, _ := http.GetVacancies("ProductManager", "7", "1", "73")
	//fmt.Println(vacancies)
	//
	//for i := 0; i < len(vacancies.Items); i++ {
	//	fmt.Println(http.GetVacancy(vacancies.Items[i].Id))
	//}
	businessLogic.Parse()
}
