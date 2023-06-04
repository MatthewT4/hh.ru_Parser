package excelExport

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"hh.ru_Parser/internal/http"
)

func vacancyStructToInterface(vacancy http.Vacancy) []interface{} {
	arr := make([]interface{}, 10)

	arr[0] = vacancy.Id
	arr[1] = vacancy.Name
	arr[2] = vacancy.Salary.Currency
	arr[3] = vacancy.Salary.From
	arr[4] = vacancy.Salary.To
	arr[5] = vacancy.Experience.Name
	arr[6] = vacancy.Schedule.Name
	arr[7] = vacancy.KeySkills
	arr[8] = vacancy.Description
	arr[9] = vacancy.Employer

	return arr
}

func WriteData(vacancies []http.Vacancy) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Data")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	f.SetSheetRow("Data", "A1", &[]interface{}{
		"Id",
		"Название",
		"Валюта",
		"ЗП от",
		"ЗП до",
		"Опыт",
		"График",
		"Скиллы",
		"Описание",
		"Компания",
	})

	for i := 0; i < len(vacancies); i++ {
		f.SetSheetRow("Data", fmt.Sprintf("A%v", i+2), &[]interface{}{
			vacancies[i].Id,
			vacancies[i].Name,
			vacancies[i].Salary.Currency,
			vacancies[i].Salary.From,
			vacancies[i].Salary.To,
			vacancies[i].Experience.Name,
			vacancies[i].Schedule.Name,
			vacancies[i].KeySkills,
			vacancies[i].Description,
			vacancies[i].Employer.Name,
		})
	}
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
