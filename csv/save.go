package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Save data to CSV
func Save(filename string, data [][]string, delimiter rune, overwrite bool) error {
	if !overwrite {
		if _, err := os.Stat(filename); err == nil {
			return fmt.Errorf("File %s already exists", filename)
		}
	}
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Creating CSV error: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = delimiter

	// write and flush
	if err := writer.WriteAll(data); err != nil {
		return fmt.Errorf("Writing CSV error: %w", err)
	}
	return nil
}
