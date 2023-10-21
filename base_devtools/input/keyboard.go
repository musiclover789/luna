package input

import (
	"github.com/musiclover789/luna/luna_utils"
	"github.com/musiclover789/luna/protocol"
)

type keyEventParams struct {
	Type           string   `json:"type"`
	Modifiers      int      `json:"modifiers,omitempty"`
	Timestamp      int64    `json:"timestamp,omitempty"`
	Text           string   `json:"text,omitempty"`
	UnmodifiedText string   `json:"unmodifiedText,omitempty"`
	KeyIdentifier  string   `json:"keyIdentifier,omitempty"`
	Code           string   `json:"code,omitempty"`
	Key            string   `json:"key,omitempty"`
	WindowsKeyCode int      `json:"windowsVirtualKeyCode,omitempty"`
	NativeKeyCode  int      `json:"nativeVirtualKeyCode,omitempty"`
	AutoRepeat     bool     `json:"autoRepeat,omitempty"`
	IsKeypad       bool     `json:"isKeypad,omitempty"`
	IsSystemKey    bool     `json:"isSystemKey,omitempty"`
	Location       int      `json:"location,omitempty"`
	Commands       []string `json:"commands,omitempty"`
}

func dispatchKeyEvent(conn *protocol.DevToolsConn, params keyEventParams) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Input.dispatchKeyEvent",
		"params": params,
	}
	conn.WriteMessage(req)
}
