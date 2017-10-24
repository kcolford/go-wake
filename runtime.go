package wake

import "log"

func go_(fn func()) {
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				log.Print("panic: %s", err)
			}
		}()
		fn()
	}()
}

func ignore_(err error) {
	if err != nil {
		log.Printf("ignored: %s", err)
	}
}
