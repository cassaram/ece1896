package mcp23s17

import "periph.io/x/conn/v3/spi"

type MCP23S17 struct {
	conn    spi.Conn
	address uint8
}

// Address is expected to be the value of the A0, A1, A2 bits (0-7)
// The connection should be established prior to calling
// Init needs to be called before any further commands
func NewMCP23S17(conn *spi.Conn, address uint8) *MCP23S17 {
	d := MCP23S17{
		conn:    *conn,
		address: address,
	}
	return &d
}

// Initialize device IO control registers
func (d *MCP23S17) Init() error {
	confRegData := 0x00
	// Disable sequential operations (address pointer does not increment after interactions)
	confRegData |= (0x01 << 5)
	// Enable hardware addressing
	confRegData |= (0x01 << 3)

	err := d.write(IOCON, byte(confRegData))
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
	switch port {
	case 0:
		// PORT A
		if output {
			// Set Direction
			val, err := d.read(IODIRA)
			if err != nil {
				return err
			}
			d.write(IODIRA, (val | (0x01 << pin)))
		} else {
			// Set Direction
			val, err := d.read(IODIRA)
			if err != nil {
				return err
			}
			d.write(IODIRA, (val &^ (0x01 << pin)))
			// Set input polarity
			val, err = d.read(IPOLA)
			if err != nil {
				return err
			}
			if invert {
				d.write(IPOLA, (val | (0x01 << pin)))
			} else {
				d.write(IPOLA, (val &^ (0x01 << pin)))
			}
			// Set interrupt enable
			val, err = d.read(GPINTENA)
			if err != nil {
				return err
			}
			if interrupt {
				d.write(GPINTENA, (val | (0x01 << pin)))
			} else {
				d.write(GPINTENA, (val &^ (0x01 << pin)))
			}
			d.write(INTCONA, 0x00) // Set to use previous value
			// Set pull-up
			val, err = d.read(GPPUA)
			if err != nil {
				return err
			}
			if pullup {
				d.write(GPPUA, (val | (0x01 << pin)))
			} else {
				d.write(GPPUA, (val &^ (0x01 << pin)))
			}
		}
	case 1:
		// PORT B
		if output {
			// Set Direction
			val, err := d.read(IODIRB)
			if err != nil {
				return err
			}
			d.write(IODIRB, (val | (0x01 << pin)))
		} else {
			// Set Direction
			val, err := d.read(IODIRB)
			if err != nil {
				return err
			}
			d.write(IODIRB, (val &^ (0x01 << pin)))
			// Set input polarity
			val, err = d.read(IPOLB)
			if err != nil {
				return err
			}
			if invert {
				d.write(IPOLB, (val | (0x01 << pin)))
			} else {
				d.write(IPOLB, (val &^ (0x01 << pin)))
			}
			// Set interrupt enable
			val, err = d.read(GPINTENB)
			if err != nil {
				return err
			}
			if interrupt {
				d.write(GPINTENB, (val | (0x01 << pin)))
			} else {
				d.write(GPINTENB, (val &^ (0x01 << pin)))
			}
			// Set pull-up
			val, err = d.read(GPPUB)
			if err != nil {
				return err
			}
			if pullup {
				d.write(GPPUB, (val | (0x01 << pin)))
			} else {
				d.write(GPPUB, (val &^ (0x01 << pin)))
			}
		}
	}
	return nil
}

func (d *MCP23S17) SetPin(port uint8, pin uint8, value bool) error {
	switch port {
	case 0:
		// Port A
		val, err := d.read(GPIOA)
		if err != nil {
			return err
		}
		if value {
			err := d.write(GPIOA, (val | (0x01 << pin)))
			if err != nil {
				return err
			}
		} else {
			err := d.write(GPIOA, (val &^ (0x01 << pin)))
			if err != nil {
				return err
			}
		}
	case 1:
		// Port B
		val, err := d.read(GPIOB)
		if err != nil {
			return err
		}
		if value {
			err := d.write(GPIOB, (val | (0x01 << pin)))
			if err != nil {
				return err
			}
		} else {
			err := d.write(GPIOB, (val &^ (0x01 << pin)))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *MCP23S17) ReadPort(port uint8) (byte, error) {
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

func (d *MCP23S17) write(register mcp23s17Register, data byte) error {
	controlByte := 0x40 | ((d.address & 0x07) << 1)
	txData := []byte{controlByte, byte(register), data}
	rxData := make([]byte, len(txData))
	err := d.conn.Tx(txData, rxData)
	if err != nil {
		return err
	}
	return nil
}

func (d *MCP23S17) read(register mcp23s17Register) (byte, error) {
	controlByte := 0x41 | ((d.address & 0x07) << 1)
	txData := []byte{controlByte, byte(register)}
	rxData := make([]byte, len(txData))
	err := d.conn.Tx(txData, rxData)
	if err != nil {
		return 0, err
	}
	return rxData[0], nil
}
