package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
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
	} `json:"salary"`
	KeySkills []struct {
		Name string `json:"name"`
	} `json:"key_skills"`
	Schedule struct {
		Name string `json:"name"`
	} `json:"schedule"`
}

func AsyncGetVacancy(vacancyId string, result chan Vacancy, wg *sync.WaitGroup) {
	vac, err := getVacancy(vacancyId)
	if err == nil {
		result <- vac
	}
	wg.Done()
}

func getVacancy(vacancyId string) (Vacancy, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", url+fmt.Sprintf("/vacancies/%v", vacancyId), nil)

	if err != nil {
		return Vacancy{}, err
	}

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
