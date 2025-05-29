package exporter

// Generator defines methods to create export files in different formats.
type Generator interface {
	SetHeaders(headers []string) error
	Generate(entries [][]string) ([]byte, error)
	Close() error
}
