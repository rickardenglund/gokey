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
	if same(cs, k.keysDown) {
		return
	}

	fmt.Printf("%d: ", len(cs))
	for c := range cs {
		fmt.Printf("D(%d,%d) ", c.Row, c.Col)
	}

	fmt.Printf("\r\n")

	k.keysDown = cs
}

func same(xs, ys map[keyboard.Coordinates]bool) bool {
	if len(xs) != len(ys) {
		return false
	}

	for k, v := range xs {
		if v != ys[k] {
			return false
		}
	}
	return true
}
