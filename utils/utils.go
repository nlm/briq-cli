package utils

import (
	"fmt"
)

// PrefixError adds a prefix to an error and wrap it.
// If the error is nil, this function returns nil.
func PrefixError(prefix string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", prefix, err)
	}
	return nil
}
