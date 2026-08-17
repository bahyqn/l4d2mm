package main

import (
	"l4d2mm/internal"
	"l4d2mm/internal/theme"
	"l4d2mm/internal/ui/components"
	"l4d2mm/internal/ui/pages"

	"github.com/go-gui-org/go-gui/gui"
	"github.com/go-gui-org/go-gui/gui/backend"
)

func main() {
	gui.SetTheme(gui.ThemeLight.WithBorders(true))

	w := gui.NewWindow(gui.WindowCfg{
		State:     internal.NewApp(),
		Title:     "l4d2mm",
		Width:     int(internal.GLOBALAPP.AppConfig.Width),
		Height:    int(internal.GLOBALAPP.AppConfig.Height),
		FixedSize: false,
		OnInit: func(w *gui.Window) {

			// vpk := internal.Vpk{}
			// vpk.ReadAllVpk()
			internal.GLOBALAPP.RegisterAllDependencies(w)
			w.UpdateView(mainView)
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
