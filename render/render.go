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

// DefaultRenderer returns a pointer to the default Renderer
func DefaultRenderer() *Renderer {
	return defaultRenderer
}

// Renerer is a text renderer in which you can register
// types associated to specific rendering functions.
type Renderer struct {
	out       io.Writer
	renderers map[reflect.Type]func(table.Writer, any)
}

// NewRenderer returns a new Renderer writing to "out".
func NewRenderer(out io.Writer) *Renderer {
	return &Renderer{
		out:       out,
		renderers: make(map[reflect.Type]func(table.Writer, any)),
	}
}

// Register registers a rendering function to the Renderer's registry.
func (r *Renderer) Register(obj any, renderFunc func(table.Writer, any)) {
	r.renderers[reflect.TypeOf(obj)] = renderFunc
}

// Render attempts to find and use a specific renderer for obj's type
// and falls back to a JSON representation if none is available.
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
			return fmt.Errorf("unable to render object: %w", err)
		}
		fmt.Fprintln(r.out, string(bytes))
	}
	return nil
}

func Render(obj any) error {
	return DefaultRenderer().Render(obj)
}
