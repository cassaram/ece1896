package gpiowrappers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cassaram/ece1896/hardware/backend"
	"github.com/cassaram/ece1896/hardware/mcp3008"
)

type FaderWrapper struct {
	device  *mcp3008.MCP3008
	backend *backend.BackendConnection
	cache   []float64
	leds    *LEDWrapper
}

func NewFaderWrapper(dev *mcp3008.MCP3008, backend *backend.BackendConnection, leds *LEDWrapper) *FaderWrapper {
	w := FaderWrapper{
		device:  dev,
		backend: backend,
		cache:   make([]float64, 5),
		leds:    leds,
	}
	return &w
}

func (w *FaderWrapper) Start(pollRate time.Duration) {
	ticker := time.NewTicker(pollRate)
	go w.rxLoop(ticker)
}

func (w *FaderWrapper) rxLoop(ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			// Time passed
			for i := uint8(0); i < 5; i++ {
				val, err := w.device.ReadPort(i)
				if err != nil {
					fmt.Println(err)
					continue
				}
				if val > w.cache[i]+0.001 || val < w.cache[i]-0.001 {
					w.cache[i] = val
					lvl := mapFaderLevel(val)
					w.setBackend(i, lvl)
				}
			}
		}
	}
}

func (w *FaderWrapper) setBackend(channel uint8, value float64) {
	path := fmt.Sprintf("crosspoint_cfgs.%d.0.volume", channel)
	if channel == 4 {
		path = "bus_cfgs.0.volume"
	}
	w.backend.Send(backend.Request{
		Method: backend.SHOW_SET,
		Path:   path,
		Data:   strconv.FormatFloat(value, 'f', 0, 64),
	})
}

func mapFaderLevel(val float64) float64 {
	return (((val - 0) * 110) / 1) - 100
}
