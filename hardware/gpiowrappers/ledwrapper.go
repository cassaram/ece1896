package gpiowrappers

import (
	"strconv"
	"strings"

	"github.com/cassaram/ece1896/backend/config"
	"github.com/cassaram/ece1896/hardware/backend"
	"github.com/cassaram/ece1896/hardware/mcp23s17"
)

type LEDWrapper struct {
	expander *mcp23s17.MCP23S17
	backend  *backend.BackendConnection
}

func NewLEDWrapper(expander *mcp23s17.MCP23S17, backend *backend.BackendConnection) *LEDWrapper {
	w := LEDWrapper{
		expander: expander,
		backend:  backend,
	}
	return &w
}

func (w *LEDWrapper) Start() {
	w.backend.Subscribe(w)
}

func (w *LEDWrapper) ReloadConfig(cfg config.ShowConfig) {
	// Update for each path we support
	paths := []string{
		"channel_cfgs.0.monitor",
		"channel_cfgs.1.monitor",
		"channel_cfgs.2.monitor",
		"channel_cfgs.3.monitor",
		"bus_cfgs.0.monitor",
		"crosspoint_cfgs.0.0.enable",
		"crosspoint_cfgs.1.0.enable",
		"crosspoint_cfgs.2.0.enable",
		"crosspoint_cfgs.3.0.enable",
	}
	for _, path := range paths {
		w.UpdatePath(path, cfg)
	}
}

func (w *LEDWrapper) UpdatePath(path string, cfg config.ShowConfig) {
	switch path {
	case "channel_cfgs.0.monitor":
		valStr, _ := cfg.GetValue(strings.Split(path, "."))
		val, _ := strconv.ParseUint(valStr, 10, 8)
		w.HandleMonitorVal(1, config.ChannelMonitorType(val))
	case "channel_cfgs.1.monitor":
		valStr, _ := cfg.GetValue(strings.Split(path, "."))
		val, _ := strconv.ParseUint(valStr, 10, 8)
		w.HandleMonitorVal(2, config.ChannelMonitorType(val))
	case "channel_cfgs.2.monitor":
		valStr, _ := cfg.GetValue(strings.Split(path, "."))
		val, _ := strconv.ParseUint(valStr, 10, 8)
		w.HandleMonitorVal(3, config.ChannelMonitorType(val))
	case "channel_cfgs.3.monitor":
		valStr, _ := cfg.GetValue(strings.Split(path, "."))
		val, _ := strconv.ParseUint(valStr, 10, 8)
		w.HandleMonitorVal(4, config.ChannelMonitorType(val))
	case "bus_cfgs.0.monitor":
		valStr, _ := cfg.GetValue(strings.Split(path, "."))
		val, _ := strconv.ParseUint(valStr, 10, 8)
		w.HandleMonitorVal(5, config.ChannelMonitorType(val))
	case "crosspoint_cfgs.0.0.enable":
		valStr, _ := cfg.GetValue(strings.Split(path, "."))
		val, _ := strconv.ParseBool(valStr)
		w.SetMuteLED(1, !val)
	case "crosspoint_cfgs.1.0.enable":
		valStr, _ := cfg.GetValue(strings.Split(path, "."))
		val, _ := strconv.ParseBool(valStr)
		w.SetMuteLED(2, !val)
	case "crosspoint_cfgs.2.0.enable":
		valStr, _ := cfg.GetValue(strings.Split(path, "."))
		val, _ := strconv.ParseBool(valStr)
		w.SetMuteLED(3, !val)
	case "crosspoint_cfgs.3.0.enable":
		valStr, _ := cfg.GetValue(strings.Split(path, "."))
		val, _ := strconv.ParseBool(valStr)
		w.SetMuteLED(4, !val)
	}
}

func (w *LEDWrapper) HandleMonitorVal(channel int, monitor config.ChannelMonitorType) {
	switch monitor {
	case config.NONE:
		w.SetPFLLED(channel, false)
		w.SetAFLLED(channel, false)
	case config.PFL:
		w.SetPFLLED(channel, true)
		w.SetAFLLED(channel, false)
	case config.AFL:
		w.SetPFLLED(channel, false)
		w.SetAFLLED(channel, true)
	}
}

func (w *LEDWrapper) SetMuteLED(channel int, on bool) error {
	switch channel {
	case 1:
		return w.expander.SetPin(0, 0, !on)
	case 2:
		return w.expander.SetPin(0, 3, !on)
	case 3:
		return w.expander.SetPin(0, 6, !on)
	case 4:
		return w.expander.SetPin(1, 1, !on)
	case 5:
		// Main
		return w.expander.SetPin(1, 4, !on)
	}
	return nil
}

func (w *LEDWrapper) SetPFLLED(channel int, on bool) error {
	switch channel {
	case 1:
		return w.expander.SetPin(0, 1, !on)
	case 2:
		return w.expander.SetPin(0, 4, !on)
	case 3:
		return w.expander.SetPin(0, 7, !on)
	case 4:
		return w.expander.SetPin(1, 2, !on)
	case 5:
		// Main
		return w.expander.SetPin(1, 5, !on)
	}
	return nil
}

func (w *LEDWrapper) SetAFLLED(channel int, on bool) error {
	switch channel {
	case 1:
		return w.expander.SetPin(0, 2, !on)
	case 2:
		return w.expander.SetPin(0, 5, !on)
	case 3:
		return w.expander.SetPin(1, 0, !on)
	case 4:
		return w.expander.SetPin(1, 3, !on)
	case 5:
		// Main
		return w.expander.SetPin(1, 6, !on)
	}
	return nil
}
