package wake

import (
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	t.Parallel()
	Sleep(1 * time.Second)
}

func TestAfter(t *testing.T) {
	t.Parallel()
	<-After(1 * time.Second)
}

func TestPreventSleep(t *testing.T) {
	err := PreventSleep()
	if err != nil {
		if err == NotImplemented {
			t.Skip(err)
		}
		t.Fatal(err)
	}
	err = AllowSleep()
	if err != nil {
		t.Error(err)
	}
}

func TestAllowSleep(t *testing.T) {
	err := PreventSleep()
	if err != nil {
		t.Skip(err)
	}
	err = AllowSleep()
	if err != nil {
		t.Error(err)
	}
}
