package browser

import (
	"luna/luna_utils"
	"luna/protocol"
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
