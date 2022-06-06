package render

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/nlm/briq-cli/briq"
)

func init() {
	DefaultRenderer().Register(&briq.ListUsersResponse{}, renderListUsersResponse)
}

func renderListUsersResponse(t table.Writer, v any) {
	obj := v.(*briq.ListUsersResponse)
	t.AppendHeader(table.Row{"Id", "Username", "Display Name"})
	for _, user := range obj.Users {
		t.AppendRow(table.Row{
			user.Id,
			user.Username,
			user.DisplayName,
		})
	}
}
