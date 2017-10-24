package wake

var setThreadExecutionState = kernel32.NewProc("SetThreadExecutionState")

const (
	esContinuous       = 0x80000000
	esSystemRequired   = 0x00000001
	esDisplayRequired  = 0x00000002
	esAwayModeRequired = 0x00000040
)

func preventSleep() (err error) {
	// https://msdn.microsoft.com/en-us/library/windows/desktop/aa373208(v=vs.85).aspx
	last, _, err = setThreadExecutionState.Call(esContinuous | esSystemRequired)
	if last != 0 {
		err = nil
	}
	return
}

func allowSleep() (err error) {
	// https://msdn.microsoft.com/en-us/library/windows/desktop/aa373208(v=vs.85).aspx
	last, _, err = setThreadExecutionState.Call(esContinuous)
	if last != 0 {
		err = nil
	}
	return
}
