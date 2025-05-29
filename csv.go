package exporter

import (
	csvpkg "encoding/csv"
	"fmt"
	"strings"
)

type csv struct {
	write   *csvpkg.Writer
	builder *strings.Builder
}

func NewCSV() (Generator, error) {
	var b strings.Builder
	return &csv{
		write:   csvpkg.NewWriter(&b),
		builder: &b,
	}, nil
}

func (c *csv) SetHeaders(headers []string) error {
	if err := c.write.Write(headers); err != nil {
		return fmt.Errorf("failed to set headers: %v", err)
	}

	return nil
}

func (c *csv) Generate(entries [][]string) ([]byte, error) {
	if err := c.write.WriteAll(entries); err != nil {
		return nil, fmt.Errorf("failed to write all row: %v", err)
	}

	c.write.Flush()
	if err := c.write.Error(); err != nil {
		return nil, fmt.Errorf("failed to flush: %v", err)
	}

	return []byte(c.builder.String()), nil
}

func (c *csv) Close() error {
	c.builder.Reset()
	return nil
}
