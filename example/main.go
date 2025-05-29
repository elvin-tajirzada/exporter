package main

import (
	"fmt"
	"github.com/elvin-tajirzada/exporter"
	"github.com/elvin-tajirzada/exporter/pkg/types"
	"os"
)

func main() {
	opts := exporter.Options{
		Extension: types.CSV,
		Headers:   []string{"Name", "Age", "Country"},
		Entries: [][]string{
			{"Alice", "30", "USA"},
			{"Bob", "25", "UK"},
		},
		//Style: exporter.Style{
		//	Horizontal:  "center",
		//	ColumnWidth: 20,
		//},
	}

	exp, err := exporter.New(opts)
	if err != nil {
		panic(err)
	}

	data, err := exp.Export()
	if err != nil {
		panic(err)
	}

	// Use 'data' bytes to save to file or send in response
	fmt.Printf("Exported file size: %d bytes\n", len(data))
}

// saveToFile saves the exported data to a file.
// Use this if you want to visualize the exported output (e.g., open in Excel or a text editor).
func saveToFile(data []byte, filename string) error {
	return os.WriteFile(filename, data, 0644)
}
