package render

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/jedib0t/go-pretty/v6/table"
)

var defaultRenderer *Renderer = NewRenderer(os.Stdout)

func DefaultRenderer() *Renderer {
	return defaultRenderer
}

type Renderer struct {
	out       io.Writer
	renderers map[reflect.Type]func(table.Writer, any)
}

func NewRenderer(out io.Writer) *Renderer {
	return &Renderer{
		out:       out,
		renderers: make(map[reflect.Type]func(table.Writer, any)),
	}
}

func (r *Renderer) Register(obj any, renderFunc func(table.Writer, any)) {
	r.renderers[reflect.TypeOf(obj)] = renderFunc
}

func (r Renderer) Render(obj any) error {
	t := table.NewWriter()
	t.SetOutputMirror(r.out)
	objType := reflect.TypeOf(obj)
	if renderFunc, ok := r.renderers[objType]; ok {
		renderFunc(t, obj)
		t.Render()
	} else {
		bytes, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(r.out, string(bytes))
	}
	return nil
}
