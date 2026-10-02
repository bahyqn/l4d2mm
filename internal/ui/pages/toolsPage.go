package pages

import (
	"fmt"

	"github.com/go-gui-org/go-gui/gui"
)

var s string

func ToolsView() gui.View {
	return gui.Column(gui.ContainerCfg{
		Content: []gui.View{
			gui.Breadcrumb(gui.BreadcrumbCfg{
				ID: "tools-nav",
				Items: []gui.BreadcrumbItemCfg{
					gui.NewBreadcrumbItem("tool", "Tools", nil),
					gui.NewBreadcrumbItem("docs", "Docs", nil),
					gui.NewBreadcrumbItem("api", "API Reference", nil),
				},
				Selected: s,
				OnSelect: func(s string, ec gui.EventCtx) {
					fmt.Println(s)
				},
			}),
		},
	})
}
