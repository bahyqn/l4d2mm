package components

import (
	"l4d2mm/internal"
	"l4d2mm/internal/schema"

	"github.com/go-gui-org/go-gui/gui"
)

func ModsHeader() gui.View {
	return gui.Row(gui.ContainerCfg{
		Sizing:  gui.FillFit,
		Padding: gui.NoPadding,
		Content: []gui.View{
			ButtonChooseFolder(),
			Input(internal.GLOBALAPP.ComponentStatus.ModsSearchValue),
			Select(100, schema.TemplateSelect{
				ID:       "select",
				Selected: internal.GLOBALAPP.ComponentStatus.SelectedSubLabel,
				Options:  internal.GLOBALAPP.ComponentStatus.SubLabels,
				OnSelectFunc: func(s []string, e *gui.Event, w *gui.Window) {
					internal.GLOBALAPP.ComponentStatus.SelectedSubLabel = s
					// fmt.Printf("%v", internal.GLOBALAPP.ComponentStatus.SelectedSubLabel)
				},
			}),
			// VerticalSpacer(),
			Select(70, schema.TemplateSelect{
				ID:       "mod-pagination",
				Selected: internal.GLOBALAPP.ComponentStatus.PageEnd,
				Options:  internal.GLOBALAPP.DynamicPageSelect(),
				OnSelectFunc: func(s []string, e *gui.Event, w *gui.Window) {
					internal.GLOBALAPP.ComponentStatus.PageEnd = s
					internal.GLOBALAPP.DynamicMods()
				},
			}),
		},
	})
}
