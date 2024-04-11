package hweq

import (
	"github.com/cassaram/ece1896/backend/ad5254"
	"github.com/cassaram/ece1896/backend/config"
)

type HardwareEQ struct {
	devices []*ad5254.AD5254
}

func NewHardwareEQ(devices []*ad5254.AD5254) *HardwareEQ {
	d := HardwareEQ{
		devices: devices,
	}

	return &d
}

// Endpoint for updating the all EQs from a config file
func (h *HardwareEQ) WriteFullConfig(config config.ShowConfig) error {
	return nil
}

// Endpoint for updating the EQs from path variables
func (h *HardwareEQ) UpdateFromPath(path string, config config.ShowConfig) error {
	return nil
}
