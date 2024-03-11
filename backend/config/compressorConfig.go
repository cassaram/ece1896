package config

import (
	"fmt"
	"strconv"
)

type CompressorConfig struct {
	Bypass      bool    `json:"bypass"`
	PreGain     float64 `json:"pre_gain"`
	PostGain    float64 `json:"post_gain"`
	Threshold   float64 `json:"threshold"`
	Ratio       float64 `json:"ratio"`
	AttackTime  float64 `json:"attack_time"`
	ReleaseTime float64 `json:"release_time"`
}

func NewCompressorConfig() *CompressorConfig {
	c := CompressorConfig{
		Bypass:      true,
		PreGain:     0,
		PostGain:    0,
		Threshold:   -18,
		Ratio:       2,
		AttackTime:  5,
		ReleaseTime: 75,
	}
	return &c
}

func (c *CompressorConfig) GetValue(path []string) (string, error) {
	switch path[0] {
	case "bypass":
		return strconv.FormatBool(c.Bypass), nil
	case "pre_gain":
		return strconv.FormatFloat(c.PreGain, 'f', -1, 64), nil
	case "post_gain":
		return strconv.FormatFloat(c.PostGain, 'f', -1, 64), nil
	case "threshold":
		return strconv.FormatFloat(c.Threshold, 'f', -1, 64), nil
	case "ratio":
		return strconv.FormatFloat(c.Ratio, 'f', -1, 64), nil
	case "attack_time":
		return strconv.FormatFloat(c.AttackTime, 'f', -1, 64), nil
	case "release_time":
		return strconv.FormatFloat(c.ReleaseTime, 'f', -1, 64), nil
	default:
		return "", fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}

func (c *CompressorConfig) SetValue(path []string, value string) error {
	switch path[0] {
	case "bypass":
		val, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		c.Bypass = val
		return nil
	case "pre_gain":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.PreGain = val
		return nil
	case "post_gain":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.PostGain = val
		return nil
	case "threshold":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Threshold = val
		return nil
	case "ratio":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Ratio = val
		return nil
	case "attack_time":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.AttackTime = val
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
