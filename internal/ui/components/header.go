package components

import (
	"l4d2mm/internal"
	"l4d2mm/internal/schema"

	"github.com/go-gui-org/go-gui/gui"
)

func ModsHeader() gui.View {
	pageMods := &internal.GLOBALAPP.ComponentStatus.PageMods

	components := []gui.View{
		ButtonChooseFolder(),
		Input(pageMods.ModsSearchValue),
		// categories
		Select(schema.TemplateSelect{
			ID:       "select-category",
			MaxWidth: 140,
			Selected: pageMods.SelectedCategory,
			Options:  pageMods.Categories,
			OnSelectFunc: func(s []string, ec gui.EventCtx) {
				pageMods.SelectedCategory = s
				internal.GetSubCategories(s[0])

				internal.QueryVPKs()
				SelectSourceMode()
			},
		}),
	}

	// sublabel
	if len(pageMods.SelectedSubCategory) == 1 && pageMods.SelectedSubCategory[0] != "" {
		components = append(components, Select(schema.TemplateSelect{
			ID:       "select-sub-categories",
			MaxWidth: 160,
			Selected: pageMods.SelectedSubCategory,
			Options:  pageMods.SubCategories,
			OnSelectFunc: func(s []string, ec gui.EventCtx) {
				pageMods.SelectedSubCategory[0] = s[0]

				internal.QueryVPKs()
				SelectSourceMode()
			},
		}))
	}

	if len(pageMods.ModsBySelectIdx) != 0 {
		components = append(components, Select(schema.TemplateSelect{
			ID:        "mod-pagination",
			MaxWidth:  120,
			Selected:  pageMods.PageEnd,
			Options:   pageMods.DynamicPageSelect,
			Invisible: len(pageMods.ModsBySelectIdx) == 0,
			OnSelectFunc: func(s []string, ec gui.EventCtx) {
				pageMods.PageEnd = s
				// internal.GLOBALAPP.DynamicMods()
			},
		}))
	}

	components = append(components, Select(schema.TemplateSelect{
		ID:       "source-mode",
		MaxWidth: 180,
		Selected: []string{pageMods.SourceMode},
		Options:  internal.AllSourceModes,
		OnSelectFunc: func(s []string, ec gui.EventCtx) {
			pageMods.SourceMode = s[0]
			internal.QueryVPKs()
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
		ID:      "mods-header",
		Sizing:  gui.FillFit,
		Padding: gui.NoPadding,
		Content: components,
	})
}
