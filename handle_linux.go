package wake

func newTimerHandle() (t timerHandle, err error) {
	sd, err := newSdTimerHandle()
	if err != NotImplemented {
		t = &sd
		return
	}

	raw, err := newSysfsTimerHandle()
	if err != nil {
		return
	}
	t = &raw
	return
}
