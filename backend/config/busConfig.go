package config

import (
	"fmt"
	"strconv"
)

type BusConfig struct {
	Name    string             `json:"name"`
	ID      uint64             `json:"id"`
	Master  bool               `json:"master"`
	Monitor ChannelMonitorType `json:"monitor"`
	Volume  float64            `json:"volume"`
	Pan     int64              `json:"pan"`
}

func NewBusConfig(id uint64) *BusConfig {
	c := BusConfig{
		Name: fmt.Sprintf("Bus %d", id+1),
		ID:   id,
	}
	return &c
}

func (c *BusConfig) GetValue(path []string) (string, error) {
	switch path[0] {
	case "name":
		return c.Name, nil
	case "id":
		return strconv.FormatUint(c.ID, 10), nil
	case "master":
		return strconv.FormatBool(c.Master), nil
	case "monitor":
		return strconv.FormatUint(uint64(c.Monitor), 10), nil
	case "volume":
		return strconv.FormatFloat(c.Volume, 'f', -1, 64), nil
	case "pan":
		return strconv.FormatInt(c.Pan, 10), nil
	default:
		return "", fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}

func (c *BusConfig) SetValue(path []string, value string) error {
	switch path[0] {
	case "name":
		c.Name = value
		return nil
	case "master":
		val, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		c.Master = val
		return nil
	case "monitor":
		val, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		c.Monitor = ChannelMonitorType(val)
		return nil
	case "volume":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Volume = val
		return nil
	case "pan":
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		c.Pan = val
		return nil
	default:
		return fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}
