// Package exporter provides functionality to export tabular data
// in different formats such as XLSX and CSV.
package exporter

import (
	"fmt"
	"github.com/elvin-tajirzada/exporter/pkg/types"
)

type (
	Exporter interface {
		Export() ([]byte, error)
	}

	exporter struct {
		ext     types.Extension
		headers []string
		entries [][]string
		style   Style
	}

	Options struct {
		Extension types.Extension
		Headers   []string
		Entries   [][]string
		Style     Style
	}

	Style struct {
		Horizontal  string
		ColumnWidth int
	}
)

func New(opts Options) (Exporter, error) {
	if err := validateOptions(opts); err != nil {
		return nil, err
	}

	return &exporter{
		ext:     opts.Extension,
		headers: opts.Headers,
		entries: opts.Entries,
		style:   opts.Style,
	}, nil
}

func (e *exporter) Export() ([]byte, error) {
	var (
		generator Generator
		err       error
	)

	// Initialize generator depending on extension
	ext := e.ext
	switch ext {
	case types.XLSX:
		generator, err = NewXLSX(len(e.headers), e.style)
	case types.CSV:
		generator, err = NewCSV()
	}

	if err != nil {
		return nil, fmt.Errorf("unable to initialize %s generator: %v", ext, err)
	}

	if err := generator.SetHeaders(e.headers); err != nil {
		return nil, fmt.Errorf("unable to set %s headers: %v", ext, err)
	}

	body, err := generator.Generate(e.entries)
	if err != nil {
		return nil, fmt.Errorf("unable to generate %s: %v", ext, err)
	}

	if err := generator.Close(); err != nil {
		return nil, fmt.Errorf("unable to close %s: %v", ext, err)
	}

	return body, nil
}
