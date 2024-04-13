package main

import (
	"fmt"
	"time"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)

func main() {
	// Initialize libraries
	host.Init()
	reg, _ := spireg.Open("/dev/spidev0.0")
	conn, _ := reg.Connect(1000*physic.KiloHertz, spi.Mode0, 8)

	// Run reset pin low to force reset
	driverreg.Init()
	reset := gpioreg.ByName("GPIO4")
	reset.Out(gpio.High)
	time.Sleep(time.Millisecond)
	reset.Out(gpio.Low)
	time.Sleep(time.Millisecond)
	reset.Out(gpio.High)
	time.Sleep(time.Millisecond)

	// Enable HAEN bit in IOCON register
	rxBuffer := make([]byte, 3)
	conn.Tx([]byte{0x40, 0x0A, 0x08}, rxBuffer)
	// Used for MCP23S17 hardware bug with pin A2 when HAEN=0 (which it is by default)
	conn.Tx([]byte{0x48, 0x0A, 0x08}, rxBuffer)

	// Read from IOCON
	conn.Tx([]byte{0x41, 0x0A, 0x00}, rxBuffer)
	fmt.Printf("Addr=0 IOCON=%02x\n", rxBuffer[2])
	conn.Tx([]byte{0x43, 0x0A, 0x00}, rxBuffer)
	fmt.Printf("Addr=1 IOCON=%02x\n", rxBuffer[2])

	// Set value on 0
	conn.Tx([]byte{0x40, 0x09, 0xEE}, rxBuffer)
	// Set value on 1
	conn.Tx([]byte{0x42, 0x0A, 0x08}, rxBuffer)

	// Read data from all chips on bus
	read(conn, 0)
	read(conn, 1)
	read(conn, 2)
	read(conn, 3)
	read(conn, 4)
	read(conn, 5)
	read(conn, 6)
	read(conn, 7)
}

func read(conn spi.Conn, address byte) {
	fmt.Printf("\n----- ADDRESS %d -----\n", address)
	rxBuffer := make([]byte, 3)
	for i := byte(0); i < 0x1A; i++ {
		conn.Tx([]byte{0x41 | (address << 1), i, 0x00}, rxBuffer)
		fmt.Printf("Addr=%d %02x=%02x\n", address, i, rxBuffer[2])
	}
}
