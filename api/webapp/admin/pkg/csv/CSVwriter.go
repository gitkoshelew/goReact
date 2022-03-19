package csv

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
)

//MakeCSV make csv file and write data
func MakeCSV(data interface{}, name string) (string, error) {

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = ','
		return gocsv.NewSafeCSVWriter(writer)
	})

	path := filepath.Join("/rest-api/pkg/csv/" + name)

	file, err := os.Create(path)	

	if err != nil {
		return "", err
	}
	defer file.Close()

	err = gocsv.MarshalFile(data, file)
	if err != nil {
		return "", err
	}
	return path, nil

	
}
