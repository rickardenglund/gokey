package keymapv1

import (
	"fmt"
	"machine/usb/hid/keyboard"
)

type Keymap struct{}

func New() Keymap {
	return Keymap{}
}

func (k Keymap) GetLayerKey() (int, int) {
	return 6, 4
}

func (k Keymap) GetCode(layer, row, col int) keyboard.Keycode {
	l := layer1
	switch layer {
	case 1:
		l = layer1
	case 2:
		l = layer2
	default:
		fmt.Printf("unknown layer: %d", layer)
		return keyboard.Keycode(0)
	}

	if col > 6 {
		col -= 7
	}
	i := row*7 + col
	if i >= len(l) {
		fmt.Printf("c: %d >= %d", i, len(l))

		return keyboard.Keycode(0)
	}

	return l[i]
}

var layer1 = []keyboard.Keycode{
	keyboard.KeyTilde, keyboard.Key1, keyboard.Key2, keyboard.Key3, keyboard.Key4, keyboard.Key5, keyboard.Keycode(0),
	keyboard.KeyTab, keyboard.KeyQ, keyboard.KeyW, keyboard.KeyE, keyboard.KeyR, keyboard.KeyT, keyboard.Keycode(0),
	keyboard.KeyEsc, keyboard.KeyA, keyboard.KeyS, keyboard.KeyD, keyboard.KeyF, keyboard.KeyG, keyboard.Keycode(0),
	keyboard.KeyLeftShift, keyboard.KeyZ, keyboard.KeyX, keyboard.KeyC, keyboard.KeyV, keyboard.KeyB, keyboard.Keycode(0),
	keyboard.KeyLeftCtrl, keyboard.Keycode(0), keyboard.KeyLeftAlt, keyboard.KeyLeftGUI, keyboard.KeySpace, keyboard.KeyReturn, keyboard.Keycode(0),

	keyboard.Keycode(0), keyboard.Key6, keyboard.Key7, keyboard.Key8, keyboard.Key9, keyboard.Key0, keyboard.KeyMinus,
	keyboard.Keycode(0), keyboard.KeyY, keyboard.KeyU, keyboard.KeyI, keyboard.KeyO, keyboard.KeyP, keyboard.KeyEqual,
	keyboard.Keycode(0), keyboard.KeyH, keyboard.KeyJ, keyboard.KeyK, keyboard.KeyL, keyboard.KeySemicolon, keyboard.KeyQuote,
	keyboard.Keycode(0), keyboard.KeyN, keyboard.KeyM, keyboard.KeyComma, keyboard.KeyPeriod, keyboard.KeySlash, keyboard.KeyRightShift,
	keyboard.KeyLeftGUI, keyboard.KeyBackspace, keyboard.KeySpace, keyboard.KeyLeftBrace, keyboard.KeyRightBrace, keyboard.KeyBackslash, keyboard.KeyModifierRightCtrl,
}

var layer2 = []keyboard.Keycode{
	keyboard.KeyF1, keyboard.KeyF2, keyboard.KeyF3, keyboard.KeyF4, keyboard.KeyF5, keyboard.KeyF6, keyboard.Keycode(0),
	keyboard.KeyTab, keyboard.KeyQ, keyboard.KeyW, keyboard.KeyE, keyboard.KeyR, keyboard.KeyT, keyboard.Keycode(0),
	keyboard.KeyEsc, keyboard.KeyA, keyboard.KeyS, keyboard.KeyD, keyboard.KeyF, keyboard.KeyG, keyboard.Keycode(0),
	keyboard.KeyLeftShift, keyboard.KeyZ, keyboard.KeyX, keyboard.KeyC, keyboard.KeyV, keyboard.KeyB, keyboard.Keycode(0),
	keyboard.KeyLeftCtrl, keyboard.Keycode(0), keyboard.KeyLeftAlt, keyboard.KeyLeftGUI, keyboard.KeySpace, keyboard.KeyReturn, keyboard.Keycode(0),

	keyboard.Keycode(0), keyboard.KeyF7, keyboard.KeyF7, keyboard.KeyF8, keyboard.KeyF9, keyboard.KeyF10, keyboard.KeyF11, keyboard.KeyF12,
	keyboard.Keycode(0), keyboard.KeyY, keyboard.KeyU, keyboard.KeyI, keyboard.KeyO, keyboard.KeyP, keyboard.KeyEqual,
	keyboard.Keycode(0), keyboard.KeyLeft, keyboard.KeyDown, keyboard.KeyUp, keyboard.KeyRight, keyboard.KeySemicolon, keyboard.KeyQuote,
	keyboard.Keycode(0), keyboard.KeyN, keyboard.KeyM, keyboard.KeyComma, keyboard.KeyPeriod, keyboard.KeySlash, keyboard.KeyRightShift,
	keyboard.KeyLeftGUI, keyboard.KeyBackspace, keyboard.KeyLeftGUI, keyboard.KeyLeftBrace, keyboard.KeyRightBrace, keyboard.KeyBackslash, keyboard.KeyModifierRightCtrl,
}
