package main

import (
	"fmt"

	"github.com/cassaram/ece1896/hardware/mcp23s17"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)

func main() {
	host.Init()
	reg, err := spireg.Open("/dev/spidev0.0")
	if err != nil {
		fmt.Print(err)
	}
	// Connect
	c, err := reg.Connect(physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		fmt.Print(err)
	}

	ledReg := mcp23s17.NewMCP23S17(&c, 0)
	ledReg.Init()
	ledReg.ConfigurePin(0, 0, true, false, false, false)
	ledReg.ConfigurePin(0, 1, true, false, false, false)
	ledReg.ConfigurePin(0, 2, true, false, false, false)
	ledReg.ConfigurePin(0, 3, true, false, false, false)
	ledReg.ConfigurePin(0, 4, true, false, false, false)
	ledReg.ConfigurePin(0, 5, true, false, false, false)
	ledReg.ConfigurePin(0, 6, true, false, false, false)
	ledReg.ConfigurePin(0, 7, true, false, false, false)
	ledReg.ConfigurePin(1, 0, true, false, false, false)
	ledReg.ConfigurePin(1, 1, true, false, false, false)
	ledReg.ConfigurePin(1, 2, true, false, false, false)
	ledReg.ConfigurePin(1, 3, true, false, false, false)
	ledReg.ConfigurePin(1, 4, true, false, false, false)
	ledReg.ConfigurePin(1, 5, true, false, false, false)
	ledReg.ConfigurePin(1, 6, true, false, false, false)
	ledReg.ConfigurePin(1, 7, true, false, false, false)
	ledReg.SetPin(0, 0, true)
	ledReg.SetPin(0, 1, true)
	ledReg.SetPin(0, 2, true)
	ledReg.SetPin(0, 3, true)
	ledReg.SetPin(0, 4, true)
	ledReg.SetPin(0, 5, true)
	ledReg.SetPin(0, 6, true)
	ledReg.SetPin(0, 7, true)
	ledReg.SetPin(1, 0, true)
	ledReg.SetPin(1, 1, true)
	ledReg.SetPin(1, 2, true)
	ledReg.SetPin(1, 3, true)
	ledReg.SetPin(1, 4, true)
	ledReg.SetPin(1, 5, true)
	ledReg.SetPin(1, 6, true)
	ledReg.SetPin(1, 7, true)
}
