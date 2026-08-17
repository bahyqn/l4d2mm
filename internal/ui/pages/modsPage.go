package pages

import (
	"l4d2mm/internal"
	"l4d2mm/internal/schema"
	"l4d2mm/internal/ui/components"
	"strconv"

	"github.com/go-gui-org/go-gui/gui"
)

func ModsView1() gui.View {
	return gui.Column(gui.ContainerCfg{
		// Width:     internal.GLOBALAPP.AppConfig.Width - internal.GLOBALAPP.AppConfig.AsideWidth - 1,
		// MaxWidth:  internal.GLOBALAPP.AppConfig.Width - internal.GLOBALAPP.AppConfig.AsideWidth - 1,
		// Height:    internal.GLOBALAPP.AppConfig.Height,
		// MaxHeight: internal.GLOBALAPP.AppConfig.Height,
		Spacing: gui.SomeF(1),

		Sizing: gui.FillFill,
		// Padding: gui.NoPadding,
		Content: []gui.View{
			components.ModsHeader(),
			components.HorizontalGap(7, gui.FillFixed),
			components.HDivider(1),
			components.HorizontalGap(5, gui.FillFixed),
			gui.Wrap(gui.ContainerCfg{
				ID:         "mods",
				Sizing:     gui.FillFill,
				Spacing:    gui.SomeF(13),
				Scrollable: true,
				ScrollMode: gui.ScrollVerticalOnly,
				Content: []gui.View{
					// gui.Button(gui.ButtonCfg{
					// 	Width:  150,
					// 	Height: 100,
					// }),
					gui.Button(gui.ButtonCfg{
						Width:  300,
						Height: 150,
					}),
					gui.Button(gui.ButtonCfg{
						Width:  300,
						Height: 150,
					}),
					gui.Button(gui.ButtonCfg{
						Width:  300,
						Height: 150,
					}),
					gui.Button(gui.ButtonCfg{
						Width:  300,
						Height: 150,
					}),
					gui.Button(gui.ButtonCfg{
						Width:  300,
						Height: 150,
					}),
					gui.Button(gui.ButtonCfg{
						Width:  300,
						Height: 150,
					}),
					gui.Button(gui.ButtonCfg{
						Width:  300,
						Height: 150,
					}),
					gui.Button(gui.ButtonCfg{
						Width:  300,
						Height: 150,
					}),
					gui.Button(gui.ButtonCfg{
						Width:  300,
						Height: 150,
					}),
					gui.Button(gui.ButtonCfg{
						Width:  300,
						Height: 150,
					}),
					gui.Button(gui.ButtonCfg{
						Width:  300,
						Height: 150,
					}),
					gui.Button(gui.ButtonCfg{
						Width:  300,
						Height: 150,
					}),
					gui.Button(gui.ButtonCfg{
						Width:  300,
						Height: 150,
					}),
					gui.Button(gui.ButtonCfg{
						Width:  300,
						Height: 150,
					}),
				},
			}),
		},
	})
}

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
				ID:         "mods",
				Sizing:     gui.FillFill,
				Spacing:    gui.SomeF(16),
				Scrollable: true,
				ScrollMode: gui.ScrollVerticalOnly,
				// Wrap:       true,
				// Overflow:   true,
				Content: renderCards(),
			}),
		},
	})
}

func renderCards() []gui.View {
	fakeData()

	t := []gui.View{}

	if len(internal.GLOBALAPP.DI.Vpk.Mods) > 0 {
		for idx, item := range internal.GLOBALAPP.DI.Vpk.Mods {
			if idx > 70 {
				break
			}
			t = append(t, components.ModCard(item))
		}
	}

	return t
}

func fakeData() {
	for i := 0; i < 100; i++ {
		internal.GLOBALAPP.DI.Vpk.Mods = append(internal.GLOBALAPP.DI.Vpk.Mods, schema.Mod{
			Id: strconv.Itoa(i),
		})
	}
}
