package wake

import "time"

type timerHandle interface {
	Start(wait, period time.Duration) (err error)
	Wait(timeout time.Duration) (again bool, err error)
	Close()
}
