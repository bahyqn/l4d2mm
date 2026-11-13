package components

import (
	"strconv"

	"l4d2mm/internal"
	"l4d2mm/internal/schema"
	"l4d2mm/internal/theme"

	"github.com/go-gui-org/go-gui/gui"
)

func TextModsAcount() gui.View {
	amount := len(internal.GLOBALAPP.DI.Task.Selected)

	return gui.Column(gui.ContainerCfg{
		ID:      "column-mods-amount",
		Sizing:  gui.FillFit,
		Padding: gui.NewPadding(8, 8, 8, 30),
		Content: []gui.View{
			gui.Text(gui.TextCfg{
				ID:   "mods-amount",
				Text: strconv.Itoa(amount) + " mods",
			}),
		},
	})
}

func SettingRow(settingRow schema.PreferenceRowConfig) gui.View {
	return gui.Column(gui.ContainerCfg{
		ID:      settingRow.ColumnID,
		Padding: gui.NoPadding,
		Spacing: gui.SpacingPx(5),
		Content: []gui.View{
			gui.Text(gui.TextCfg{
				ID:   settingRow.PrimaryTextID,
				Text: settingRow.PrimaryTextLabel,
				TextStyle: gui.TextStyle{
					Size:  16,
					Color: theme.DefaultTheme.TextPrimary,
				},
			}),
			gui.Text(gui.TextCfg{
				ID:   settingRow.SecondaryTextID,
				Text: settingRow.SecondaryTextLabel,
				TextStyle: gui.TextStyle{
					Size:  12,
					Color: theme.DefaultTheme.TextSecondary,
				},
			}),
		},
	})
}

func TextVPKBoard(id string, enabled int, disabled int) gui.View {
	return gui.Column(gui.ContainerCfg{
		ID:     "",
		Sizing: gui.FillFit,
		Content: []gui.View{
			gui.Text(gui.TextCfg{
				ID:   "title-" + id,
				Text: id,
				TextStyle: gui.TextStyle{
					Size:  20,
					Color: theme.DefaultTheme.TextPrimary,
				},
			}),
			gui.Text(gui.TextCfg{
				ID:   id + "-enabled",
				Text: "Enabled: " + strconv.Itoa(enabled),
			}), gui.Text(gui.TextCfg{
				ID:   id + "-enabled",
				Text: "Disabled: " + strconv.Itoa(disabled),
			}),
		},
	})
}
