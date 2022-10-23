package keyboard

import (
	"fmt"
	"machine/usb/hid/keyboard"
)

type Coordinates struct {
	Row, Col int
}

type Keyboard interface {
	PressedKeys(map[Coordinates]bool)
}

func New() Keyboard {
	hidKeyboard := keyboard.New()

	return &kb{
		keysDown: map[Coordinates]bool{},
		hk:       hidKeyboard,
	}
}

type kbi interface {
	Down(keyboard.Keycode) error
	Up(keyboard.Keycode) error
}
type kb struct {
	keysDown map[Coordinates]bool
	hk       kbi
}

func (k *kb) PressedKeys(cs map[Coordinates]bool) {
	printed := false
	for c := range cs {
		if !k.keysDown[c] {
			// fmt.Printf("D(%d,%d) ", c.Row, c.Col)
			k.hk.Down(getCode(c))
			printed = true
		}
	}

	for c := range k.keysDown {
		if !cs[c] {
			fmt.Printf("U(%d, %d) ", c.Row, c.Col)
			k.hk.Up(getCode(c))
			// printed = true
		}
	}

	k.keysDown = cs

	if printed {
		fmt.Print("\r\n\r\n")
	}
}

func getCode(c Coordinates) keyboard.Keycode {
	if c.Col > 5 {
		c.Col -= 6
	}
	i := c.Row*6 + c.Col
	if i >= len(keymap) {
		fmt.Printf("c: %d >= %d", i, len(keymap))

		return keyboard.Keycode(0)
	}

	return keymap[i]
}

var keymap = []keyboard.Keycode{
	keyboard.KeyTilde, keyboard.Key1, keyboard.Key2, keyboard.Key3, keyboard.Key4, keyboard.Key5,
	keyboard.KeyTab, keyboard.KeyQ, keyboard.KeyW, keyboard.KeyE, keyboard.KeyR, keyboard.KeyT,
	keyboard.KeyEsc, keyboard.KeyA, keyboard.KeyS, keyboard.KeyD, keyboard.KeyF, keyboard.KeyG,
	keyboard.KeyLeftShift, keyboard.KeyZ, keyboard.KeyX, keyboard.KeyC, keyboard.KeyV, keyboard.KeyB,
	keyboard.KeyLeftCtrl, keyboard.KeyLeftAlt, keyboard.KeyLeftGUI, keyboard.KeySpace, keyboard.KeyReturn, keyboard.Keycode(0),

	keyboard.Key6, keyboard.Key7, keyboard.Key8, keyboard.Key9, keyboard.Key0, keyboard.KeyMinus,
	keyboard.KeyY, keyboard.KeyU, keyboard.KeyI, keyboard.KeyO, keyboard.KeyP, keyboard.KeyEqual,
	keyboard.KeyH, keyboard.KeyJ, keyboard.KeyK, keyboard.KeyL, keyboard.KeySemicolon, keyboard.KeyQuote,
	keyboard.KeyN, keyboard.KeyM, keyboard.KeyComma, keyboard.KeyPeriod, keyboard.KeySlash, keyboard.KeyRightShift,
	keyboard.KeyModifierRightCtrl, keyboard.KeyBackspace, keyboard.KeySpace, keyboard.KeyLeftBrace, keyboard.KeyRightBrace, keyboard.KeyBackslash,
}
