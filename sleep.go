// +build !windows
// +build !linux

package wake

func preventSleep() (err error) {
	err = NotImplemented
	return
}

func allowSleep() (err error) {
	err = NotImplemented
	return
}
