package main

import (
	"embed"
	"l4d2mm/internal"
	theme "l4d2mm/internal/theme"
	"l4d2mm/internal/ui/components"
	"l4d2mm/internal/ui/pages"

	"github.com/go-gui-org/go-gui/gui"
	"github.com/go-gui-org/go-gui/gui/backend"
)

//go:embed assets/*
var assetFiles embed.FS

func main() {
	app := internal.NewApp(assetFiles)
	gui.SetTheme(gui.ThemeLight.WithBorders(true))

	w := gui.NewWindow(gui.WindowCfg{
		State:     app,
		Title:     "l4d2mm",
		Width:     int(internal.GLOBALAPP.AppConfig.Width),
		Height:    int(internal.GLOBALAPP.AppConfig.Height),
		FixedSize: true,
		OnInit: func(w *gui.Window) {

			// vpk := internal.Vpk{}
			// vpk.ReadAllVpk()
			internal.GLOBALAPP.RegisterAllDependencies(w)
			internal.GLOBALAPP.Prepare()

			w.SetView(mainView)
		},
	})

	backend.Run(w)
}

func mainView(w *gui.Window) gui.View {

	return gui.Row(gui.ContainerCfg{
		Sizing: gui.FillFill,
		// HAlign: gui.HAlignCenter,
		// VAlign: gui.VAlignMiddle,
		Padding:    gui.NoPadding,
		Spacing:    gui.NoSpacing,
		Color:      theme.DefaultLightGNOME().WindowBackground,
		SizeBorder: gui.NoBorder,
		Content: []gui.View{
			pages.AsideView(),
			components.VDivider(1),
			pages.MainView(),
		},
	})
}
