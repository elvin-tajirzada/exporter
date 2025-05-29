# exporter

`exporter` is a Go package for exporting tabular data to common file formats such as XLSX (Excel) and CSV. It provides a simple interface to generate export files with customizable headers, entries, and style options.

---

## Features

- Export data to XLSX or CSV formats.
- Customizable column width and horizontal alignment for XLSX.
- Simple API to set headers and entries.
- Extensible with new formats by implementing the `Generator` interface.

---

## Installation

```bash
go get github.com/elvin-tajirzada/exporter
```

## Usage

You can use the exporter like this:

```go
func main() {
	opts := exporter.Options{
		Extension: types.CSV,
		Headers:   []string{"Name", "Age", "Country"},
		Entries: [][]string{
			{"Alice", "30", "USA"},
			{"Bob", "25", "UK"},
		},
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


