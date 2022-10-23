package keyboard

import (
	"gokey/keyboard/keymap"
	"machine"
	"machine/usb/hid/keyboard"
)

type Coordinates struct {
	Row, Col int
}

type Keyboard interface {
	PressedKeys(map[Coordinates]bool)
}

func New(ledPin machine.Pin) Keyboard {
	hidKeyboard := keyboard.New()

	return &kb{
		activeCodes: map[keyboard.Keycode]bool{},
		hk:          hidKeyboard,
		led:         ledPin,
	}
}

type kbi interface {
	Down(keyboard.Keycode) error
	Up(keyboard.Keycode) error
}

type kb struct {
	activeCodes map[keyboard.Keycode]bool
	hk          kbi
	led         machine.Pin
}

func (k *kb) PressedKeys(cs map[Coordinates]bool) {
	l := 1

	if cs[Coordinates{Row: 4, Col: 5}] {
		l = 2
		k.led.High()
	} else {
		k.led.Low()
	}

	desiredKeycodes := map[keyboard.Keycode]bool{}
	for c := range cs {
		code := keymap.GetCode(l, c.Row, c.Col)
		desiredKeycodes[code] = true
	}

	for activeCode := range k.activeCodes {
		if !desiredKeycodes[activeCode] {
			k.hk.Up(activeCode)
		}

	}

	for code := range desiredKeycodes {
		if !k.activeCodes[code] {
			k.hk.Down(code)
		}
	}

	k.activeCodes = desiredKeycodes
}
