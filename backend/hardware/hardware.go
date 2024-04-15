package hardware

import (
	"log"

	"github.com/cassaram/ece1896/backend/ad5254"
	"github.com/cassaram/ece1896/backend/config"
	"github.com/cassaram/ece1896/backend/fpga"
	"github.com/cassaram/ece1896/backend/hweq"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
)

type Hardware struct {
	FPGA       fpga.FPGA
	HardwareEQ hweq.HardwareEQ
	Logger     *log.Logger
}

func NewHardware(log *log.Logger) *Hardware {
	// Get I2C bus
	bus, err := i2creg.Open("/dev/i2c-1")
	if err != nil {
		panic(err)
	}

	// Get AD5254 devices (&i2c.Dev{Bus: bus, Addr: 0x2C}), (&i2c.Dev{Bus: bus, Addr: 0x2D}), (&i2c.Dev{Bus: bus, Addr: 0x2E})
	ad5254devices := []*ad5254.AD5254{
		ad5254.NewAD5254(&i2c.Dev{Bus: bus, Addr: 0x2C}),
		ad5254.NewAD5254(&i2c.Dev{Bus: bus, Addr: 0x2D}),
		ad5254.NewAD5254(&i2c.Dev{Bus: bus, Addr: 0x2E}),
	}

	h := Hardware{
		FPGA:       *fpga.NewFPGA(&i2c.Dev{Bus: bus, Addr: 0x40}, log),
		HardwareEQ: *hweq.NewHardwareEQ(ad5254devices),
	}
	return &h
}

func (h *Hardware) WriteFullConfig(config config.ShowConfig) error {
	err := h.FPGA.WriteFullConfig(config)
	if err != nil {
		return err
	}
	err = h.HardwareEQ.WriteFullConfig(config)
	if err != nil {
		return err
	}
	return nil
}

func (h *Hardware) UpdateFromPath(path string, config config.ShowConfig) error {
	err := h.FPGA.UpdateFromPath(path, config)
	if err != nil {
		return err
	}
	err = h.HardwareEQ.UpdateFromPath(path, config)
	if err != nil {
		return err
	}
	return nil
}
