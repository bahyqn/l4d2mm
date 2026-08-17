package components

import (
	"fmt"
	"l4d2mm/internal"
	"l4d2mm/internal/schema"

	"github.com/go-gui-org/go-gui/gui"
)

func ModsHeader() gui.View {
	return gui.Row(gui.ContainerCfg{
		Sizing:  gui.FillFit,
		Padding: gui.NoPadding,
		Content: []gui.View{
			Input("gnome", internal.GLOBALAPP.ComponentStatus.ModsSearchValue),
			Select("gnome", schema.TemplateSelect{
				ID:       "select",
				Selected: internal.GLOBALAPP.ComponentStatus.SelectedSubLabel,
				Options:  internal.GLOBALAPP.ComponentStatus.SubLabels,
				OnSelectFunc: func(s []string, e *gui.Event, w *gui.Window) {
					internal.GLOBALAPP.ComponentStatus.SelectedSubLabel = s

					fmt.Printf("%v", internal.GLOBALAPP.ComponentStatus.SelectedSubLabel)
				},
			}),
			// VerticalSpacer(),
			ButtonChooseFolder("gnome"),
		},
	})
}
