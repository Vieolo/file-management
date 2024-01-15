package filemanagement

import (
	"encoding/csv"
	"os"
)

// Reads a CSV files and converts to an array of array of strings
func ReadCsvFile(filePath string) ([][]string, error) {
	f, fErr := os.Open(filePath)
	if fErr != nil {
		return make([][]string, 0), fErr
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, rErr := csvReader.ReadAll()
	if rErr != nil {
		return make([][]string, 0), rErr
	}

	return records, nil
}

// Receives the content of the CSV file as a slice of string slice and writes it to file
func WriteCsvFile(content [][]string, filename string) error {
	file, oErr := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if oErr != nil {
		return oErr
	}

	csvWriter := csv.NewWriter(file)
	for _, line := range content {

		wErr := csvWriter.Write(line)
		if wErr != nil {
			return wErr
		}
	}
	csvWriter.Flush()

	return file.Close()
}
