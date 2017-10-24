package wake

import (
	"fmt"
	"log"
)

func go_(fn func()) {
	go func() {
		defer func() { ignore_(recover_()) }()
		fn()
	}()
}

func recover_() (err error) {
	var ok bool
	pnc := recover()
	if err, ok = pnc.(error); !ok {
		err = fmt.Errorf("panic: %s", pnc)
	}
	return
}

func ignore_(err error) {
	if err != nil {
		log.Printf("ignored: %s", err)
	}
}
