package wake

import (
	"time"
	"unsafe"
)

var (
	createWaitableTimer = kernel32.NewProc("CreateWaitableTimer")
	setWaitableTimer    = kernel32.NewProc("SetWaitableTimer")
	closeHandle         = kernel32.NewProc("CloseHandle")
	waitForSingleObject = kernel32.NewProc("WaitForSingleObject")
)

const (
	waitAbandoned = 0x00000080
	waitObject0   = 0x00000000
	waitTimeout   = 0x00000102
	waitFailed    = 0xFFFFFFFF
)

type winTimerHandle struct {
	hdl uintptr
}

func newWinTimerHandle() (h *winTimerHandle, err error) {
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms682492(v=vs.85).aspx
	hdl, _, err = createWaitableTimer.Call(0, 0, 0)
	if h.hdl != 0 {
		err = nil
	}
	if err != nil {
		return
	}
	h = &winTimerHandle{hdl: hdl}
	return
}

func (h *winTimerHandle) Close() {
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms724211(v=vs.85).aspx
	ok, _, err := closeHandle.Call(h.hdl)
	if ok != 0 {
		err = nil
	}
	ignore_(err)
}

func (h *winTimerHandle) Start(wait, period time.Duration) (err error) {
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms686289(v=vs.85).aspx
	duetime := -int64(wait / time.Nanosecond / 100)
	ok, _, err = setWaitableTimer.Call(h.hdl,
		uintptr(unsafe.Pointer(&duetime)),
		uintptr(period/time.Millisecond), 0, 0, 1)
	if ok != 0 {
		err = nil
	}
	return
}

func (h *winTimerHandle) Wait(timeout time.Duration) (again bool, err error) {
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms687032(v=vs.85).aspx
	res, _, err := waitForSingleObject.Call(h.hdl,
		uintptr(timeout/time.Millisecond))
	if res != waitFailed {
		err = nil
	}
	if err != nil {
		return
	}
	if res == waitTimeout {
		again = true
	}
	return
}

func newTimerHandle() (h timerHandle, err error) {
	return newWinTimerHandle()
}
