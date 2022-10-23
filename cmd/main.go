package main

import (
	"gokey/keyboard"
	"machine"
	"time"
)

func main() {
	l := machine.LED
	l.Configure(machine.PinConfig{Mode: machine.PinOutput})

	c1 := setupCol(machine.GP4)
	c2 := setupCol(machine.GP5)
	c3 := setupCol(machine.GP6)
	c4 := setupCol(machine.GP7)
	c5 := setupCol(machine.GP8)
	c6 := setupCol(machine.GP9)

	c7 := setupCol(machine.GP15)
	c8 := setupCol(machine.GP14)
	c9 := setupCol(machine.GP13)
	c10 := setupCol(machine.GP12)
	c11 := setupCol(machine.GP11)
	c12 := setupCol(machine.GP10)
	cols := []machine.Pin{
		c1, c2, c3, c4, c5, c6,
		c7, c8, c9, c10, c11, c12,
	}

	r1 := setupRow(machine.GP21)
	r2 := setupRow(machine.GP22)
	r3 := setupRow(machine.GP26)
	r4 := setupRow(machine.GP27)
	r5 := setupRow(machine.GP28)

	r6 := setupRow(machine.GP16)
	r7 := setupRow(machine.GP17)
	r8 := setupRow(machine.GP18)
	r9 := setupRow(machine.GP19)
	r10 := setupRow(machine.GP20)

	rows := []machine.Pin{
		r1, r2, r3, r4, r5,
		r6, r7, r8, r9, r10,
	}

	kb := keyboard.New(l)
	for {
		pressed := map[keyboard.Coordinates]bool{}
		for r := range rows {
			enableRow(rows, r)

			for i := range cols {
				if cols[i].Get() {
					pressed[keyboard.Coordinates{Row: r, Col: i}] = true
				}
			}

		}

		kb.PressedKeys(pressed)
		time.Sleep(5 * time.Millisecond)
	}
}

func enableRow(rows []machine.Pin, r int) {
	for i := range rows {
		if r == i {
			rows[i].High()
		} else {
			rows[i].Low()
		}
	}
}

func setupCol(pin machine.Pin) machine.Pin {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	return pin
}

func setupRow(pin machine.Pin) machine.Pin {
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return pin
}
