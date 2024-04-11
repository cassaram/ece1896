package config

import (
	"fmt"
	"strconv"
)

type EQConfig struct {
	Bypass         bool    `json:"bypass"`
	HighPassEnable bool    `json:"high_pass_enable"`
	HighPassGain   float64 `json:"high_pass_gain"`
	MidEnable      bool    `json:"mid_enable"`
	MidGain        float64 `json:"mid_gain"`
	LowPassEnable  bool    `json:"low_pass_enable"`
	LowPassGain    float64 `json:"low_pass_gain"`
}

func NewEQConfig() *EQConfig {
	c := EQConfig{
		Bypass:         true,
		HighPassEnable: false,
		HighPassGain:   0,
		MidEnable:      false,
		MidGain:        0,
		LowPassEnable:  false,
		LowPassGain:    0,
	}
	return &c
}

func (c *EQConfig) GetValue(path []string) (string, error) {
	switch path[0] {
	case "bypass":
		return strconv.FormatBool(c.Bypass), nil
	case "high_pass_enable":
		return strconv.FormatBool(c.HighPassEnable), nil
	case "high_pass_gain":
		return strconv.FormatFloat(c.HighPassGain, 'f', -1, 64), nil
	case "mid_enable":
		return strconv.FormatBool(c.MidEnable), nil
	case "mid_gain":
		return strconv.FormatFloat(c.MidGain, 'f', -1, 64), nil
	case "low_pass_enable":
		return strconv.FormatBool(c.LowPassEnable), nil
	case "low_pass_gain":
		return strconv.FormatFloat(c.LowPassGain, 'f', -1, 64), nil
	default:
		return "", fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}

func (c *EQConfig) SetValue(path []string, value string) error {
	switch path[0] {
	case "bypass":
		val, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		c.Bypass = val
		return nil
	case "high_pass_enable":
		val, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		c.HighPassEnable = val
		return nil
	case "high_pass_gain":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.HighPassGain = val
		return nil
	case "mid_enable":
		val, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		c.MidEnable = val
		return nil
	case "mid_gain":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.MidGain = val
		return nil
	case "low_pass_enable":
		val, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		c.LowPassEnable = val
		return nil
	case "low_pass_gain":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.LowPassGain = val
		return nil
	default:
		return fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}
