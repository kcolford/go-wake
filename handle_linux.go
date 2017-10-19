package wake

func newTimerHandle() (t timerHandle, err error) {
	raw, err := newRawTimerHandle()
	if err != nil {
		return
	}
	t = &raw
	return
}
