package protocol

type DevtoolsRoot struct {
	Port int
}

func NewDevtoolsRoot(port int) *DevtoolsRoot {
	return &DevtoolsRoot{
		Port: port,
	}
}

func (devtoolsRoot *DevtoolsRoot) FirstConn() (error, *DevToolsConn) {
	webSocketDebuggerUrl := *GetDefaultWebSocketDebuggerUrl(devtoolsRoot.Port)
	return CreteDevToolsConn(webSocketDebuggerUrl)
}
