package wake

func preventSleep() (err error) {
	err = sdPreventSleep()
	if err != NotImplemented {
		return
	}

	return
}

func allowSleep() (err error) {
	err = sdAllowSleep()
	if err != NotImplemented {
		return
	}

	return
}
