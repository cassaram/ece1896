package config

import (
	"fmt"
	"strconv"
)

type InputConfig struct {
	InvertPhase bool    `json:"invert_phase"`
	StereoGroup bool    `json:"stereo_group"`
	Gain        float64 `json:"gain"`
}

func NewInputConfig() *InputConfig {
	c := InputConfig{
		InvertPhase: false,
		StereoGroup: false,
		Gain:        0,
	}
	return &c
}

func (c *InputConfig) GetValue(path []string) (string, error) {
	switch path[0] {
	case "invert_phase":
		return strconv.FormatBool(c.InvertPhase), nil
	case "stereo_group":
		return strconv.FormatBool(c.StereoGroup), nil
	case "gain":
		return strconv.FormatFloat(c.Gain, 'f', -1, 64), nil
	default:
		return "", fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}

func (c *InputConfig) SetValue(path []string, value string) error {
	switch path[0] {
	case "invert_phase":
		val, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		c.InvertPhase = val
		return nil
	case "stereo_group":
		val, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		c.StereoGroup = val
		return nil
	case "gain":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Gain = val
		return nil
	default:
		return fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}
