package my_table

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

// TableRowProvider is a generic interface for converting any struct to table rows.
type TableRowProvider interface {
	ToTableRows() []table.Row
	ToTableHeader() table.Row // Optional: Define headers generically if needed
}

func RenderTable(provider TableRowProvider) {
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)
	tw.AppendHeader(provider.ToTableHeader())
	tw.AppendRows(provider.ToTableRows())
	tw.Render()
}
