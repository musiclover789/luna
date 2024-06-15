package emulation

import (
	"github.com/musiclover789/luna/luna_utils"
	"github.com/musiclover789/luna/protocol"
	"github.com/tidwall/gjson"
	"strconv"
)

func SetDeviceMetricsOverride(conn *protocol.Session, width, height, scale int64) gjson.Result {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "Emulation.setDeviceMetricsOverride",
		"params": map[string]interface{}{
			"width":             width,
			"height":            height,
			"deviceScaleFactor": scale,
			"scale":             scale,
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

func SetTouchEmulationEnabled(conn *protocol.Session, maxTouchPoints int) {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "Emulation.setTouchEmulationEnabled",
		"params": map[string]interface{}{
			"enabled":        true,
			"maxTouchPoints": maxTouchPoints,
		},
	}
	conn.WriteMessage(req)
}

func DispatchTouchEvent(conn *protocol.Session, touchType string, x, y int) {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "Input.dispatchTouchEvent",
		"params": map[string]interface{}{
			"type": touchType,
			"touchPoints": []map[string]interface{}{
				{
					"x": x,
					"y": y,
				},
			},
		},
	}
	conn.WriteMessage(req)
}
