package main

import (
	"encoding/csv"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type table struct {
	Date     string
	Open     string
	High     string
	Low      string
	Close    string
	Volume   string
	AdjClose string
}

func main() {
	// Open the CSV file
	file, err := os.Open("./table.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Create a slice to store the data
	var mainTable []table

	// Iterate through the CSV records and populate the data structure
	for _, record := range records {
		// Assuming the CSV columns are in the same order as the struct fields
		// age, _ := strconv.Atoi(record[1]) // Convert age to an integer
		table := table{
			Date:     record[0],
			Open:     record[1],
			High:     record[2],
			Low:      record[3],
			Close:    record[4],
			Volume:   record[5],
			AdjClose: record[6],
		}
		mainTable = append(mainTable, table)
	}

	// Now, 'people' contains the data from the CSV file in your data structure

	excutionErr := tpl.Execute(os.Stdout, mainTable)
	if err != nil {
		log.Fatalln(excutionErr)
	}
}
