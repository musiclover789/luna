package emulation

import (
	"github.com/tidwall/gjson"
	"github.com/musiclover789/luna/luna_utils"
	"github.com/musiclover789/luna/protocol"
	"strconv"
)

func SetDeviceMetricsOverride(conn *protocol.DevToolsConn, width, height int64) gjson.Result {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "Emulation.setDeviceMetricsOverride",
		"params": map[string]interface{}{
			"width":             width,
			"height":            height,
			"deviceScaleFactor": 0,
			"mobile":            false,
		},
	}
	ch := make(chan map[string]interface{}) // 创建一个int类型的channel
	conn.SubscribeOneTimeEvent(strconv.Itoa(id), func(param interface{}) {
		ch <- param.(map[string]interface{})
	})
	conn.WriteMessage(req)
	return gjson.ParseBytes(luna_utils.FormatJSONAsBytes(<-ch))
}
