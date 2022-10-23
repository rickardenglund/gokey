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
	l := layer1
	if cs[Coordinates{Row: 4, Col: 5}] {
		l = layer2
	}

	for c := range cs {
		if !k.keysDown[c] {
			k.hk.Down(getCode(l, c))
		}
	}

	for c := range k.keysDown {
		if !cs[c] {
			k.hk.Up(getCode(l, c))
		}
	}

	k.keysDown = cs
}

func getCode(l []keyboard.Keycode, c Coordinates) keyboard.Keycode {
	if c.Col > 5 {
		c.Col -= 6
	}
	i := c.Row*6 + c.Col
	if i >= len(l) {
		fmt.Printf("c: %d >= %d", i, len(l))

		return keyboard.Keycode(0)
	}

	return l[i]
}

var layer1 = []keyboard.Keycode{
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

var layer2 = []keyboard.Keycode{
	keyboard.KeyF1, keyboard.KeyF2, keyboard.KeyF3, keyboard.KeyF4, keyboard.KeyF5, keyboard.KeyF6,
	keyboard.KeyTab, keyboard.KeyQ, keyboard.KeyW, keyboard.KeyE, keyboard.KeyR, keyboard.KeyT,
	keyboard.KeyEsc, keyboard.KeyA, keyboard.KeyS, keyboard.KeyD, keyboard.KeyF, keyboard.KeyG,
	keyboard.KeyLeftShift, keyboard.KeyZ, keyboard.KeyX, keyboard.KeyC, keyboard.KeyV, keyboard.KeyB,
	keyboard.KeyLeftCtrl, keyboard.KeyLeftAlt, keyboard.KeyLeftGUI, keyboard.KeySpace, keyboard.KeyReturn, keyboard.Keycode(0),

	keyboard.KeyF7, keyboard.KeyF8, keyboard.KeyF9, keyboard.KeyF10, keyboard.KeyF11, keyboard.KeyF12,
	keyboard.KeyY, keyboard.KeyU, keyboard.KeyI, keyboard.KeyO, keyboard.KeyP, keyboard.KeyEqual,
	keyboard.KeyLeft, keyboard.KeyDown, keyboard.KeyUp, keyboard.KeyRight, keyboard.KeySemicolon, keyboard.KeyQuote,
	keyboard.KeyN, keyboard.KeyM, keyboard.KeyComma, keyboard.KeyPeriod, keyboard.KeySlash, keyboard.KeyRightShift,
	keyboard.Keycode(0), keyboard.KeyBackspace, keyboard.KeyLeftGUI, keyboard.KeyLeftBrace, keyboard.KeyRightBrace, keyboard.KeyBackslash,
}
