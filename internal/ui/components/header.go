package components

import (
	"fmt"
	"l4d2mm/internal"
	"l4d2mm/internal/schema"

	"github.com/go-gui-org/go-gui/gui"
)

func ModsHeader() gui.View {
	pageMods := &internal.GLOBALAPP.ComponentStatus.PageMods

	components := []gui.View{
		ButtonChooseFolder(),
		Input(pageMods.ModsSearchValue),
		Select(schema.TemplateSelect{
			ID:       "select-category",
			MaxWidth: 140,
			Selected: pageMods.SelectedategoryLabel,
			Options:  pageMods.Categories,
			OnSelectFunc: func(s []string, ec gui.EventCtx) {
				pageMods.SelectedategoryLabel = s
				internal.GetSubCategories(s[0])

				if s[0] == "All" {
					pageMods.SelectedSubLabel[0] = ""
					clear(pageMods.ModsBySelectIdx)
					pageMods.ModsBySelectIdx = pageMods.ModsBySelectIdx[:0]
					pageMods.ModsBySelectIdx = append(pageMods.ModsBySelectIdx, internal.GLOBALAPP.DI.Vpk.Mods...)
					internal.GLOBALAPP.DynamicPageSelect()
					return
				}

				if subLabel, ok := internal.ModLables[s[0]]; ok {
					clear(pageMods.ModsBySelectIdx)
					pageMods.ModsBySelectIdx = pageMods.ModsBySelectIdx[:0]
					if len(subLabel) == 0 {
						for _, item := range internal.GLOBALAPP.DI.Vpk.Mods {
							if item.Category == s[0] {
								pageMods.ModsBySelectIdx = append(pageMods.ModsBySelectIdx, item)
							}
						}
					} else {
						pageMods.SubLabels[0] = subLabel[0]

						// clear(pageMods.ModsBySelectIdx)
						// pageMods.ModsBySelectIdx = pageMods.ModsBySelectIdx[:0]

						for _, item := range internal.GLOBALAPP.DI.Vpk.Mods {
							if item.Category == s[0] && item.SubType == subLabel[0].Value {
								pageMods.ModsBySelectIdx = append(pageMods.ModsBySelectIdx, item)
							}
						}
					}
				}
				internal.GLOBALAPP.DynamicPageSelect()
			},
		}),
	}

	if len(pageMods.SelectedSubLabel) == 1 && pageMods.SelectedSubLabel[0] != "" {
		components = append(components, Select(schema.TemplateSelect{
			ID:       "select-sub-categories",
			MaxWidth: 160,
			Selected: internal.GLOBALAPP.ComponentStatus.PageMods.SelectedSubLabel,
			Options:  internal.GLOBALAPP.ComponentStatus.PageMods.SubLabels,
			OnSelectFunc: func(s []string, ec gui.EventCtx) {
				pageMods.SelectedSubLabel[0] = s[0]

				fmt.Printf("category: %s, sublabel: %s \n\n", pageMods.SelectedategoryLabel[0], pageMods.SelectedSubLabel[0])

				clear(pageMods.ModsBySelectIdx)
				pageMods.ModsBySelectIdx = pageMods.ModsBySelectIdx[:0]
				for _, item := range internal.GLOBALAPP.DI.Vpk.Mods {
					if item.Category == pageMods.SelectedategoryLabel[0] && item.SubType == pageMods.SelectedSubLabel[0] {
						pageMods.ModsBySelectIdx = append(pageMods.ModsBySelectIdx, item)
					}
				}
				internal.GLOBALAPP.DynamicPageSelect()
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
