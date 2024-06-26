package target

import (
	"fmt"
	"github.com/musiclover789/luna/luna_utils"
	"github.com/musiclover789/luna/protocol"
	"github.com/tidwall/gjson"
	"strconv"
	"time"
)

func CreateTarget(conn *protocol.Session, url string) gjson.Result {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "Target.createTarget",
		"params": map[string]interface{}{
			"url": url,
		},
	}
	ch := make(chan map[string]interface{}) // 创建一个int类型的channel
	conn.SubscribeOneTimeEvent(strconv.Itoa(id), func(param interface{}) {
		ch <- param.(map[string]interface{})
	})
	conn.WriteMessage(req)
	return gjson.ParseBytes(luna_utils.FormatJSONAsBytes(<-ch))
}

func AttachToTargetSync(conn *protocol.Session, targetID string, timeout time.Duration) (err error, sessionId string) {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "Target.attachToTarget",
		"params": map[string]interface{}{
			"targetId": targetID,
		},
	}
	resultChan := make(chan string, 1)
	eventHandler := func(param interface{}) {
		value := gjson.Get(luna_utils.FormatJSONAsString(param.(map[string]interface{})), "result.targetId").String()
		resultChan <- value
	}
	//先订阅后执行
	conn.SubscribePersistentEvent(strconv.Itoa(id), eventHandler)
	conn.WriteMessage(req)
	defer conn.UnsubscribePersistentEvent(strconv.Itoa(id))
	for {
		select {
		case result := <-resultChan:
			return nil, result
		case <-time.After(timeout * time.Second):
			return fmt.Errorf("CalculateScalingFactor timeout"), ""
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func ActivateTarget(conn *protocol.Session, targetID string) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Target.activateTarget",
		"params": map[string]interface{}{
			"targetId": targetID,
		},
	}
	conn.WriteMessage(req)
}

func GetTargets(conn *protocol.Session) gjson.Result {
	id := luna_utils.IdGen.NextID()
	req := map[string]interface{}{
		"id":     id,
		"method": "Target.getTargets",
	}
	conn.WriteMessage(req)
	ch := make(chan map[string]interface{}) // 创建一个int类型的channel
	conn.SubscribeOneTimeEvent(strconv.Itoa(id), func(param interface{}) {
		ch <- param.(map[string]interface{})
	})
	conn.WriteMessage(req)
	return gjson.ParseBytes(luna_utils.FormatJSONAsBytes(<-ch))
}

func CloseTarget(conn *protocol.Session, targetId string) {
	req := map[string]interface{}{
		"id":     luna_utils.IdGen.NextID(),
		"method": "Target.closeTarget",
		"params": map[string]interface{}{
			"targetId": targetId,
		},
	}
	conn.WriteMessage(req)
}
