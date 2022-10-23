package keymap

import (
	"fmt"
	"machine/usb/hid/keyboard"
)

func GetCode(layer, row, col int) keyboard.Keycode {
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

	if col > 5 {
		col -= 6
	}
	i := row*6 + col
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
