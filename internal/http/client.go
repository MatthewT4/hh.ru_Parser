package http

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const url = "https://api.hh.ru"

type vacanciesList struct {
	Items []struct {
		Id string `json:"id"`
	} `json:"items"`
}

func GetVacancies(vacanciesName string, industry, area, professionalRole string) (vacanciesList, error) {

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", url+"/vacancies", nil)

	if err != nil {
		return vacanciesList{}, err
	}

	q := req.URL.Query()
	q.Add("text", vacanciesName)
	q.Add("industry", industry)
	q.Add("area", area)
	q.Add("professional_role", professionalRole)
	q.Add("per_page", "100")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return vacanciesList{}, err
	}
	defer resp.Body.Close()

	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			return vacanciesList{}, err
		}
	}

	var list vacanciesList
	errMarshal := json.Unmarshal(bodyBytes, &list)
	if errMarshal != nil {
		return vacanciesList{}, errMarshal
	}

	return list, nil
}

type Vacancy struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Employer    struct {
		Id   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"employer"`
	Experience struct {
		Id   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"experience"`
	Salary struct {
		Currency string `json:"currency,omitempty"`
		From     int    `json:"from,omitempty"`
		To       int    `json:"to,omitempty"`
	}
}

func GetVacancy(vacancyId string) (Vacancy, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", url+"/vacancy", nil)

	if err != nil {
		return Vacancy{}, err
	}

	q := req.URL.Query()
	q.Add("vacancy_id", vacancyId)

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return Vacancy{}, err
	}
	defer resp.Body.Close()

	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			return Vacancy{}, err
		}
	}

	var vacancy Vacancy
	errMarshal := json.Unmarshal(bodyBytes, &vacancy)
	if errMarshal != nil {
		return Vacancy{}, errMarshal
	}

	return vacancy, nil
}
