package utils

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-asargin-dev/models"
)

func CsvReader(fileName string) ([]models.Book, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var result []models.Book

	for _, row := range rows[1:] {
		id, _ := strconv.Atoi(row[0])
		stockNumber, _ := strconv.Atoi(row[2])
		pageNumber, _ := strconv.Atoi(row[3])
		price, _ := strconv.ParseFloat(row[4], 64)

		data := models.Book{
			ID:          id,
			Name:        row[1],
			StockNumber: stockNumber,
			PageNumber:  pageNumber,
			Price:       price,
			StockCode:   row[5],
			Isbn:        row[6],
			AuthorName:  row[7],
		}
		result = append(result, data)
	}
	return result, nil
}
