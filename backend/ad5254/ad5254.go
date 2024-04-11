package ad5254

import "periph.io/x/conn/v3/i2c"

type AD5254 struct {
	device *i2c.Dev
}

func NewAD5254(device *i2c.Dev) *AD5254 {
	d := AD5254{
		device: device,
	}

	return &d
}

func (d *AD5254) WriteToRegister(register uint8, value uint8) error {
	firstByte := 0x00 | (byte(register) & 0x1F)
	data := []byte{firstByte, value}
	_, err := d.device.Write(data)
	return err
}
