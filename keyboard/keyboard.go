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
		keysDown: map[Coordinates]bool{},
		hk:       hidKeyboard,
		led:      ledPin,
	}
}

type kbi interface {
	Down(keyboard.Keycode) error
	Up(keyboard.Keycode) error
}
type kb struct {
	keysDown map[Coordinates]bool
	hk       kbi
	led      machine.Pin
}

func (k *kb) PressedKeys(cs map[Coordinates]bool) {
	l := 1

	if cs[Coordinates{Row: 4, Col: 5}] {
		l = 2
		k.led.High()
	} else {
		k.led.Low()
	}

	for c := range cs {
		if !k.keysDown[c] {
			k.hk.Down(keymap.GetCode(l, c.Row, c.Col))
		}
	}

	for c := range k.keysDown {
		if !cs[c] {
			k.hk.Up(keymap.GetCode(l, c.Row, c.Col))
		}
	}

	k.keysDown = cs
}
