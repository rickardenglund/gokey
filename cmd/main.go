package main

import (
	"gokey/keyboard"
	"gokey/keyboard/keymapv1"
	//"gokey/keyboard/echo"
	"machine"
	"time"
)

func main() {
	l := machine.LED
	l.Configure(machine.PinConfig{Mode: machine.PinOutput})

	cols := setuptPins(machine.PinInputPulldown,
		machine.GP5,
		machine.GP15, // 1
		machine.GP6,
		machine.GP7,
		machine.GP8, // 5
		machine.GP9,
		machine.GP28,

		machine.GP21,
		machine.GP20,
		machine.GP19,
		machine.GP18,
		machine.GP22,
		machine.GP17,
		machine.GP16,
	)

	rows := setuptPins(machine.PinOutput,
		machine.GP26,
		machine.GP27,
		machine.GP4,
		machine.GP3,
		machine.GP2,

		machine.GP10,
		machine.GP11,
		machine.GP12,
		machine.GP13,
		machine.GP14,
	)

	kb := keyboard.New(l, keymapv1.New())
	//kb := echo.New()
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
		time.Sleep(7 * time.Millisecond)
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

func setuptPins(mode machine.PinMode, pins ...machine.Pin) []machine.Pin {
	for i := range pins {
		pins[i].Configure(machine.PinConfig{Mode: mode})
	}

	return pins
}
