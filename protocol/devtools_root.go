package protocol

import (
	"strings"
)

type DevtoolsRoot struct {
	Port int
}

func NewDevtoolsRoot(port int) *DevtoolsRoot {
	return &DevtoolsRoot{
		Port: port,
	}
}

func (devtoolsRoot *DevtoolsRoot) FirstConn() (error, *Session, string) {
	webSocketDebuggerUrl, err := GetDefaultWebSocketDebuggerUrl(devtoolsRoot.Port)
	if err != nil {
		return err, nil, ""
	}
	parts := strings.Split(*webSocketDebuggerUrl, "/")
	id := parts[len(parts)-1]
	e, r := CreteSession(*webSocketDebuggerUrl)
	return e, r, id
}
