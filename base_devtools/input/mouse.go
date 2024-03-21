package input

import (
	"github.com/musiclover789/luna/luna_utils"
	"github.com/musiclover789/luna/protocol"
)

const (
	MouseButtonNone    mouseButton = "none"
	MouseButtonLeft    mouseButton = "left"
	MouseButtonRight   mouseButton = "right"
	MouseButtonMiddle  mouseButton = "middle"
	MouseButtonBack    mouseButton = "back"
	MouseButtonForward mouseButton = "forward"
)

type mouseButton string

type PointerType string

type mouseEventParams struct {
	Type               string      `json:"type"`
	X                  float64     `json:"x"`
	Y                  float64     `json:"y"`
	Modifiers          int         `json:"modifiers,omitempty"`
	Timestamp          int64       `json:"timestamp,omitempty"`
	Button             mouseButton `json:"button,omitempty"`
	Buttons            int         `json:"buttons,omitempty"`
	ClickCount         int         `json:"clickCount,omitempty"`
	DurationMillis     int         // 修改了此处
	Force              float64     `json:"force,omitempty"`
	TangentialPressure float64     `json:"tangentialPressure,omitempty"`
	TiltX              int         `json:"tiltX,omitempty"`
	TiltY              int         `json:"tiltY,omitempty"`
	Twist              int         `json:"twist,omitempty"`
	DeltaX             float64     `json:"deltaX,omitempty"`
	DeltaY             float64     `json:"deltaY,omitempty"`
	PointerType        PointerType `json:"pointerType,omitempty"`
}

func dispatchMouseEvent(conn *protocol.DevToolsConn, params mouseEventParams) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Input.dispatchMouseEvent",
		"params": params,
	}
	conn.WriteMessage(req)
}

func mouseWheel(conn *protocol.DevToolsConn, x, y float64, deltaX, deltaY int) {
	params := map[string]interface{}{
		"type":   "mouseWheel",
		"x":      x,
		"y":      y,
		"deltaX": deltaX,
		"deltaY": deltaY,
	}
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "Input.dispatchMouseEvent",
		"params": params,
	}
	//fmt.Println("我们测试一下:",id, x, y, deltaX, deltaY)
	conn.WriteMessage(req)
}
