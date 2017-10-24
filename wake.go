package wake

import (
	"errors"
	"runtime"
	"time"

	builtins "github.com/kcolford/go-builtins"
)

var NotImplemented = errors.New("not implemented")

type Timer struct {
	*time.Timer
	h timerHandle
}

func newTimer(d time.Duration, ti *time.Timer) (t *Timer) {
	t = &Timer{ti, nil}
	var err error
	t.h, err = newTimerHandle()
	builtins.Ignore(err)
	if err != nil {
		return
	}
	runtime.SetFinalizer(t.h, func(h timerHandle) {
		h.Close()
	})
	err = t.h.Start(d, 0)
	builtins.Ignore(err)
	if err != nil {
		return
	}
	builtins.Go(func() { builtins.Ignore(keepAlive(t.h)) })
	return
}

func (t *Timer) Reset(d time.Duration) (active bool) {
	active = t.Timer.Reset(d)
	if t.h != nil {
		err := t.h.Start(d, 0)
		builtins.Ignore(err)
		if err != nil {
			return
		}

		builtins.Go(func() { keepAlive(t.h) })
	}
	return
}

func After(d time.Duration) <-chan time.Time {
	return NewTimer(d).C
}

func Sleep(d time.Duration) {
	<-After(d)
}

func AfterFunc(d time.Duration, f func()) *Timer {
	return newTimer(d, time.AfterFunc(d, f))
}

func NewTimer(d time.Duration) *Timer {
	return newTimer(d, time.NewTimer(d))
}

func (t *Timer) Stop() (active bool) {
	active = t.Timer.Stop()
	if t.h != nil {
		builtins.Ignore(t.h.Start(0, 0))
	}
	return
}

type Ticker struct {
	*time.Ticker
	h timerHandle
}

func NewTicker(d time.Duration) (t *Ticker) {
	t = &Ticker{time.NewTicker(d), nil}
	var err error
	t.h, err = newTimerHandle()
	builtins.Ignore(err)
	if err != nil {
		return
	}
	runtime.SetFinalizer(t.h, func(h timerHandle) {
		h.Close()
	})
	err = t.h.Start(d, d)
	builtins.Ignore(err)
	if err != nil {
		return
	}
	builtins.Go(func() { builtins.Ignore(keepAlive(t.h)) })
	return
}

func keepAlive(h timerHandle) error {
	defer runtime.KeepAlive(h)
	return waitAll(h)
}

func waitAll(h timerHandle) (err error) {
	for again, err := h.Wait(1 * time.Second); again && err != nil; {
	}
	return
}

func Tick(d time.Duration) <-chan time.Time {
	return NewTicker(d).C
}

func (t *Ticker) Stop() {
	t.Ticker.Stop()
	if t.h != nil {
		builtins.Ignore(t.h.Start(0, 0))
	}
}

var SleepAlreadyPrevented = errors.New("sleep mode has already been prevented")

func PreventSleep() error {
	return preventSleep()
}

var SleepAlreadyAllowed = errors.New("sleep mode as already been allowed")

func AllowSleep() error {
	return allowSleep()
}
