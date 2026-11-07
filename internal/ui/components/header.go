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
		// categories
		Select(schema.TemplateSelect{
			ID:       "select-category",
			MaxWidth: 140,
			Selected: pageMods.SelectedategoryLabel,
			Options:  pageMods.Categories,
			OnSelectFunc: func(s []string, ec gui.EventCtx) {
				pageMods.SelectedategoryLabel = s
				internal.GetSubCategories(s[0])

				if s[0] == "All" {
					clear(pageMods.ModsBySelectIdx)
					pageMods.ModsBySelectIdx = pageMods.ModsBySelectIdx[:0]

					pageMods.SelectedSubLabel[0] = ""

					mods := internal.GLOBALAPP.DI.Vpk.Mods
					for i := range internal.GLOBALAPP.DI.Vpk.Mods {
						pageMods.ModsBySelectIdx = append(pageMods.ModsBySelectIdx, &mods[i])
					}

					internal.GLOBALAPP.DynamicPageSelect()
					return
				}

				if s[0] == "Collections" {
					mslice := []string{}

					for _, item := range internal.GLOBALAPP.DI.Vpk.Collections {
						if item.Name == pageMods.SelectedSubLabel[0] {
							mslice = append(mslice, item.Mods...)
							break
						}
					}

					if len(mslice) > 0 {
						clear(pageMods.ModsBySelectIdx)
						pageMods.ModsBySelectIdx = pageMods.ModsBySelectIdx[:0]

						for _, id := range mslice {
							for _, mod := range internal.GLOBALAPP.DI.Vpk.Mods {
								if id == mod.ID {
									pageMods.ModsBySelectIdx = append(pageMods.ModsBySelectIdx, &mod)
									continue
								}
							}
						}
						internal.GLOBALAPP.DynamicPageSelect()
						fmt.Println(len(pageMods.ModsBySelectIdx))
					}
					return
				}
				internal.QueryVPKs()
			},
		}),
	}

	// sublabel
	if len(pageMods.SelectedSubLabel) == 1 && pageMods.SelectedSubLabel[0] != "" {
		components = append(components, Select(schema.TemplateSelect{
			ID:       "select-sub-categories",
			MaxWidth: 160,
			Selected: pageMods.SelectedSubLabel,
			Options:  pageMods.SubLabels,
			OnSelectFunc: func(s []string, ec gui.EventCtx) {
				pageMods.SelectedSubLabel[0] = s[0]

				// fmt.Printf("category: %s, sublabel: %s \n\n", pageMods.SelectedategoryLabel[0], pageMods.SelectedSubLabel[0])

				clear(pageMods.ModsBySelectIdx)
				pageMods.ModsBySelectIdx = pageMods.ModsBySelectIdx[:0]
				for _, item := range internal.GLOBALAPP.DI.Vpk.Mods {
					if item.Category == pageMods.SelectedategoryLabel[0] && item.SubType == pageMods.SelectedSubLabel[0] {
						pageMods.ModsBySelectIdx = append(pageMods.ModsBySelectIdx, &item)
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
