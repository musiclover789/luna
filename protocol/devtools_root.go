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
	webSocketDebuggerUrl, err := GetDefaultWebSocketDebuggerUrl(devtoolsRoot.Port)
	if err != nil {
		return err, nil
	}
	return CreteDevToolsConn(*webSocketDebuggerUrl)
}
