package keyboard

import (
	"machine"
	"machine/usb/hid/keyboard"
)

type Coordinates struct {
	Row, Col int
}

type Keyboard interface {
	PressedKeys(map[Coordinates]bool)
}

type Keymap interface {
	GetLayerKey() (int, int)
	GetCode(layer, row, col int) keyboard.Keycode
}

func New(ledPin machine.Pin, keymap Keymap) Keyboard {
	hidKeyboard := keyboard.New()

	return &kb{
		activeCodes: map[keyboard.Keycode]bool{},
		hk:          hidKeyboard,
		led:         ledPin,
		keymap:      keymap,
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
	keymap      Keymap
}

func (k *kb) PressedKeys(cs map[Coordinates]bool) {
	l := 1

	lCol, lRow := k.keymap.GetLayerKey()
	if cs[Coordinates{Row: lRow, Col: lCol}] {
		l = 2
		k.led.High()
	} else {
		k.led.Low()
	}

	desiredKeycodes := map[keyboard.Keycode]bool{}
	for c := range cs {
		code := k.keymap.GetCode(l, c.Row, c.Col)
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
