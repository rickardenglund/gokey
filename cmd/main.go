package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	l := machine.LED
	l.Configure(machine.PinConfig{Mode: machine.PinOutput})

	c1 := machine.GP15
	c2 := machine.GP14
	c3 := machine.GP13

	r1 := machine.GP16

	c1.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	c2.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	c3.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	r1.Configure(machine.PinConfig{Mode: machine.PinOutput})

	r1.High()
	fmt.Printf("hejhej\n")
	for {
		time.Sleep(500 * time.Millisecond)

		fmt.Printf("%t, ", c1.Get())
		fmt.Printf("%t", c2.Get())
		fmt.Printf("%t\r\n", c3.Get())
	}
}
