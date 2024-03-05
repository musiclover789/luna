package dom

import (
	"github.com/tidwall/gjson"
	"luna/luna_utils"
	"luna/protocol"
	"strconv"
)

func DOMEnable(conn *protocol.DevToolsConn) gjson.Result {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "DOM.enable",
	}
	ch := make(chan map[string]interface{}) // 创建一个int类型的channel
	conn.SubscribeOneTimeEvent(strconv.Itoa(id), func(param interface{}) {
		ch <- param.(map[string]interface{})
	})
	conn.WriteMessage(req)
	return gjson.ParseBytes(luna_utils.FormatJSONAsBytes(<-ch))
}

func GetDocument(conn *protocol.DevToolsConn, depth int, pierce bool) gjson.Result {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "DOM.getDocument",
		"params": map[string]interface{}{
			"depth":  depth,
			"pierce": pierce,
		},
	}
	ch := make(chan map[string]interface{})
	conn.SubscribeOneTimeEvent(strconv.Itoa(id), func(param interface{}) {
		ch <- param.(map[string]interface{})
	})
	conn.WriteMessage(req)
	return gjson.ParseBytes(luna_utils.FormatJSONAsBytes(<-ch))
}

func QuerySelector(conn *protocol.DevToolsConn, nodeId int64, selector string) gjson.Result {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "DOM.querySelector",
		"params": map[string]interface{}{
			"nodeId":   nodeId,
			"selector": selector,
		},
	}
	ch := make(chan map[string]interface{})
	conn.SubscribeOneTimeEvent(strconv.Itoa(id), func(param interface{}) {
		ch <- param.(map[string]interface{})
	})
	conn.WriteMessage(req)
	return gjson.ParseBytes(luna_utils.FormatJSONAsBytes(<-ch))
}

func SetFileInputFiles(conn *protocol.DevToolsConn, nodeId int64, files []string) {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "DOM.setFileInputFiles",
		"params": map[string]interface{}{
			"files":  files,
			"nodeId": nodeId,
		},
	}
	conn.WriteMessage(req)
}

func GetOuterHTML(conn *protocol.DevToolsConn, nodeID int64) gjson.Result {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "DOM.getOuterHTML",
		"params": map[string]interface{}{
			"nodeId": nodeID,
		},
	}
	ch := make(chan map[string]interface{})
	conn.SubscribeOneTimeEvent(strconv.Itoa(id), func(param interface{}) {
		ch <- param.(map[string]interface{})
	})
	conn.WriteMessage(req)
	return gjson.ParseBytes(luna_utils.FormatJSONAsBytes(<-ch))
}

func GetOuterHTMLWithMap(conn *protocol.DevToolsConn, nodeID int64) map[string]interface{} {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "DOM.getOuterHTML",
		"params": map[string]interface{}{
			"nodeId": nodeID,
		},
	}
	ch := make(chan map[string]interface{})
	conn.SubscribeOneTimeEvent(strconv.Itoa(id), func(param interface{}) {
		ch <- param.(map[string]interface{})
	})
	conn.WriteMessage(req)
	return <-ch
}

func SetOuterHTML(conn *protocol.DevToolsConn, nodeID int64, outerHTML string) gjson.Result {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "DOM.setOuterHTML",
		"params": map[string]interface{}{
			"nodeId":    nodeID,
			"outerHTML": outerHTML,
		},
	}
	ch := make(chan map[string]interface{})
	conn.SubscribeOneTimeEvent(strconv.Itoa(id), func(param interface{}) {
		ch <- param.(map[string]interface{})
	})
	conn.WriteMessage(req)
	return gjson.ParseBytes(luna_utils.FormatJSONAsBytes(<-ch))
}
