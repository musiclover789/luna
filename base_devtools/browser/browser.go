package browser

import (
	"github.com/musiclover789/luna/luna_utils"
	"github.com/musiclover789/luna/protocol"
)

func CloseBrowser(conn *protocol.DevToolsConn) {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "Browser.close",
		"params": map[string]interface{}{},
	}
	conn.WriteMessage(req)
}

func SetWindowBounds(conn *protocol.DevToolsConn, left, top, width, height int) {
	id := luna_utils.IdGen.NextID()
	bounds := map[string]interface{}{
		"left":   left,
		"top":    top,
		"width":  width,
		"height": height,
	}
	req := map[string]interface{}{
		"id":     id,
		"method": "Browser.setWindowBounds",
		"params": map[string]interface{}{
			"windowId": 1,
			"bounds":   bounds,
		},
	}
	conn.WriteMessage(req)
}
