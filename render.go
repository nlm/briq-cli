package main

import (
	"github.com/nlm/briq-cli/render"
)

func Render(obj any) error {
	return render.DefaultRenderer().Render(obj)
}
