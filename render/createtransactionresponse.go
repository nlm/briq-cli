package render

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/nlm/briq-cli/briq"
)

func init() {
	DefaultRenderer().Register(&briq.CreateTransactionResponse{}, renderCreateTransactionResponse)
}

func renderCreateTransactionResponse(t table.Writer, v any) {
	obj := v.(*briq.CreateTransactionResponse)
	t.AppendHeader(table.Row{"From", "Amount", "To"})
	t.AppendRow(table.Row{obj.From, obj.Amount, obj.To})
}
