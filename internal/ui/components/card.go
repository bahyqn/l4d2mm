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
	// cardWidth  float32 = 160
	// cardHeight float32 = 170
	// imgHeight  float32 = 90
	cardWidth   float32 = 250
	cardHeight  float32 = 270
	imgHeight   float32 = 140
	addontitle  float32 = 16
	addonautohr float32 = 14
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
		// Radius:       gui.SomeF(12),
		Radius:        gui.RadiusPx(12),
		ClipContents:  true,
		// ClipContents: true,
		ColorBorder: theme.DefaultLightGNOME().BorderColor,
		Shadow: &gui.BoxShadow{
			Color:      gui.RGBA(0, 0, 0, 26),
			OffsetX:    0,
			OffsetY:    3,
			BlurRadius: 8,
		},
		Content: []gui.View{
			gui.Row(gui.ContainerCfg{
				ID:           "block-mod-from",
				Padding:      gui.NewPadding(4, 4, 4, 4),
				Sizing:       gui.FillFit,
				Float:        true,
				FloatOffsetY: 10,
				FloatOffsetX: 215,
				// FloatZIndex:  2,
				Radius: gui.RadiusPx(13),
				Color: func(mod *schema.Mod) gui.Color {
					if mod.IsFromWorkshop {
						return gui.RGBA(0, 120, 212, 230)
					}
					return gui.RGBA(51, 65, 85, 230)
				}(&mod),
				Content: []gui.View{
					// VerticalSpacer(),
					gui.Text(gui.TextCfg{
						ID: "mod-from-letter",
						Text: func(mod *schema.Mod) string {
							if mod.IsFromWorkshop {
								return "W"
							}
							return "L"
						}(&mod),
						TextStyle: gui.TextStyle{
							Size:  16,
							Color: gui.RGBA(255, 255, 255, 255),
							// BgColor: func(mod *schema.Mod) gui.Color {
							// 	if mod.IsFromWorkshop {
							// 		return gui.RGBA(0, 120, 212, 230)
							// 	}
							// 	return gui.RGBA(51, 65, 85, 230)
							// }(&mod),
						},
					}),
				},
			}),
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
						Height:    cardHeight - imgHeight - 48,
						MaxHeight: cardHeight - imgHeight - 48,
						// ColorBorder: theme.DefaultLightGNOME().BorderColor,

						// Padding:    gui.NoPadding,
						Padding: gui.NewPadding(10, 13, 0, 13),
						// Spacing:    gui.NoSpacing,
						// Spacing:    gui.SomeF(2),
						Spacing:    gui.SpacingPx(2),
						SizeBorder: gui.NoBorder,
						Clip:       true,
						// ClipContents: true,
						Content: []gui.View{
							gui.Text(gui.TextCfg{
								ID:     "mod-title" + mod.Id,
								Sizing: gui.FillFit,
								Text:   mod.Name,
								Mode:   gui.TextModeWrap,
								Clip:   true,
								TextStyle: gui.TextStyle{
									Size:        addontitle,
									LineSpacing: 2,
									Color:       gui.RGB(0, 0, 0),
									CellWidth:   120,
								},
							}),
							gui.Text(gui.TextCfg{
								ID: "mod-description" + mod.Id,
								// Text: mod.Remark,
								Sizing: gui.FixedFit,
								Text:   mod.Author,
								Mode:   gui.TextModeSingleLine,
								Clip:   true,
								TextStyle: gui.TextStyle{
									Size:      addonautohr,
									Color:     gui.RGB(0, 0, 0),
									CellWidth: 120,
								},
							}),
						},
					}),
					// buttons
					gui.Row(gui.ContainerCfg{
						VAlign:    gui.VAlignMiddle,
						Width:     cardWidth,
						MaxWidth:  cardWidth,
						Height:    48,
						MaxHeight: 48,
						// Padding:    gui.NoPadding,
						Padding:    gui.NewPadding(0, 10, 10, 10),
						SizeBorder: gui.NoBorder,
						// Spacing:    gui.SomeF(2),
						Spacing: gui.SpacingPx(8),
						// Button()
						// Color: gui.RGBA(0, 0, 0, 30),
						Content: []gui.View{
							SwitchDsiableMod(&mod),
							// ButtonShowModInfo(&mod),
							// ButtonDisableMod(&mod),
							ButtonDeleteMod(&mod),
							// VerticalSpacer(),
							// MoreHoriz(&mod),

							DisplayModProfileButton(&mod),
							CRCScanButton(&mod),

							gui.Toggle(gui.ToggleCfg{
								ID:         "mod-toggle-" + mod.Id,
								Label:      "",
								Size:       gui.SomeF(26),
								TextSelect: "✔",
								// MinWidth: theme.DefaultTheme.ModCardToggleMinWidth,
								Selected: internal.GLOBALAPP.DI.Task.IsSelected(mod.Id),
								OnClick: func(ec gui.EventCtx) {
									selected := internal.GLOBALAPP.DI.Task.IsSelected(mod.Id)

									if selected {
										internal.GLOBALAPP.DI.Task.RemoveTask(mod.Id)
										return
									}

									internal.GLOBALAPP.DI.Task.AddTask(mod.Id)
								},
							}),
						},
					}),
				},
			}),
		},
	})
}
func modColor(mod *schema.Mod) gui.Color {
	if mod.IsFromWorkshop {
		return theme.DefaultTheme.ColorWorkshopMod
	}
	return theme.DefaultTheme.ColorLocalMod
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
