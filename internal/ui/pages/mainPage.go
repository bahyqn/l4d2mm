package pages

import (
	"l4d2mm/internal"
	"l4d2mm/internal/theme"

	"github.com/go-gui-org/go-gui/gui"
)

type AsideItem struct {
	name string
	icon string
}

var taside = []AsideItem{
	{
		name: "Mods",
		icon: "extension.svg",
	},
	{
		name: "Workshop",
		icon: "build.svg",
	},
	{
		name: "Download",
		icon: "download.svg",
	},
	{
		name: "Servers",
		icon: "host.svg",
	},
	{
		name: "Settings",
		icon: "settings.svg",
	},
}

func forloopAside(app *internal.App) []gui.View {
	tmp := []gui.View{}

	// tmp = append(tmp, components.HorizontalGap(1, gui.FillFixed))

	for idx, el := range taside {
		tmp = append(tmp, gui.Button(gui.ButtonCfg{
			ID:     el.name,
			HAlign: gui.Some(gui.HAlignLeft),
			VAlign: gui.Some(gui.VAlignMiddle),
			Sizing: gui.FillFit,
			Height: 35,
			// Color:       gui.RGBA(0, 0, 0, 13),
			// Color:       gui.RGBA(53, 132, 228, 255),
			Color: app.ButtonColorChoice(idx, theme.DefaultLightGNOME().ButtonActive, theme.DefaultLightGNOME().ButtonDefault),
			Colors: gui.ColorSet{
				Base:   app.ButtonColorChoice(idx, theme.DefaultLightGNOME().ButtonActive, theme.DefaultLightGNOME().ButtonDefault),
				Border: theme.DefaultLightGNOME().BorderColor,
				Hover:  theme.DefaultLightGNOME().BorderColor,
			},
			SizeBorder: gui.NoBorder,
			Padding:    gui.NewPadding(8, 22, 8, 20),
			Radius:     gui.SomeF(8),
			OnClick: func(ec gui.EventCtx) {
				ec.Event.IsHandled = true
				app.SetAsideIdx(idx)
			},
			Content: []gui.View{
				gui.Svg(gui.SvgCfg{
					Width:    15,
					Height:   15,
					Sizing:   gui.FixedFixed,
					FileName: "assets/icons/" + el.icon,
				}),
				gui.Text(gui.TextCfg{
					Text: el.name,
					TextStyle: gui.TextStyle{
						Size:  10,
						Color: gui.RGB(0, 0, 0),
					},
				}),
			},
		}))
	}
	return tmp
}

func AsideView() gui.View {
	app := internal.GLOBALAPP

	return gui.Column(gui.ContainerCfg{
		ID:       "aside",
		Sizing:   gui.FixedFill,
		Width:    float32(internal.GLOBALAPP.AppConfig.AsideWidth),
		MaxWidth: float32(internal.GLOBALAPP.AppConfig.AsideWidth),
		// Height:   float32(internal.GLOBALAPP.AppConfig.Height),
		// 5,5,5,5
		// Padding:  gui.Some(gui.PaddingSmall),
		Padding: gui.NewPadding(9, 5, 9, 5),
		Radius:  gui.NoRadius,
		Spacing: gui.SomeF(5),
		// Color:       gui.RGBA(255, 255, 255, 255),
		Color:       theme.DefaultLightGNOME().ButtonHover,
		SizeBorder:  gui.NoBorder,
		ColorBorder: theme.DefaultLightGNOME().BorderColor,
		Shadow: &gui.BoxShadow{
			Color:      gui.RGBA(255, 255, 255, 128),
			OffsetX:    1,
			BlurRadius: 10,
		},
		Content: forloopAside(app),
	})
}

func MainView() gui.View {

	var tmp gui.View

	switch internal.GLOBALAPP.ComponentStatus.AsideIdx {
	case 0:
		tmp = ModsView()
	case 1:
		tmp = WorkshopView()
	case 2:
		tmp = DownloadView()
	case 3:
		tmp = ServersView()
	case 4:
		tmp = SettingsView()

	}
	return tmp
}
