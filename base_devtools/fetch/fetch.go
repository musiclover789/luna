package fetch

import (
	"luna/luna_utils"
	"luna/protocol"
)

func GetResponseBody(conn *protocol.DevToolsConn, requestId string) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Fetch.getResponseBody",
		"params": map[string]interface{}{
			"requestId": requestId,
		},
	}
	conn.WriteMessage(req)
}

