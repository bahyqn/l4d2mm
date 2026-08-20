package components

import (
	"fmt"
	"l4d2mm/internal"
	"l4d2mm/internal/schema"
	"l4d2mm/internal/theme"
	"os"
	"path/filepath"

	"github.com/go-gui-org/go-gui/gui"
)

const (
	cardWidth  float32 = 160
	cardHeight float32 = 170
	imgHeight  float32 = 90
)

func ModCard(mod schema.Mod) gui.View {
	return gui.Column(gui.ContainerCfg{
		ID:           mod.Id,
		Width:        cardWidth,
		MaxWidth:     cardWidth,
		Height:       cardHeight,
		MaxHeight:    cardHeight,
		Padding:      gui.NoPadding,
		Spacing:      gui.NoSpacing,
		Sizing:       gui.FitFit,
		Radius:       gui.SomeF(18),
		ClipContents: true,
		// ClipContents: true,
		ColorBorder: theme.DefaultLightGNOME().BorderColor,
		Shadow: &gui.BoxShadow{
			Color:      gui.RGBA(0, 0, 0, 26),
			OffsetX:    0,
			OffsetY:    3,
			BlurRadius: 8,
		},
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
			gui.Column(gui.ContainerCfg{
				ID:         mod.Id + "-bottom-row",
				Width:      cardWidth,
				Sizing:     gui.FixedFill,
				Padding:    gui.NoPadding,
				Spacing:    gui.NoSpacing,
				SizeBorder: gui.NoBorder,

				Content: []gui.View{
					// title, description, remask
					gui.Column(gui.ContainerCfg{
						ID:        mod.Id + "-text-col",
						Sizing:    gui.FixedFixed,
						Width:     cardWidth,
						Height:    cardHeight - imgHeight - 32,
						MaxHeight: cardHeight - imgHeight - 32,
						// ColorBorder: theme.DefaultLightGNOME().BorderColor,

						// Padding:    gui.NoPadding,
						Padding: gui.NewPadding(5, 10, 5, 10),
						// Spacing:    gui.NoSpacing,
						Spacing:    gui.SomeF(2),
						SizeBorder: gui.NoBorder,
						Content: []gui.View{
							gui.Text(gui.TextCfg{
								ID:   "mod-title" + mod.Id,
								Text: "Urban flight",
								Mode: gui.TextModeSingleLine,
								Clip: true,
								TextStyle: gui.TextStyle{
									Size:  11,
									Color: gui.RGB(0, 0, 0),
								},
							}),
							gui.Text(gui.TextCfg{
								ID:   "mod-id" + mod.Id,
								Text: mod.Id + ".vpk",
								Mode: gui.TextModeSingleLine,
								Clip: true,
								TextStyle: gui.TextStyle{
									Size:  10,
									Color: gui.RGB(0, 0, 0),
								},
							}),
							gui.Text(gui.TextCfg{
								ID: "mod-description" + mod.Id,
								// Text: mod.Remark,
								Text: "Author",
								Mode: gui.TextModeSingleLine,
								Clip: true,
								TextStyle: gui.TextStyle{
									Size:  10,
									Color: gui.RGB(0, 0, 0),
								},
							}),
						},
					}),
					// buttons
					gui.Row(gui.ContainerCfg{
						Width:     cardWidth,
						MaxWidth:  cardWidth,
						Height:    20,
						MaxHeight: 20,
						// Padding:    gui.NoPadding,
						Padding:    gui.NewPadding(0, 10, 0, 10),
						SizeBorder: gui.NoBorder,
						// Button()
						// Color: gui.RGBA(0, 0, 0, 30),
						Content: []gui.View{
							ButtonShowModInfo(&mod),
							ButtonDisableMod(&mod),
							ButtonDeleteMod(&mod),
							gui.Switch(gui.SwitchCfg{
								ID:       "mod-switch-" + mod.Id,
								Selected: mod.IsEnable,
								OnClick: func(ec gui.EventCtx) {
									internal.GLOBALAPP.DI.Vpk.DisableMod(&mod)
									fmt.Println("mod isEnable: ", mod.IsEnable)
								},
							}),
						},
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
