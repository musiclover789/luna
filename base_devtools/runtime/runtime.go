package runtime

import (
	"fmt"
	"github.com/tidwall/gjson"
	"luna/luna_utils"
	"luna/protocol"
	"strconv"
	"time"
)

func Evaluate(conn *protocol.DevToolsConn, expression string) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Runtime.evaluate",
		"params": map[string]interface{}{
			"expression":                  expression,
			"includeCommandLineAPI":       false,
			"silent":                      true,
			"returnByValue":               false,
			"generatePreview":             true,
			"userGesture":                 true,
			"awaitPromise":                false,
			"throwOnSideEffect":           false,
			"disableBreaks":               false,
			"replMode":                    false,
			"allowUnsafeEvalBlockedByCSP": true,
			"serializationOptions":        map[string]interface{}{"serialization": "auto"},
		},
	}
	conn.WriteMessage(req)
}

func EvaluateWithResultSync(conn *protocol.DevToolsConn, expression string, timeout time.Duration) (error, gjson.Result) {
	id := luna_utils.IdGen.NextID()
	ch := make(chan map[string]interface{})
	eventHandler := func(param interface{}) {
		ch <- param.(map[string]interface{})
	}
	//先订阅后执行
	conn.SubscribePersistentEvent(strconv.Itoa(id), eventHandler)
	EvaluateById(conn, expression, id)
	defer conn.UnsubscribePersistentEvent(strconv.Itoa(id))
	for {
		select {
		case result := <-ch:
			return nil, gjson.ParseBytes(luna_utils.FormatJSONAsBytes(result))
		case <-time.After(timeout * time.Second):
			return fmt.Errorf("EvaluateWithResultSync timeout"), gjson.Result{}
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func EvaluateById(conn *protocol.DevToolsConn, expression string, id int) {
	req := map[string]interface{}{
		"id":     id,
		"method": "Runtime.evaluate",
		"params": map[string]interface{}{
			"expression":                  expression,
			"includeCommandLineAPI":       false,
			"silent":                      true,
			"returnByValue":               false,
			"generatePreview":             true,
			"userGesture":                 true,
			"awaitPromise":                false,
			"throwOnSideEffect":           false,
			"disableBreaks":               false,
			"replMode":                    false,
			"allowUnsafeEvalBlockedByCSP": true,
			//"serializationOptions":        map[string]interface{}{"serialization": "auto"},
		},
	}
	conn.WriteMessage(req)
}

func RuntimeEnable(conn *protocol.DevToolsConn) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Runtime.enable",
		"params": map[string]interface{}{},
	}
	conn.WriteMessage(req)
}

func RuntimeDisable(conn *protocol.DevToolsConn) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Runtime.disable",
		"params": map[string]interface{}{},
	}
	conn.WriteMessage(req)
}
