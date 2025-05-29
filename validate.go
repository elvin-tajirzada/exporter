package exporter

import (
	"errors"
	"fmt"
	"github.com/elvin-tajirzada/exporter/pkg/types"
)

// validateOptions checks for required fields and validates supported extensions.
func validateOptions(opts Options) error {
	switch opts.Extension {
	case types.XLSX, types.CSV:
		if len(opts.Headers) == 0 {
			return errors.New("headers cannot be empty")
		}

		if len(opts.Entries) == 0 {
			return errors.New("entries cannot be empty")
		}

		for i, row := range opts.Entries {
			if len(row) == 0 {
				return fmt.Errorf("row %d has no columns", i)
			}
		}

		return nil
	default:
		return fmt.Errorf("unexpected extension: %s", opts.Extension)
	}
}
