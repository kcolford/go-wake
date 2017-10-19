// +build !windows
// +build !linux

package wake

func newTimerHandle() (t timerHandle, err error) {
	err = NotImplemented
	return
}
