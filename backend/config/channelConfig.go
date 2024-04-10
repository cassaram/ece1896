package config

import (
	"fmt"
	"strconv"
)

type ChannelConfig struct {
	Name          string           `json:"name"`
	ID            uint64           `json:"id"`
	Color         string           `json:"color"`
	InputCfg      InputConfig      `json:"input_cfg"`
	EQCfg         EQConfig         `json:"eq_cfg"`
	CompressorCfg CompressorConfig `json:"compressor_cfg"`
	GateCfg       GateConfig       `json:"gate_cfg"`
	PFL           bool             `json:"pfl"`
	AFL           bool             `json:"afl"`
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
		PFL:           false,
		AFL:           false,
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
	case "pfl":
		return strconv.FormatBool(c.PFL), nil
	case "afl":
		return strconv.FormatBool(c.AFL), nil
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
	case "pfl":
		val, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		c.PFL = val
		return nil
	case "afl":
		val, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		c.AFL = val
		return nil
	default:
		return fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}
