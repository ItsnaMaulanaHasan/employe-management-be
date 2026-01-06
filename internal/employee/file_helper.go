package employee

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const reportDir = "storage/reports"

func SaveJSONToFile(filename string, data any) error {
	err := os.MkdirAll(reportDir, 0755)
	if err != nil {
		return err
	}

	filepath := filepath.Join(reportDir, filename)

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, bytes, 0644)
}

func ReadJSONFromFile(filename string) ([]byte, error) {
	filepath := filepath.Join(reportDir, filename)
	return os.ReadFile(filepath)
}
