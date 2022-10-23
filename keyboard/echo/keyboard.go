package echo

import (
	"fmt"
	"gokey/keyboard"
)

func New() keyboard.Keyboard {
	return &kb{keysDown: make(map[keyboard.Coordinates]bool)}

}

type kb struct {
	keysDown map[keyboard.Coordinates]bool
}

func (k *kb) PressedKeys(cs map[keyboard.Coordinates]bool) {
	for c := range cs {
		if !k.keysDown[c] {
			fmt.Printf("D(%d,%d) ", c.Row, c.Col)
		}
	}

	k.keysDown = cs
}
