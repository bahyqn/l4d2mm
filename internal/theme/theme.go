package theme

import "github.com/go-gui-org/go-gui/gui"

// GNOMETheme defines standard colors for the GNOME Adwaita light theme.
type GNOMETheme struct {
	WindowBackground gui.Color // Window background color (`#fafafa`)
	ViewBackground   gui.Color // Sidebar / View background color (`#ffffff`)
	BorderColor      gui.Color // General border color (`1px` divider line)
	ButtonDefault    gui.Color // Default button background (fully transparent / seamless)
	ButtonHover      gui.Color // Button hover state (`Hover`)
	ButtonActive     gui.Color // Button active/selected state (`Clicked`)
	TextPrimary      gui.Color // Primary text color
	TextDim          gui.Color // Secondary / label text color (`--dim-label`)
	AccentBlue       gui.Color // GNOME standard accent blue (`#3584e4`)
}

// DefaultLightGNOME returns a standard GNOME light theme RGBA color instance.
func DefaultLightGNOME() GNOMETheme {
	return GNOMETheme{
		WindowBackground: gui.RGBA(250, 250, 250, 255), // rgba(250, 250, 250, 1)
		ViewBackground:   gui.RGBA(255, 255, 255, 255), // rgba(255, 255, 255, 1)
		BorderColor:      gui.RGBA(0, 0, 0, 20),        // rgba(0, 0, 0, 0.08) Approx. 0.08 opacity
		ButtonDefault:    gui.RGBA(0, 0, 0, 0),         // rgba(0,0,0,0) Default transparent (blends into sidebar)
		ButtonHover:      gui.RGBA(0, 0, 0, 13),        // rgba(0,0,0, 0.05) Approx. 5% light gray for hover
		// ButtonActive:     gui.RGBA(0, 0, 0, 38),       // Gray color block for active/selected state
		ButtonActive: gui.RGBA(0, 0, 0, 30),       // rgba(0,0,0, 0.12)  Gray color block for active/selected state
		TextPrimary:  gui.RGBA(0, 0, 0, 255),      // rgba(0,0,0, 1) Pure black
		TextDim:      gui.RGBA(94, 94, 94, 255),   // rgba(94,94,94, 1) `#5e5e5e`
		AccentBlue:   gui.RGBA(53, 132, 228, 255), // rgba(53, 132, 228, 1) `#3584e4`
	}
}
