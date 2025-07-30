package parser

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// Each row is a map of column name -> value
type ParsedRow map[string]string

func ParseExcel(path string) ([]ParsedRow, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}

	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	if len(rows) < 1 {
		return nil, fmt.Errorf("no rows found in the sheet")
	}

	headers := rows[0]
	var parsed []ParsedRow

	for _, row := range rows[1:] {
		entry := make(ParsedRow)
		for i, cell := range row {
			if i < len(headers) {
				entry[headers[i]] = cell
			}
		}
		// Fill missing columns with empty strings
		for i := len(row); i < len(headers); i++ {
			entry[headers[i]] = ""
		}
		parsed = append(parsed, entry)
	}

	return parsed, nil
}
