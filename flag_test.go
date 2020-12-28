package concurrent_test

import (
	"testing"

	"github.com/alcortesm/concurrent"
)

func TestFlag(t *testing.T) {
	subtests := map[string]func(t *testing.T){
		"new flags start unset":                flagStartUnset,
		"an unset flag can be set":             flagCanBeSet,
		"setting a flag twice has no effect":   flagSettingTwiceIsNop,
		"done is open when flag is unset":      flagDoneIsOpenWhenUnset,
		"done is closed after the flag is set": flagDoneIsClosedAfterSet,
	}

	for name, testFn := range subtests {
		testFn := testFn
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			testFn(t)
		})
	}
}

func flagTODO(t *testing.T) {
}

// Tests that new flags start unset.
func flagStartUnset(t *testing.T) {
	flag := concurrent.NewFlag()

	if flag.IsSet() {
		t.Errorf("wanted an unset flag, but got a set one")
	}
}

// Tests that you can set an unset flag.
func flagCanBeSet(t *testing.T) {
	flag := concurrent.NewFlag()
	flag.Set()

	if !flag.IsSet() {
		t.Errorf("wanted a set flag, but got an unset one")
	}
}

// Tests that setting and already set flag has no effect.
func flagSettingTwiceIsNop(t *testing.T) {
	flag := concurrent.NewFlag()
	flag.Set()
	flag.Set()

	if !flag.IsSet() {
		t.Errorf("wanted a set flag, but got an unset one")
	}
}

// Tests that the done channel is open when the flag is unset.
func flagDoneIsOpenWhenUnset(t *testing.T) {
	flag := concurrent.NewFlag()

	select {
	case _, ok := <-flag.Done():
		if ok {
			t.Fatal("unexpected data in done channel")
		}
		t.Errorf("the done channel is closed")
	default:
		// success
	}
}

// Tests that the done channel is closed after setting the flag.
func flagDoneIsClosedAfterSet(t *testing.T) {
	flag := concurrent.NewFlag()
	flag.Set()

	select {
	case _, ok := <-flag.Done():
		if ok {
			t.Fatal("unexpected data in done channel")
		}
		// success
	default:
		t.Errorf("the done channel is open")
	}
}
