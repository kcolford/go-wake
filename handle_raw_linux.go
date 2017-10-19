package wake

import (
	"fmt"
	"os"
	"time"
)

type rawTimerHandle struct {
	stop chan<- struct{}
	sig  chan struct{}
}

func newRawTimerHandle() (t rawTimerHandle, err error) {
	t.sig = make(chan struct{})
	return
}

func (t *rawTimerHandle) waitfor(stop <-chan struct{}, d time.Duration) (err error) {
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

func (t *rawTimerHandle) Start(wait, period time.Duration) (err error) {
	close(t.stop)
	// use a separate stop so that the goroutine binds to this and
	// doesn't cause a race condition when we modify t
	stop := make(chan struct{})
	t.stop = stop
	go func() {
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
	}()
	return
}

func (t *rawTimerHandle) Wait(timeout time.Duration) (again bool, err error) {
	select {
	case <-t.sig:
	case <-time.After(timeout):
		again = true
	}
	return
}

func (t *rawTimerHandle) Close() {
	close(t.stop)
}

func raw(d time.Duration) (err error) {
	file, err := os.Create("/sys/class/rtc/rtc0/wakealarm")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(file, "+%d", d/time.Second)
	return
}
