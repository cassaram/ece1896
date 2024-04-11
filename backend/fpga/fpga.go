package fpga

import (
	"errors"
	"log"
	"strings"

	"github.com/cassaram/ece1896/backend/config"
	"golang.org/x/exp/maps"
	"periph.io/x/conn/v3/i2c"
)

type FPGA struct {
	device          *i2c.Dev
	log             *log.Logger
	monitorBusCache uint8
}

func NewFPGA(device *i2c.Dev, log *log.Logger) *FPGA {
	f := FPGA{
		device:          device,
		log:             log,
		monitorBusCache: 0x00,
	}

	return &f
}

func (f *FPGA) WriteFullConfig(config config.ShowConfig) error {
	errs := error(nil)
	for _, key := range maps.Keys(fpga_memMap) {
		err := f.UpdateFromPath(key, config)
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}

	if errs != nil {
		return errs
	}
	return nil
}

func (f *FPGA) UpdateFromPath(path string, config config.ShowConfig) error {
	// Check if this is a special case
	switch path {
	case "channel_cfgs.0.monitor":
		return f.updateMonitorBus(config)
	case "channel_cfgs.1.monitor":
		return f.updateMonitorBus(config)
	case "channel_cfgs.2.monitor":
		return f.updateMonitorBus(config)
	case "channel_cfgs.3.monitor":
		return f.updateMonitorBus(config)
	case "channel_cfgs.0.compressor_cfg.bypass":
		return f.updateBypassRegister(config)
	case "channel_cfgs.1.compressor_cfg.bypass":
		return f.updateBypassRegister(config)
	case "channel_cfgs.2.compressor_cfg.bypass":
		return f.updateBypassRegister(config)
	case "channel_cfgs.3.compressor_cfg.bypass":
		return f.updateBypassRegister(config)
	case "channel_cfgs.0.gate_cfg.bypass":
		return f.updateBypassRegister(config)
	case "channel_cfgs.1.gate_cfg.bypass":
		return f.updateBypassRegister(config)
	case "channel_cfgs.2.gate_cfg.bypass":
		return f.updateBypassRegister(config)
	case "channel_cfgs.3.gate_cfg.bypass":
		return f.updateBypassRegister(config)
	}

	// Get address of parameter
	addr, ok := fpga_memMap[path]
	if !ok {
		// We don't write that parameter, Ignore call
		return nil
	}

	// Split path
	paths := strings.Split(path, ".")

	// Get value
	valStr, err := config.GetValue(paths)
	if err != nil {
		return err
	}
	valBytes := addr.formatter(valStr)
	b := []byte{addr.address, valBytes[0], valBytes[1], valBytes[2]}

	// Write
	f.device.Write(b)

	return nil
}

func (f *FPGA) updateMonitorBus(config config.ShowConfig) error {
	// Get all values from config and format into uint8
	ch1 := byte(config.ChannelCfgs[0].Monitor)
	ch2 := byte(config.ChannelCfgs[1].Monitor)
	ch3 := byte(config.ChannelCfgs[2].Monitor)
	ch4 := byte(config.ChannelCfgs[3].Monitor)
	data := ((ch1 & 0x03) << 6) | ((ch2 & 0x03) << 4) | ((ch3 & 0x03) << 2) | ((ch4 & 0x03) << 0)

	// Write to device
	b := []byte{0x60, 0x00, 0x00, data}
	f.device.Write(b)

	return nil
}

func (f *FPGA) updateBypassRegister(config config.ShowConfig) error {
	// Format data
	var data byte = 0x00
	for i := 0; i < 4; i++ {
		data <<= 0x01
		if config.ChannelCfgs[i].CompressorCfg.Bypass {
			data |= 0x01
		}
	}
	for i := 0; i < 4; i++ {
		data <<= 0x01
		if config.ChannelCfgs[i].GateCfg.Bypass {
			data |= 0x01
		}
	}

	// Write to device
	b := []byte{0x61, 0x00, 0x00, data}
	f.device.Write(b)

	return nil
}
