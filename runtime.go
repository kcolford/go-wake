package wake

import "log"

func go_(fn func()) {
	go func() {
		defer func() {
			ignore_(recover())
		}()
		fn()
	}()
}

func ignore_(err interface{}) {
	if err != nil {
		log.Print(err)
	}
}
