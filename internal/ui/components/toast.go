package components

import (
	"sync"
	"time"

	"l4d2mm/internal"

	"github.com/go-gui-org/go-gui/gui"
)

// type ToastItem struct {
// 	IconName  string
// 	IconColor gui.Color
// 	Hint      string
// 	HintColor gui.Color
// }

// var toastConfig = []ToastItem{
// 	{
// 		IconName:  "check_cicle.svg",
// 		IconColor: gui.RGBA(40, 167, 69, 255),
// 		Hint:      "",
// 		HintColor: gui.RGBA(40, 167, 69, 255),
// 	},
// 	{
// 		IconName:  "cancel.svg",
// 		IconColor: gui.RGBA(220, 53, 69, 255),
// 		Hint:      "",
// 		HintColor: gui.RGBA(220, 53, 69, 255),
// 	},
// }

// MaxAlive 2s
var MaxAlive = 2

type ToastType string

const (
	ToastSuccess ToastType = "Success"
	ToastFailed  ToastType = "Failed"
)

type ToastItem struct {
	ID    int64
	Type  ToastType
	Msg   string
	Timer *time.Timer
}

type Toast struct {
	mu   sync.Mutex
	Msgs []ToastItem
}

func NewToast() *Toast {
	return &Toast{}
}

func (t *Toast) Add(msg string, flag ToastType) {
	t.mu.Lock()
	id := time.Now().UnixNano()

	item := ToastItem{
		ID:   id,
		Msg:  msg,
		Type: flag,
	}

	item.Timer = time.AfterFunc(time.Duration(MaxAlive), func() {
		t.Remove(id)
	})

	t.Msgs = append(t.Msgs, item)
	t.mu.Unlock()
}

func (t *Toast) Remove(id int64) {
	t.mu.Lock()
	defer t.mu.Unlock()

	newItems := make([]ToastItem, 0, len(t.Msgs))

	for _, item := range t.Msgs {
		if item.ID != id {
			newItems = append(newItems, item)
		} else {
			if item.Timer != nil {
				item.Timer.Stop()
			}
		}
	}

	t.Msgs = newItems

	// if len(t.Msgs) == 0 {

	// }
}

func (t *Toast) ToastSuccessMsg(hint string) {
	gui.Row(gui.ContainerCfg{
		ID:       "toast-success-msg",
		Width:    200,
		MaxWidth: 200,
		Content: []gui.View{
			gui.Svg(gui.SvgCfg{
				ID:       "svg-success" + hint,
				FileName: internal.SetupIcon("check_cicle.svg"),
				Width:    10,
				Height:   10,
				Color:    gui.RGBA(40, 167, 69, 255),
			}),
			gui.Text(gui.TextCfg{
				ID:   "toast-hint",
				Mode: gui.TextModeWrap,
				TextStyle: gui.TextStyle{
					Color: gui.RGBA(40, 167, 69, 255),
					Size:  10,
				},
			}),
		},
	})
}

func (t *Toast) ToastFailedMsg(hint string) {
	gui.Row(gui.ContainerCfg{
		ID:       "toast-failed-msg",
		Width:    200,
		MaxWidth: 200,
		Content: []gui.View{
			gui.Svg(gui.SvgCfg{
				ID:       "svg-success" + hint,
				FileName: internal.SetupIcon("check_cicle.svg"),
				Width:    10,
				Height:   10,
				Color:    gui.RGBA(40, 167, 69, 255),
			}),
			gui.Text(gui.TextCfg{
				ID:   "toast-hint",
				Mode: gui.TextModeWrap,
				TextStyle: gui.TextStyle{
					Color: gui.RGBA(40, 167, 69, 255),
					Size:  10,
				},
			}),

			gui.Svg(gui.SvgCfg{}),
		},
	})
}
