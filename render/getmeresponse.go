package render

import (
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/nlm/briq-cli/briq"
)

func init() {
	DefaultRenderer().Register(&briq.GetMeResponse{}, renderGetMeResponse)
}

func renderGetMeResponse(t table.Writer, v any) {
	obj := v.(*briq.GetMeResponse)
	t.AppendHeader(table.Row{"Key", "Value"})
	t.AppendRows([]table.Row{
		{"Id", obj.Id},
		{"Username", obj.Username},
		{"Email", obj.Email},
		{"Full Name", strings.Join([]string{obj.FirstName, obj.LastName}, " ")},
		{"Display Name", obj.DisplayName},
	})
}
