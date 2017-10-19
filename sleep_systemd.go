package wake

import (
	"os"
	"strings"
	"sync"

	"github.com/coreos/go-systemd/login1"
	"github.com/coreos/go-systemd/util"
)

var (
	sdSleepLock  *os.File
	sdSleepLockL sync.Mutex
)

func sdPreventSleep() (err error) {
	if !util.IsRunningSystemd() {
		return NotImplemented
	}

	sdSleepLockL.Lock()
	defer sdSleepLockL.Unlock()

	if sdSleepLock != nil {
		return SleepAlreadyPrevented
	}

	// https://www.freedesktop.org/wiki/Software/systemd/inhibit/
	conn, err := login1.New()
	if err != nil {
		return
	}
	sdSleepLock, err = conn.Inhibit(
		"idle:sleep:shutdown",
		strings.Join(os.Args, " "),
		"unknown reason",
		"block",
	)

	return
}

func sdAllowSleep() (err error) {
	if !util.IsRunningSystemd() {
		return NotImplemented
	}

	sdSleepLockL.Lock()
	defer sdSleepLockL.Unlock()

	if sdSleepLock == nil {
		return SleepAlreadyAllowed
	}

	// https://www.freedesktop.org/wiki/Software/systemd/inhibit/
	sdSleepLock.Close()
	sdSleepLock = nil

	return
}
