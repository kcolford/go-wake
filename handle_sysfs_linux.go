package wake

import (
	"fmt"
	"os"
	"time"
)

type sysfsTimerHandle struct {
	stop chan<- struct{}
	sig  chan struct{}
}

func newSysfsTimerHandle() (t sysfsTimerHandle, err error) {
	t.sig = make(chan struct{}, 1)
	t.stop = make(chan struct{})
	return
}

func (t *sysfsTimerHandle) waitfor(stop <-chan struct{}, d time.Duration) (err error) {
	file, err := os.Create("/sys/class/rtc/rtc0/wakealarm")
	if err != nil {
		return
	}
	defer file.Close()
	_, err = fmt.Fprintf(file, "%d\n", time.Now().Add(d).Unix())
	if err != nil {
		return
	}
	select {
	case <-stop:
	case <-time.After(d):
		select {
		case t.sig <- struct{}{}:
		default:
		}
	}
	return
}

func (t *sysfsTimerHandle) Start(wait, period time.Duration) (err error) {
	close(t.stop)
	// use a separate stop so that the goroutine binds to this and
	// doesn't cause a race condition when we modify t
	stop := make(chan struct{})
	t.stop = stop
	go_(func() {
		t.waitfor(stop, wait)
		select {
		case <-stop:
			return
		default:
		}
		if period == 0 {
			return
		}
		for {
			t.waitfor(stop, period)
			select {
			case <-stop:
				return
			default:
			}
		}
	})
	return
}

func (t *sysfsTimerHandle) Wait(timeout time.Duration) (again bool, err error) {
	select {
	case <-t.sig:
	case <-time.After(timeout):
		again = true
	}
	return
}

func (t *sysfsTimerHandle) Close() {
	close(t.stop)
}
