package wake

var setThreadExecutionState = kernel32.NewProc("SetThreadExecutionState")

const (
	esContinuous       = 0x80000000
	esSystemRequired   = 0x00000001
	esDisplayRequired  = 0x00000002
	esAwayModeRequired = 0x00000040
)

func preventSleep() (err error) {
	_, _, err = setThreadExecutionState.Call(esContinuous | esSystemRequired)
	return
}

func allowSleep() (err error) {
	_, _, err = setThreadExecutionState.Call(esContinuous)
	return
}
