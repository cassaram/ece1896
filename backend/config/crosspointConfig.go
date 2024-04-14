package config

import (
	"fmt"
	"strconv"
)

type CrosspointConfig struct {
	BusID     uint64  `json:"bus_id"`
	ChannelID uint64  `json:"channel_id"`
	Enable    bool    `json:"enable"`
	Pan       int64   `json:"pan"`
	Volume    float64 `json:"volume"`
}

func NewCrosspointConfig(channelID uint64, busID uint64) *CrosspointConfig {
	c := CrosspointConfig{
		BusID:     busID,
		ChannelID: channelID,
		Enable:    false,
		Pan:       0,
		Volume:    -100,
	}
	return &c
}

func (c *CrosspointConfig) GetValue(path []string) (string, error) {
	switch path[0] {
	case "bus_id":
		return strconv.FormatUint(c.BusID, 10), nil
	case "channel_id":
		return strconv.FormatUint(c.ChannelID, 10), nil
	case "enable":
		return strconv.FormatBool(c.Enable), nil
	case "pan":
		return strconv.FormatInt(c.Pan, 10), nil
	case "volume":
		return strconv.FormatFloat(c.Volume, 'f', -1, 64), nil
	default:
		return "", fmt.Errorf("encountered unexpected path varaible %s", path[0])
	}
}

func (c *CrosspointConfig) SetValue(path []string, value string) error {
	switch path[0] {
	case "enable":
		val, err := strconv.ParseBool(value)
		if err != nil {
			return nil
		}
		c.Enable = val
		return nil
	case "pan":
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		c.Pan = val
		return nil
	case "volume":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Volume = min(max(val, -100), 10)
		return nil
	default:
		return fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}
