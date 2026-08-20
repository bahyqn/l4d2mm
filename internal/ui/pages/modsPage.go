package pages

import (
	"l4d2mm/internal"
	"l4d2mm/internal/schema"
	"l4d2mm/internal/ui/components"
	"strconv"

	"github.com/go-gui-org/go-gui/gui"
)

func ModsView() gui.View {
	return gui.Column(gui.ContainerCfg{
		// Width:     internal.GLOBALAPP.AppConfig.Width - internal.GLOBALAPP.AppConfig.AsideWidth - 1,
		// MaxWidth:  internal.GLOBALAPP.AppConfig.Width - internal.GLOBALAPP.AppConfig.AsideWidth - 1,
		// Height:    internal.GLOBALAPP.AppConfig.Height,
		// MaxHeight: internal.GLOBALAPP.AppConfig.Height,
		Spacing: gui.SomeF(1),
		Sizing:  gui.FillFill,
		// Padding: gui.NoPadding,

		Content: []gui.View{
			components.ModsHeader(),
			components.HorizontalGap(7, gui.FillFixed),
			components.HDivider(1),
			components.HorizontalGap(5, gui.FillFixed),

			gui.Wrap(gui.ContainerCfg{
				ID:         "mods-wrap",
				Sizing:     gui.FitFill,
				Spacing:    gui.SomeF(16),
				Scrollable: true,
				ScrollMode: gui.ScrollVerticalOnly,
				Overflow:   false,
				// Wrap:       true,
				// Overflow:   true,
				Content: renderCards(),
			}),
		},
	})
}

func renderCards() []gui.View {
	// Uncomment this to mock data
	// fakeData()

	t := []gui.View{}

	if len(internal.GLOBALAPP.ComponentStatus.ModsBySelectIdx) > 0 {
		for _, item := range internal.GLOBALAPP.ComponentStatus.ModsBySelectIdx {
			// Control the number of elements rendered on the page
			// if idx >= 50 {
			// 	break
			// }
			t = append(t, components.ModCard(item))
		}
	}

	return t
}

func fakeData() {
	for i := 0; i < 100; i++ {
		internal.GLOBALAPP.ComponentStatus.ModsBySelectIdx = append(internal.GLOBALAPP.ComponentStatus.ModsBySelectIdx, schema.Mod{
			Id: strconv.Itoa(i),
		})
	}
}
