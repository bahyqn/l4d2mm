package components

import (
	"l4d2mm/internal"
	"l4d2mm/internal/schema"

	"github.com/go-gui-org/go-gui/gui"
)

func ModsHeader() gui.View {
	components := []gui.View{
		ButtonChooseFolder(),
		Input(internal.GLOBALAPP.ComponentStatus.ModsSearchValue),
		Select(100, schema.TemplateSelect{
			ID:       "select-category",
			Selected: internal.GLOBALAPP.ComponentStatus.SelectedategoryLabel,
			Options:  internal.GLOBALAPP.ComponentStatus.Categories,
			OnSelectFunc: func(s []string, ec gui.EventCtx) {
				internal.GLOBALAPP.ComponentStatus.SelectedategoryLabel = s
				internal.GetSubCategories(s[0])

			},
		}),
	}

	if len(internal.GLOBALAPP.ComponentStatus.SelectedSubLabel) > 0 {
		components = append(components, Select(100, schema.TemplateSelect{
			ID:       "select-sub-categories",
			Selected: internal.GLOBALAPP.ComponentStatus.SelectedSubLabel,
			Options:  internal.GLOBALAPP.ComponentStatus.SubLabels,
			OnSelectFunc: func(s []string, ec gui.EventCtx) {
				internal.GLOBALAPP.ComponentStatus.SelectedSubLabel = s
			},
		}))
	}

	if len(internal.GLOBALAPP.ComponentStatus.ModsBySelectIdx) != 0 {
		components = append(components, Select(70, schema.TemplateSelect{
			ID:        "mod-pagination",
			Selected:  internal.GLOBALAPP.ComponentStatus.PageEnd,
			Options:   internal.GLOBALAPP.DynamicPageSelect(),
			Invisible: len(internal.GLOBALAPP.ComponentStatus.ModsBySelectIdx) == 0,
			OnSelectFunc: func(s []string, ec gui.EventCtx) {
				internal.GLOBALAPP.ComponentStatus.PageEnd = s
				internal.GLOBALAPP.DynamicMods()
			},
		}))
	}

	return gui.Row(gui.ContainerCfg{
		Sizing:  gui.FillFit,
		Padding: gui.NoPadding,
		Content: components,
	})
}
