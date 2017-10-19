package wake

import (
	"time"

	sddbus "github.com/coreos/go-systemd/dbus"
	"github.com/coreos/go-systemd/util"
)

const sdUnitPrefix = "go-wake-894723409238409238403284932748392-"

type sdTimerHandle struct {
	conn *sddbus.Conn
	id   int
}

func newSdTimerHandle() (t sdTimerHandle, err error) {
	if !util.IsRunningSystemd() {
		err = NotImplemented
		return
	}
	t.conn, err = sddbus.New()
	return
}

func (t *sdTimerHandle) Close() {
	t.conn.Close()
}

func (t *sdTimerHandle) Start(wait, period time.Duration) (err error) {
	return
}
