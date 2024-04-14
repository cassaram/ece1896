package gpiowrappers

import (
	"log"
	"strconv"

	"github.com/cassaram/ece1896/hardware/backend"
	"github.com/cassaram/ece1896/hardware/mcp23s17"
	"periph.io/x/conn/v3/gpio"
)

type EncodersWrapper struct {
	expander   *mcp23s17.MCP23S17
	backend    *backend.BackendConnection
	intAPin    gpio.PinIn
	intBPin    gpio.PinIn
	logger     *log.Logger
	portACache byte
	portBCache byte
}

func NewEncodersWrapper(expander *mcp23s17.MCP23S17, backend *backend.BackendConnection, intAPin gpio.PinIn, intBPin gpio.PinIn, logger *log.Logger) *EncodersWrapper {
	w := EncodersWrapper{
		expander:   expander,
		backend:    backend,
		intAPin:    intAPin,
		intBPin:    intBPin,
		logger:     logger,
		portACache: 0x00,
		portBCache: 0x00,
	}
	return &w
}

func (w *EncodersWrapper) Start() {
	//w.backend.Subscribe(w)
	go w.intAHandler()
	go w.intBHandler()
}

func (w *EncodersWrapper) intAHandler() {
	for {
		for w.intAPin.WaitForEdge(-1) {
			val, err := w.expander.ReadPort(0)
			if err != nil {
				w.logger.Printf("error reading port A: %v\n", err)
			}
			w.logger.Printf("INT A Triggered. Old value: %08b, New value: %08b", w.portACache, val)
			go w.updateValues(val, w.portBCache)
		}
	}
}

func (w *EncodersWrapper) intBHandler() {
	for {
		for w.intBPin.WaitForEdge(-1) {
			val, err := w.expander.ReadPort(1)
			if err != nil {
				w.logger.Printf("error reading port B: %v\n", err)
			}
			w.logger.Printf("INT B Triggered. Old value: %08b, New value: %08b", w.portBCache, val)
			go w.updateValues(w.portACache, val)
		}
	}
}

func (w *EncodersWrapper) updateValues(portA byte, portB byte) {
	// Handle rotary encoders
	if isEdge(portA, w.portACache, 0) || isEdge(portA, w.portACache, 1) {
		switch processRotaryEncoder(portA, portA, w.portACache, w.portACache, 0, 1) {
		case -1:
			oldVal := w.backend.GetConfig().CrosspointCfgs[0][0].Pan
			w.backend.Send(backend.Request{
				Method: backend.SHOW_SET,
				Path:   "crosspoint_cfgs.0.0.pan",
				Data:   strconv.FormatInt(oldVal-1, 10),
			})
		case 1:
			oldVal := w.backend.GetConfig().CrosspointCfgs[0][0].Pan
			w.backend.Send(backend.Request{
				Method: backend.SHOW_SET,
				Path:   "crosspoint_cfgs.0.0.pan",
				Data:   strconv.FormatInt(oldVal+1, 10),
			})
		}
	}
	if isEdge(portA, w.portACache, 3) || isEdge(portA, w.portACache, 4) {
		switch processRotaryEncoder(portA, portA, w.portACache, w.portACache, 3, 4) {
		case -1:
			oldVal := w.backend.GetConfig().CrosspointCfgs[1][0].Pan
			w.backend.Send(backend.Request{
				Method: backend.SHOW_SET,
				Path:   "crosspoint_cfgs.1.0.pan",
				Data:   strconv.FormatInt(oldVal-1, 10),
			})
		case 1:
			oldVal := w.backend.GetConfig().CrosspointCfgs[1][0].Pan
			w.backend.Send(backend.Request{
				Method: backend.SHOW_SET,
				Path:   "crosspoint_cfgs.1.0.pan",
				Data:   strconv.FormatInt(oldVal+1, 10),
			})
		}
	}
	if isEdge(portA, w.portACache, 6) || isEdge(portA, w.portACache, 7) {
		switch processRotaryEncoder(portA, portA, w.portACache, w.portACache, 6, 7) {
		case -1:
			oldVal := w.backend.GetConfig().CrosspointCfgs[2][0].Pan
			w.backend.Send(backend.Request{
				Method: backend.SHOW_SET,
				Path:   "crosspoint_cfgs.2.0.pan",
				Data:   strconv.FormatInt(oldVal-1, 10),
			})
		case 1:
			oldVal := w.backend.GetConfig().CrosspointCfgs[2][0].Pan
			w.backend.Send(backend.Request{
				Method: backend.SHOW_SET,
				Path:   "crosspoint_cfgs.2.0.pan",
				Data:   strconv.FormatInt(oldVal+1, 10),
			})
		}
	}
	if isEdge(portB, w.portBCache, 1) || isEdge(portB, w.portBCache, 2) {
		switch processRotaryEncoder(portB, portB, w.portBCache, w.portBCache, 1, 2) {
		case -1:
			oldVal := w.backend.GetConfig().CrosspointCfgs[3][0].Pan
			w.backend.Send(backend.Request{
				Method: backend.SHOW_SET,
				Path:   "crosspoint_cfgs.3.0.pan",
				Data:   strconv.FormatInt(oldVal-1, 10),
			})
		case 1:
			oldVal := w.backend.GetConfig().CrosspointCfgs[3][0].Pan
			w.backend.Send(backend.Request{
				Method: backend.SHOW_SET,
				Path:   "crosspoint_cfgs.3.0.pan",
				Data:   strconv.FormatInt(oldVal+1, 10),
			})
		}
	}
	// Handle rotary encoder push switches
	if fallingEdge(portA, w.portACache, 2) {
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "selected_channel",
			Data:   "0",
		})
	}
	if fallingEdge(portA, w.portACache, 5) {
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "selected_channel",
			Data:   "1",
		})
	}
	if fallingEdge(portB, w.portBCache, 0) {
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "selected_channel",
			Data:   "2",
		})
	}
	if fallingEdge(portB, w.portBCache, 3) {
		w.backend.Send(backend.Request{
			Method: backend.SHOW_SET,
			Path:   "selected_channel",
			Data:   "3",
		})
	}

	// Update cache
	w.portACache = portA
	w.portBCache = portB
}

func isEdge(new byte, old byte, bit uint8) bool {
	return ((old & bit) != (new & bit))
}

// Returns -1, 0, or 1 depending on change in encoder
func processRotaryEncoder(newA byte, newB byte, oldA byte, oldB byte, bitA uint8, bitB uint8) int8 {
	maskA := byte(0x01 << bitA)
	maskB := byte(0x01 << bitB)
	if (oldA&maskA > 0) && (oldB&maskB > 0) {
		// HIGH HIGH
		if (newA&maskA > 0) && (newB&maskB == 0) {
			return -1
		} else if (newA&maskA == 0) && (newB&maskB > 0) {
			return 1
		}
	}
	if (oldA&maskA > 0) && (oldB&maskB == 0) {
		// HIGH LOW
		if (newA&maskA > 0) && (newB&maskB > 0) {
			return 1
		} else if (newA&maskA == 0) && (newB&maskB == 0) {
			return -1
		}
	}
	if (oldA&maskA == 0) && (oldB&maskB > 0) {
		// LOW HIGH
		if (newA&maskA == 0) && (newB&maskB == 0) {
			return 1
		} else if (newA&maskA > 0) && (newB&maskB > 0) {
			return -1
		}
	}
	if (oldA&maskA == 0) && (oldB&maskB == 0) {
		// LOW LOW
		if (newA&maskA > 0) && (newB&maskB == 0) {
			return 1
		} else if (newA&maskA == 0) && (newB&maskB > 0) {
			return -1
		}
	}
	return 0
}
