package gpiowrappers

import (
	"log"
	"strconv"

	"github.com/cassaram/ece1896/backend/config"
	"github.com/cassaram/ece1896/hardware/backend"
	"github.com/cassaram/ece1896/hardware/mcp23s17"
	"periph.io/x/conn/v3/gpio"
)

type SwitchesWrapper struct {
	expander   *mcp23s17.MCP23S17
	backend    *backend.BackendConnection
	intAPin    gpio.PinIn
	intBPin    gpio.PinIn
	portACache byte
	portBCache byte
	logger     *log.Logger
}

func NewSwitchesWrapper(expander *mcp23s17.MCP23S17, backend *backend.BackendConnection, intAPin gpio.PinIn, intBPin gpio.PinIn, logger *log.Logger) *SwitchesWrapper {
	w := SwitchesWrapper{
		expander: expander,
		backend:  backend,
		intAPin:  intAPin,
		intBPin:  intBPin,
		logger:   logger,
	}
	return &w
}

func (w *SwitchesWrapper) Start() {
	w.backend.Subscribe(w)
	go w.intAHandler()
	go w.intBHandler()
}

func (w *SwitchesWrapper) intAHandler() {
	for {
		for w.intAPin.WaitForEdge(-1) {
			val, err := w.expander.ReadPort(0)
			if err != nil {
				w.logger.Printf("error reading port A: %v\n", err)
			}
			w.logger.Printf("INT A Triggered. Old value: %02x, New value: %02x", w.portACache, val)
			w.updateValues(val, w.portBCache)
		}
	}
}

func (w *SwitchesWrapper) intBHandler() {
	for {
		for w.intBPin.WaitForEdge(-1) {
			val, err := w.expander.ReadPort(1)
			if err != nil {
				w.logger.Printf("error reading port B: %v\n", err)
			}
			w.logger.Printf("INT B Triggered. Old value: %02x, New value: %02x", w.portBCache, val)
			w.updateValues(w.portACache, val)
		}
	}
}

func (w *SwitchesWrapper) updateValues(portA byte, portB byte) {
	// MUTE SWITCHES
	if fallingEdge(portA, w.portACache, 0) {
		valStr := strconv.FormatBool(!w.backend.GetConfig().CrosspointCfgs[0][0].Enable)
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "crosspoint_cfgs.0.0.enable",
			Data:   valStr,
		})
	}
	if fallingEdge(portA, w.portACache, 3) {
		valStr := strconv.FormatBool(!w.backend.GetConfig().CrosspointCfgs[1][0].Enable)
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "crosspoint_cfgs.1.0.enable",
			Data:   valStr,
		})
	}
	if fallingEdge(portA, w.portACache, 6) {
		valStr := strconv.FormatBool(!w.backend.GetConfig().CrosspointCfgs[2][0].Enable)
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "crosspoint_cfgs.2.0.enable",
			Data:   valStr,
		})
	}
	if fallingEdge(portB, w.portBCache, 1) {
		valStr := strconv.FormatBool(!w.backend.GetConfig().CrosspointCfgs[3][0].Enable)
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "crosspoint_cfgs.3.0.enable",
			Data:   valStr,
		})
	}

	// PFL SWITCHES
	if fallingEdge(portA, w.portACache, 1) {
		val := config.PFL
		if w.backend.GetConfig().ChannelCfgs[0].Monitor == config.PFL {
			val = config.NONE
		}
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "channel_cfgs.0.monitor",
			Data:   strconv.FormatInt(int64(val), 10),
		})
	}
	if fallingEdge(portA, w.portACache, 4) {
		val := config.PFL
		if w.backend.GetConfig().ChannelCfgs[1].Monitor == config.PFL {
			val = config.NONE
		}
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "channel_cfgs.1.monitor",
			Data:   strconv.FormatInt(int64(val), 10),
		})
	}
	if fallingEdge(portA, w.portACache, 7) {
		val := config.PFL
		if w.backend.GetConfig().ChannelCfgs[2].Monitor == config.PFL {
			val = config.NONE
		}
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "channel_cfgs.2.monitor",
			Data:   strconv.FormatInt(int64(val), 10),
		})
	}
	if fallingEdge(portB, w.portBCache, 3) {
		val := config.PFL
		if w.backend.GetConfig().ChannelCfgs[3].Monitor == config.PFL {
			val = config.NONE
		}
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "channel_cfgs.3.monitor",
			Data:   strconv.FormatInt(int64(val), 10),
		})
	}
	if fallingEdge(portB, w.portBCache, 5) {
		val := config.PFL
		if w.backend.GetConfig().BusCfgs[0].Monitor == config.PFL {
			val = config.NONE
		}
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "bus_cfgs.0.monitor",
			Data:   strconv.FormatInt(int64(val), 10),
		})
	}

	// AFL Switches
	if fallingEdge(portA, w.portACache, 2) {
		val := config.AFL
		if w.backend.GetConfig().ChannelCfgs[0].Monitor == config.AFL {
			val = config.NONE
		}
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "channel_cfgs.0.monitor",
			Data:   strconv.FormatInt(int64(val), 10),
		})
	}
	if fallingEdge(portA, w.portACache, 5) {
		val := config.AFL
		if w.backend.GetConfig().ChannelCfgs[1].Monitor == config.AFL {
			val = config.NONE
		}
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "channel_cfgs.1.monitor",
			Data:   strconv.FormatInt(int64(val), 10),
		})
	}
	if fallingEdge(portB, w.portBCache, 0) {
		val := config.AFL
		if w.backend.GetConfig().ChannelCfgs[2].Monitor == config.AFL {
			val = config.NONE
		}
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "channel_cfgs.2.monitor",
			Data:   strconv.FormatInt(int64(val), 10),
		})
	}
	if fallingEdge(portB, w.portBCache, 3) {
		val := config.AFL
		if w.backend.GetConfig().ChannelCfgs[3].Monitor == config.AFL {
			val = config.NONE
		}
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "channel_cfgs.3.monitor",
			Data:   strconv.FormatInt(int64(val), 10),
		})
	}
	if fallingEdge(portB, w.portBCache, 6) {
		val := config.AFL
		if w.backend.GetConfig().BusCfgs[0].Monitor == config.AFL {
			val = config.NONE
		}
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "bus_cfgs.0.monitor",
			Data:   strconv.FormatInt(int64(val), 10),
		})
	}

	// Update
	w.portACache = portA
	w.portBCache = portB
}

func (w *SwitchesWrapper) ReloadConfig(config config.ShowConfig) {
	// Don't care, we just set values here
}

func (w *SwitchesWrapper) UpdatePath(path string, config config.ShowConfig) {
	// Don't care, we just set values here
}

func fallingEdge(newVal byte, oldVal byte, bit uint8) bool {
	mask := byte(0x01 << bit)
	return (newVal&mask != oldVal&mask) && (newVal&mask == 0)
}
