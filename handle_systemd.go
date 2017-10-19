package wake

import (
	"fmt"
	"time"

	sddbus "github.com/coreos/go-systemd/dbus"
	"github.com/coreos/go-systemd/util"
	"github.com/godbus/dbus"
)

const (
	sdUnitPrefix = "go-wake-894723409238409238403284932748392-"
	sdNoOpUnit   = sdUnitPrefix + "no-op.service"
)

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
	_, err = t.conn.StartTransientUnit(
		fmt.Sprintf("%s-%d.service", sdUnitPrefix, t.id),
		"fail",
		[]sddbus.Property{
			sddbus.Property{
				"OnActiveSec",
				dbus.MakeVariant(wait / time.Second),
			},
			sddbus.Property{
				"OnUnitActiveSec",
				dbus.MakeVariant(period / time.Second),
			},
			sddbus.Property{
				"WakeSystem",
				dbus.MakeVariant(true),
			},
		},
		nil,
	)
	return
}
