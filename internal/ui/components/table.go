package components

import (
	"l4d2mm/internal/schema"
	"strconv"

	"github.com/go-gui-org/go-gui/gui"
)

func Table(tableCfg *schema.TableConfig) gui.View {
	return gui.Column(gui.ContainerCfg{
		ID:     "",
		Sizing: gui.FillFill,
		Content: []gui.View{
			// table header -> row
			tableHeader(tableCfg.Header),
			// tables rows column -> row
			HDivider(2),
			tableBody(tableCfg),
		},
	})
}

func tableHeader(header []schema.TableField) gui.View {
	headers := []gui.View{}

	for _, v := range header {
		headers = append(headers, gui.Text(gui.TextCfg{
			ID:       v.Text,
			Text:     v.Text,
			MinWidth: v.MinWidth,
			// TextStyle: gui.TextStyle{
			// 	Color: gui.RGBA(0, 0, 0, 255),
			// 	Align: gui.TextAlignCenter,
			// },
		}))
	}
	return gui.Row(gui.ContainerCfg{
		ID:      "table-header",
		Sizing:  gui.FillFit,
		Content: headers,
	})
}

func tableBody(tcfg *schema.TableConfig) gui.View {
	columns := []gui.View{}

	for idx, v := range tcfg.RowData {
		columns = append(columns, tableRow(v, tcfg, idx))
	}

	return gui.Wrap(gui.ContainerCfg{
		ID:         "table-wrap",
		Sizing:     gui.FillFill,
		Scrollable: true,
		ScrollMode: gui.ScrollVerticalOnly,
		Content:    columns,
	})
}

func tableRow(row []schema.TableField, tcfg *schema.TableConfig, idx int) gui.View {
	rows := []gui.View{}

	strIdx := strconv.Itoa(idx)
	rows = append(rows, Toggle("toggle-"+strIdx, tcfg, idx, 100))

	for _, v := range row {
		rows = append(rows, gui.Text(gui.TextCfg{
			ID:       strIdx + v.Text,
			Text:     v.Text,
			MinWidth: v.MinWidth,
			Mode:     gui.TextModeWrap,
		}))
	}

	return gui.Row(gui.ContainerCfg{
		ID:      "table-row-" + strIdx,
		Sizing:  gui.FillFit,
		Padding: gui.NoPadding,
		VAlign:  gui.VAlignMiddle,
		Content: rows,
	})
}
