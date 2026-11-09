package components

import (
	"os"
	"path/filepath"

	"l4d2mm/internal"
	"l4d2mm/internal/schema"
	"l4d2mm/internal/theme"

	"github.com/go-gui-org/go-gui/gui"
)

const (
	// cardWidth  float32 = 160
	// cardHeight float32 = 170
	// imgHeight  float32 = 90
	cardWidth   float32 = 250
	cardHeight  float32 = 260
	imgHeight   float32 = 140
	addontitle  float32 = 16
	addonautohr float32 = 14
)

func ModCard(mod *schema.Mod) gui.View {
	return gui.Column(gui.ContainerCfg{
		ID:        mod.ID,
		Width:     cardWidth,
		MaxWidth:  cardWidth,
		Height:    cardHeight,
		MaxHeight: cardHeight,
		Padding:   gui.NoPadding,
		Spacing:   gui.NoSpacing,
		Sizing:    gui.FitFit,
		// Radius:       gui.SomeF(12),
		Radius:       gui.RadiusPx(12),
		ClipContents: true,
		ColorBorder:  theme.DefaultLightGNOME().BorderColor,
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
				}(mod),
				Content: []gui.View{
					// VerticalSpacer(),
					gui.Text(gui.TextCfg{
						ID: "mod-from-letter",
						Text: func(mod *schema.Mod) string {
							if mod.IsFromWorkshop {
								return "W"
							}
							return "L"
						}(mod),
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
				ID:    "img-" + mod.ID,
				Src:   img(mod),
				Width: cardWidth,
				// MaxWidth:  cardWidth,
				Height:    imgHeight,
				MaxHeight: imgHeight,
			}),
			HDivider(1),
			gui.Column(gui.ContainerCfg{
				ID:         mod.ID + "-bottom-row",
				Width:      cardWidth,
				Sizing:     gui.FixedFill,
				Padding:    gui.NoPadding,
				Spacing:    gui.NoSpacing,
				SizeBorder: gui.NoBorder,

				Content: []gui.View{
					// title, description, remask
					gui.Column(gui.ContainerCfg{
						ID:        mod.ID + "-text-col",
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
								ID:     "mod-title" + mod.ID,
								Sizing: gui.FillFit,
								Text:   mod.ModName,
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
								ID: "mod-description" + mod.ID,
								// Text: mod.Remark,
								Sizing: gui.FixedFit,
								Text:   "By " + mod.Author,
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
							SwitchDsiableMod(mod),
							// ButtonShowModInfo(&mod),
							// ButtonDisableMod(&mod),
							ButtonDeleteMod(mod),
							// VerticalSpacer(),
							// MoreHoriz(&mod),

							DisplayModProfileButton(mod),
							CRCScanButton(mod),

							gui.Toggle(gui.ToggleCfg{
								ID:         "mod-toggle-" + mod.ID,
								Label:      "",
								Size:       gui.SomeF(26),
								TextSelect: "✔",
								// MinWidth: theme.DefaultTheme.ModCardToggleMinWidth,
								Selected: internal.GLOBALAPP.DI.Task.IsSelected(mod.ID),
								OnClick: func(ec gui.EventCtx) {
									selected := internal.GLOBALAPP.DI.Task.IsSelected(mod.ID)

									if selected {
										internal.GLOBALAPP.DI.Task.RemoveTask(mod.ID)
										return
									}

									internal.GLOBALAPP.DI.Task.AddTask(mod.ID)
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
		return imgStat(filepath.Join(internal.GLOBALAPP.DI.Vpk.WorkshopDir, mod.ID+".jpg"))
	}
	return imgStat(filepath.Join(internal.GLOBALAPP.DI.Vpk.AddonsDir, mod.ID+".jpg"))
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

func SettingsGroup(data *schema.SettingsGroupConfig) gui.View {
	groups := []gui.View{}

	rowAmount := len(data.Rows) - 2

	for i, el := range data.Rows {
		groups = append(groups, gui.Row(gui.ContainerCfg{
			ID:      "card-container",
			Sizing:  gui.FillFit,
			Padding: gui.NoPadding,
			VAlign:  gui.VAlignMiddle,
			Content: []gui.View{
				gui.Column(gui.ContainerCfg{
					ID:      "left-row",
					Sizing:  gui.FitFit,
					Spacing: gui.SpacingPx(5),
					Padding: gui.NewPadding(5, 15, 5, 15),
					Content: []gui.View{
						gui.Text(gui.TextCfg{
							ID:   "left-row-title-",
							Text: el.LeftRowTitle,
							TextStyle: gui.TextStyle{
								Color:  theme.DefaultTheme.TextPrimary,
								Size:   16,
								Family: "500",
							},
						}),
						gui.Text(gui.TextCfg{
							ID:   "left-row-sub-title",
							Text: el.LeftRowSubTitle,
							TextStyle: gui.TextStyle{
								Color: theme.DefaultTheme.TextDim,
								Size:  12,
							},
						}),
					},
				}),
				HorizontalSpacer(),
				el.RightView,
				HorizontalGap(7, gui.FixedFit),
			},
		}))

		if i <= rowAmount {
			groups = append(groups, HDivider(1))
		}
	}

	return gui.Column(gui.ContainerCfg{
		ID:       "settings-row-outter-container",
		Width:    700,
		MaxWidth: 700,
		Padding:  gui.NoPadding,
		Sizing:   gui.FixedFit,
		Content: []gui.View{
			gui.Text(gui.TextCfg{
				ID:   "settings-row-outter-title",
				Text: data.Title,
				TextStyle: gui.TextStyle{
					Color: theme.DefaultTheme.TextPrimary,
					Size:  16,
				}.Bold(),
			}),

			// innner card container
			gui.Column(gui.ContainerCfg{
				ID:          "inner-container-",
				Sizing:      gui.FillFit,
				Padding:     gui.NewPadding(10, 10, 10, 10),
				SizeBorder:  gui.BorderPx(2),
				ColorBorder: theme.DefaultTheme.BorderColor,
				Radius:      gui.RadiusPx(10),
				Content:     groups,
			}),
		},
	})
}
