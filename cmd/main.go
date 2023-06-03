package main

import (
	"fmt"
	"hh.ru_Parser/internal/http"
)

func main() {

	fmt.Println(http.GetVacancies("ProductManager", "7", "1", "73"))
}
