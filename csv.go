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
