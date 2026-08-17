package components

import (
	"l4d2mm/internal"
	"l4d2mm/internal/schema"
	"l4d2mm/internal/theme"
	"os"
	"path/filepath"

	"github.com/go-gui-org/go-gui/gui"
)

const (
	cardWidth  float32 = 160
	cardHeight float32 = 150
	imgHeight  float32 = 90
)

func ModCard(mod schema.Mod) gui.View {
	return gui.Column(gui.ContainerCfg{
		ID:        mod.Id,
		Width:     cardWidth,
		MaxWidth:  cardWidth,
		Height:    cardHeight,
		MaxHeight: cardHeight,
		Padding:   gui.NoPadding,
		Spacing:   gui.NoSpacing,
		Sizing:    gui.FitFit,
		// SizeBorder: gui.NoBorder,
		ColorBorder: theme.DefaultLightGNOME().BorderColor,
		// Radius:      gui.SomeF(8),
		Radius: gui.NoRadius,
		Content: []gui.View{
			gui.Image(gui.ImageCfg{
				ID:    "img-" + mod.Id,
				Src:   img(&mod),
				Width: cardWidth,
				// MaxWidth:  cardWidth,
				Height:    imgHeight,
				MaxHeight: imgHeight,
			}),
			HDivider(1),
			gui.Row(gui.ContainerCfg{
				ID:     mod.Id + "-bottom-row",
				Width:  cardWidth,
				Sizing: gui.FixedFit,
				Content: []gui.View{
					// title, description, remask
					gui.Column(gui.ContainerCfg{
						ID:     mod.Id + "-text-col",
						Sizing: gui.FixedFit,
						Width:  cardWidth,
						Content: []gui.View{
							gui.Text(gui.TextCfg{
								ID:   "mod-title" + mod.Id,
								Text: mod.Id,
							}),
							gui.Text(gui.TextCfg{
								ID:   "mod-description" + mod.Id,
								Text: mod.Remark,
							}),
						},
					}),
					// buttons
					gui.Column(gui.ContainerCfg{
						// Button()
					}),
				},
			}),
		},
	})
}

func img(mod *schema.Mod) string {
	if mod.IsFromWorkshop {
		return imgStat(filepath.Join(internal.GLOBALAPP.DI.Vpk.WorkshopDir, mod.Id+".jpg"))
	}
	return imgStat(filepath.Join(internal.GLOBALAPP.DI.Vpk.AddonsDir, mod.Id+".jpg"))
}

func imgStat(imgPath string) string {
	_, err := os.Stat(imgPath)

	if err != nil {
		if os.IsNotExist(err) {
			return "assets/imgs/l4d2.png"
		}
	}
	return imgPath
}
