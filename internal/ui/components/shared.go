package components

import (
	"l4d2mm/internal/theme"

	"github.com/go-gui-org/go-gui/gui"
)

// Vertical
func VDividerBywidth(w, h int) gui.View {
	return gui.Column(gui.ContainerCfg{
		// Height:      500,
		// Width:       1,
		Width:       float32(w),
		Height:      float32(h),
		Color:       theme.DefaultLightGNOME().ViewBackground,
		SizeBorder:  gui.SomeF(1),
		ColorBorder: theme.DefaultLightGNOME().BorderColor,
		Padding:     gui.NoPadding,
		Radius:      gui.NoRadius,
	})
}

// horizontal
func HDividerByWidth(w, h int) gui.View {
	return gui.Row(gui.ContainerCfg{
		// Width:       500,
		// Height:      1,
		Width:       float32(w),
		Height:      float32(h),
		Color:       theme.DefaultLightGNOME().ViewBackground,
		SizeBorder:  gui.SomeF(1),
		ColorBorder: theme.DefaultLightGNOME().BorderColor,
		Padding:     gui.NoPadding,
		Radius:      gui.NoRadius,
	})
}

// Vertical
func VDivider(w int) gui.View {
	return gui.Column(gui.ContainerCfg{
		// Height:      500,
		// Width:       1,
		Width:       float32(w),
		Sizing:      gui.FixedFill,
		Color:       theme.DefaultLightGNOME().ViewBackground,
		SizeBorder:  gui.SomeF(1),
		ColorBorder: theme.DefaultLightGNOME().BorderColor,
		Padding:     gui.NoPadding,
		Radius:      gui.NoRadius,
	})
}

func HDivider(h int) gui.View {
	return gui.Row(gui.ContainerCfg{
		// Width:       500,
		// Height:      1,
		Height:      float32(h),
		Sizing:      gui.FillFixed,
		Color:       theme.DefaultLightGNOME().ViewBackground,
		SizeBorder:  gui.SomeF(1),
		ColorBorder: theme.DefaultLightGNOME().BorderColor,
		Padding:     gui.NoPadding,
		Radius:      gui.NoRadius,
	})
}

func VerticalSpacer() gui.View {
	return gui.Row(gui.ContainerCfg{
		Sizing: gui.FillFill,
	})
}

func HorizontalSpacer() gui.View {
	return gui.Column(gui.ContainerCfg{
		Sizing: gui.FillFill,
	})
}

func VerticalGap(w int, size gui.Sizing) gui.View {
	return gui.Row(gui.ContainerCfg{
		Width:  float32(w),
		Height: 1,
		Sizing: size,
	})
}

func HorizontalGap(h int, size gui.Sizing) gui.View {
	return gui.Column(gui.ContainerCfg{
		Width:  1,
		Height: float32(h),
		Sizing: size,
	})
}
