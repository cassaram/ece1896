package config

import (
	"fmt"
	"strconv"
)

type ChannelMonitorType uint8

const (
	NONE ChannelMonitorType = 0
	PFL  ChannelMonitorType = 2
	AFL  ChannelMonitorType = 3
)

type ChannelConfig struct {
	Name          string             `json:"name"`
	ID            uint64             `json:"id"`
	Color         string             `json:"color"`
	InputCfg      InputConfig        `json:"input_cfg"`
	EQCfg         EQConfig           `json:"eq_cfg"`
	CompressorCfg CompressorConfig   `json:"compressor_cfg"`
	GateCfg       GateConfig         `json:"gate_cfg"`
	Monitor       ChannelMonitorType `json:"monitor"`
}

func NewChannelConfig(id uint64) *ChannelConfig {
	c := ChannelConfig{
		Name:          fmt.Sprintf("Channel %d", id+1),
		ID:            id,
		Color:         "#F6F",
		InputCfg:      *NewInputConfig(),
		EQCfg:         *NewEQConfig(),
		CompressorCfg: *NewCompressorConfig(),
		GateCfg:       *NewGateConfig(),
		Monitor:       NONE,
	}
	return &c
}

func (c *ChannelConfig) GetValue(path []string) (string, error) {
	switch path[0] {
	case "name":
		return c.Name, nil
	case "id":
		return strconv.FormatUint(c.ID, 10), nil
	case "color":
		return c.Color, nil
	case "input_cfg":
		return c.InputCfg.GetValue(path[1:])
	case "eq_cfg":
		return c.EQCfg.GetValue(path[1:])
	case "compressor_cfg":
		return c.CompressorCfg.GetValue(path[1:])
	case "gate_cfg":
		return c.GateCfg.GetValue(path[1:])
	case "monitor":
		return strconv.FormatUint(uint64(c.Monitor), 10), nil
	default:
		return "", fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}

func (c *ChannelConfig) SetValue(path []string, value string) error {
	switch path[0] {
	case "name":
		c.Name = value
		return nil
	case "color":
		c.Color = value
		return nil
	case "input_cfg":
		return c.InputCfg.SetValue(path[1:], value)
	case "eq_cfg":
		return c.EQCfg.SetValue(path[1:], value)
	case "compressor_cfg":
		return c.CompressorCfg.SetValue(path[1:], value)
	case "gate_cfg":
		return c.GateCfg.SetValue(path[1:], value)
	case "monitor":
		val, err := strconv.ParseUint(value, 10, 8)
		if err != nil {
			return err
		}
		c.Monitor = ChannelMonitorType(val)
		return nil
	default:
		return fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}
