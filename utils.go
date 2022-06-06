package main

import (
	"fmt"

	"github.com/nlm/briq-cli/render"
)

func Render(obj any) error {
	return render.DefaultRenderer().Render(obj)
}

func PrefixError(prefix string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", prefix, err)
	}
	return nil
}
