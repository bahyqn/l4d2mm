package components

import (
	"l4d2mm/internal"
	"l4d2mm/internal/schema"

	"github.com/go-gui-org/go-gui/gui"
)

func ModsHeader() gui.View {
	components := []gui.View{
		ButtonChooseFolder(),
		Input(internal.GLOBALAPP.ComponentStatus.PageMods.ModsSearchValue),
		Select(schema.TemplateSelect{
			ID:       "select-category",
			MaxWidth: 60,
			Selected: internal.GLOBALAPP.ComponentStatus.PageMods.SelectedategoryLabel,
			Options:  internal.GLOBALAPP.ComponentStatus.PageMods.Categories,
			OnSelectFunc: func(s []string, ec gui.EventCtx) {
				internal.GLOBALAPP.ComponentStatus.PageMods.SelectedategoryLabel = s
				internal.GetSubCategories(s[0])

			},
		}),
	}

	if len(internal.GLOBALAPP.ComponentStatus.PageMods.SelectedSubLabel) > 0 {
		components = append(components, Select(schema.TemplateSelect{
			ID:       "select-sub-categories",
			MaxWidth: 60,
			Selected: internal.GLOBALAPP.ComponentStatus.PageMods.SelectedSubLabel,
			Options:  internal.GLOBALAPP.ComponentStatus.PageMods.SubLabels,
			OnSelectFunc: func(s []string, ec gui.EventCtx) {
				internal.GLOBALAPP.ComponentStatus.PageMods.SelectedSubLabel = s
			},
		}))
	}

	if len(internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx) != 0 {
		components = append(components, Select(schema.TemplateSelect{
			ID:        "mod-pagination",
			MaxWidth:  60,
			Selected:  internal.GLOBALAPP.ComponentStatus.PageMods.PageEnd,
			Options:   internal.GLOBALAPP.ComponentStatus.PageMods.DynamicPageSelect,
			Invisible: len(internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx) == 0,
			OnSelectFunc: func(s []string, ec gui.EventCtx) {
				internal.GLOBALAPP.ComponentStatus.PageMods.PageEnd = s
				// internal.GLOBALAPP.DynamicMods()
			},
		}))
	}

	components = append(components, Select(schema.TemplateSelect{
		ID:       "source-mode",
		MaxWidth: 80,
		Selected: internal.GLOBALAPP.ComponentStatus.PageMods.SourceMode,
		Options:  internal.AllSourceModes,
		OnSelectFunc: func(s []string, ec gui.EventCtx) {
			internal.GLOBALAPP.ComponentStatus.PageMods.SourceMode = s
			SelectSourceMode()
			// internal.GLOBALAPP.Refresh()
			// fmt.Println(len(internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx))
		},
	}))

	components = append(components, SelectAllButton())

	if len(internal.GLOBALAPP.DI.Task.Selected) > 0 {
		components = append(components, CancelAllButton())
	}

	return gui.Row(gui.ContainerCfg{
		Sizing:  gui.FillFit,
		Padding: gui.NoPadding,
		Content: components,
	})
}
