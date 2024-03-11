package config

import (
	"fmt"
	"strconv"
)

type GateConfig struct {
	Bypass      bool    `json:"bypass"`
	Threshold   float64 `json:"threshold"`
	Depth       float64 `json:"depth"`
	AttackTime  float64 `json:"attack_time"`
	HoldTime    float64 `json:"hold_time"`
	ReleaseTime float64 `json:"release_time"`
}

func NewGateConfig() *GateConfig {
	c := GateConfig{
		Bypass:      true,
		Threshold:   -42,
		Depth:       0,
		AttackTime:  16,
		HoldTime:    0,
		ReleaseTime: 75,
	}
	return &c
}

func (c *GateConfig) GetValue(path []string) (string, error) {
	switch path[0] {
	case "bypass":
		return strconv.FormatBool(c.Bypass), nil
	case "threshold":
		return strconv.FormatFloat(c.Threshold, 'f', -1, 64), nil
	case "depth":
		return strconv.FormatFloat(c.Depth, 'f', -1, 64), nil
	case "attack_time":
		return strconv.FormatFloat(c.AttackTime, 'f', -1, 64), nil
	case "hold_time":
		return strconv.FormatFloat(c.HoldTime, 'f', -1, 64), nil
	case "release_time":
		return strconv.FormatFloat(c.ReleaseTime, 'f', -1, 64), nil
	default:
		return "", fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}

func (c *GateConfig) SetValue(path []string, value string) error {
	switch path[0] {
	case "bypass":
		val, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		c.Bypass = val
		return nil
	case "threshold":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Threshold = val
		return nil
	case "depth":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Depth = val
		return nil
	case "attack_time":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.AttackTime = val
		return nil
	case "hold_time":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.HoldTime = val
		return nil
	case "release_time":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.ReleaseTime = val
		return nil
	default:
		return fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}
