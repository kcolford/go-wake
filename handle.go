// +build !windows

package wake

import "time"

type timerHandle struct{}

func newTimerHandle() (t timerHandle, err error) {
	err = NotImplemented
	return
}

func (t *timerHandle) Start(wait, period time.Duration) (err error) {
	err = NotImplemented
	return
}

func (t *timerHandle) Wait(timeout time.Duration) (again bool, err error) {
	err = NotImplemented
	return
}

func (t *timerHandle) Close() (err error) {
	err = NotImplemented
	return
}
