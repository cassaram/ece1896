package mcp23s17

import (
	"sync"

	"periph.io/x/conn/v3/spi"
)

type MCP23S17 struct {
	conn     spi.Conn
	address  uint8
	regCache map[mcp23s17Register]byte
	regMute  sync.Mutex
}

// Address is expected to be the value of the A0, A1, A2 bits (0-7)
// The connection should be established prior to calling
// Init needs to be called before any further commands
func NewMCP23S17(conn *spi.Conn, address uint8) *MCP23S17 {
	d := MCP23S17{
		conn:     *conn,
		address:  address,
		regCache: make(map[mcp23s17Register]byte),
	}

	// Set regCache values to POR values
	d.regMute.Lock()
	defer d.regMute.Unlock()
	d.regCache[IODIRA] = 0xFF
	d.regCache[IODIRB] = 0xFF
	d.regCache[IPOLA] = 0x00
	d.regCache[IPOLB] = 0x00
	d.regCache[GPINTENA] = 0x00
	d.regCache[GPINTENB] = 0x00
	d.regCache[DEFVALA] = 0x00
	d.regCache[DEFVALB] = 0x00
	d.regCache[INTCONA] = 0x00
	d.regCache[INTCONB] = 0x00
	d.regCache[GPPUA] = 0x00
	d.regCache[GPPUB] = 0x00
	d.regCache[INTFA] = 0x00
	d.regCache[INTFB] = 0x00
	d.regCache[INTCAPA] = 0x00
	d.regCache[INTCAPB] = 0x00
	d.regCache[GPIOA] = 0x00
	d.regCache[GPIOB] = 0x00
	d.regCache[OLATA] = 0x00
	d.regCache[OLATB] = 0x00

	return &d
}

// Initialize device IO control registers
func (d *MCP23S17) Init() error {
	d.regMute.Lock()
	defer d.regMute.Unlock()

	// Init GPINTEN
	err := d.write(GPINTENA, 0xFF)
	if err != nil {
		return err
	}
	err = d.write(GPINTENB, 0xFF)
	if err != nil {
		return err
	}

	// Init pull-ups
	err = d.write(GPPUA, 0xFF)
	if err != nil {
		return err
	}
	err = d.write(GPPUB, 0xFF)
	if err != nil {
		return err
	}

	// Init outputs high
	err = d.write(OLATA, 0xFF)
	if err != nil {
		return err
	}
	err = d.write(OLATB, 0xFF)
	if err != nil {
		return err
	}

	// Init port direction
	err = d.write(IODIRA, 0xFF)
	if err != nil {
		return err
	}
	err = d.write(IODIRB, 0xFF)
	if err != nil {
		return err
	}

	// Set both ports to trigger interrupts based off previous value only
	err = d.write(INTCONA, 0x00)
	if err != nil {
		return err
	}
	err = d.write(INTCONB, 0x00)
	if err != nil {
		return err
	}

	return nil
}

func (d *MCP23S17) ConfigurePin(port uint8, pin uint8, output bool, invert bool, interrupt bool, pullup bool) error {
	d.regMute.Lock()
	defer d.regMute.Unlock()

	switch port {
	case 0:
		// PORT A
		if output {
			// Set Direction
			d.write(IODIRA, (d.regCache[IODIRA] &^ (0x01 << pin)))
		} else {
			// Set Direction
			d.write(IODIRA, (d.regCache[IODIRA] | (0x01 << pin)))
			// Set input polarity
			if invert {
				d.write(IPOLA, (d.regCache[IPOLA] | (0x01 << pin)))
			} else {
				d.write(IPOLA, (d.regCache[IPOLA] &^ (0x01 << pin)))
			}
			// Set interrupt enable
			if interrupt {
				d.write(GPINTENA, (d.regCache[GPINTENA] | (0x01 << pin)))
			} else {
				d.write(GPINTENA, (d.regCache[GPINTENA] &^ (0x01 << pin)))
			}
			// Set pull-up
			if pullup {
				d.write(GPPUA, (d.regCache[GPPUA] | (0x01 << pin)))
			} else {
				d.write(GPPUA, (d.regCache[GPPUA] &^ (0x01 << pin)))
			}
		}
	case 1:
		// PORT B
		if output {
			// Set Direction
			d.write(IODIRB, (d.regCache[IODIRB] &^ (0x01 << pin)))
		} else {
			// Set Direction
			d.write(IODIRB, (d.regCache[IODIRB] | (0x01 << pin)))
			// Set input polarity
			if invert {
				d.write(IPOLB, (d.regCache[IPOLB] | (0x01 << pin)))
			} else {
				d.write(IPOLB, (d.regCache[IPOLB] &^ (0x01 << pin)))
			}
			// Set interrupt enable
			if interrupt {
				d.write(GPINTENB, (d.regCache[GPINTENB] | (0x01 << pin)))
			} else {
				d.write(GPINTENB, (d.regCache[GPINTENB] &^ (0x01 << pin)))
			}
			// Set pull-up
			if pullup {
				d.write(GPPUB, (d.regCache[GPPUB] | (0x01 << pin)))
			} else {
				d.write(GPPUB, (d.regCache[GPPUB] &^ (0x01 << pin)))
			}
		}
	}
	return nil
}

func (d *MCP23S17) SetPin(port uint8, pin uint8, value bool) error {
	d.regMute.Lock()
	defer d.regMute.Unlock()

	switch port {
	case 0:
		// Port A
		if value {
			err := d.write(GPIOA, (d.regCache[GPIOA] | (0x01 << pin)))
			if err != nil {
				return err
			}
		} else {
			err := d.write(GPIOA, (d.regCache[GPIOA] &^ (0x01 << pin)))
			if err != nil {
				return err
			}
		}
	case 1:
		// Port B
		if value {
			err := d.write(GPIOB, (d.regCache[GPIOB] | (0x01 << pin)))
			if err != nil {
				return err
			}
		} else {
			err := d.write(GPIOB, (d.regCache[GPIOB] &^ (0x01 << pin)))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *MCP23S17) ReadPort(port uint8) (byte, error) {
	d.regMute.Lock()
	defer d.regMute.Unlock()

	switch port {
	case 0:
		// Port A
		return d.read(GPIOA)
	case 1:
		// Port B
		return d.read(GPIOB)
	}
	return 0, nil
}

func (d *MCP23S17) ReadINTCAP(port uint8) (byte, error) {
	d.regMute.Lock()
	defer d.regMute.Unlock()

	switch port {
	case 0:
		return d.read(INTCAPA)
	case 1:
		return d.read(INTCAPB)
	}
	return 0, nil
}

func (d *MCP23S17) write(register mcp23s17Register, data byte) error {
	// Write
	controlByte := 0x40 | ((d.address & 0x07) << 1)
	txData := []byte{controlByte, byte(register), data}
	rxData := make([]byte, len(txData))
	err := d.conn.Tx(txData, rxData)
	if err != nil {
		return err
	}
	// Cache
	d.regCache[register] = data
	return nil
}

func (d *MCP23S17) read(register mcp23s17Register) (byte, error) {
	controlByte := 0x41 | ((d.address & 0x07) << 1)
	txData := []byte{controlByte, byte(register), 0xFF}
	rxData := make([]byte, len(txData))
	err := d.conn.Tx(txData, rxData)
	if err != nil {
		return 0, err
	}
	// Cache result
	d.regCache[register] = rxData[0]
	return rxData[2], nil
}
