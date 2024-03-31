package config

import (
	"fmt"
	"strconv"
)

type ShowConfig struct {
	Name            string               `json:"name"`
	FileName        string               `json:"filename"`
	SelectedChannel uint64               `json:"selected_channel"`
	ChannelCfgs     []ChannelConfig      `json:"channel_cfgs"`
	BusCfgs         []BusConfig          `json:"bus_cfgs"`
	CrosspointCfgs  [][]CrosspointConfig `json:"crosspoint_cfgs"`
}

func NewShowConfig(name string, filename string, channels uint64, busses uint64) *ShowConfig {
	c := ShowConfig{
		Name:           name,
		FileName:       filename,
		ChannelCfgs:    make([]ChannelConfig, channels),
		BusCfgs:        make([]BusConfig, busses),
		CrosspointCfgs: make([][]CrosspointConfig, channels),
	}
	for i := uint64(0); i < channels; i++ {
		c.CrosspointCfgs[i] = make([]CrosspointConfig, busses)
		c.ChannelCfgs[i] = *NewChannelConfig(i)
		for j := uint64(0); j < busses; j++ {
			if i == 0 {
				c.BusCfgs[j] = *NewBusConfig(j)
				if j == 0 {
					c.BusCfgs[j].Master = true
				} else if j == busses-1 {
					c.BusCfgs[j].PFL = true
					c.BusCfgs[j].AFL = true
				}
			}
			c.CrosspointCfgs[i][j] = *NewCrosspointConfig(i, j)
		}
	}
	return &c
}

func (c *ShowConfig) GetValue(path []string) (string, error) {
	switch path[0] {
	case "name":
		return c.Name, nil
	case "filename":
		return c.FileName, nil
	case "selected_channel":
		return strconv.FormatUint(c.SelectedChannel, 10), nil
	case "channel_cfgs":
		id, err := strconv.Atoi(path[1])
		if err != nil {
			return "", err
		}
		return c.ChannelCfgs[id].GetValue(path[2:])
	case "bus_cfgs":
		id, err := strconv.Atoi(path[1])
		if err != nil {
			return "", err
		}
		return c.BusCfgs[id].GetValue(path[2:])
	case "crosspoint_cfgs":
		input_id, err := strconv.Atoi(path[1])
		if err != nil {
			return "", err
		}
		bus_id, err := strconv.Atoi(path[2])
		if err != nil {
			return "", err
		}
		return c.CrosspointCfgs[input_id][bus_id].GetValue(path[3:])
	default:
		return "", fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}

func (c *ShowConfig) SetValue(path []string, value string) error {
	switch path[0] {
	case "name":
		c.Name = value
		return nil
	case "filename":
		c.FileName = value
		return nil
	case "selected_channel":
		val, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		c.SelectedChannel = val
		return nil
	case "channel_cfgs":
		id, err := strconv.Atoi(path[1])
		if err != nil {
			return err
		}
		return c.ChannelCfgs[id].SetValue(path[2:], value)
	case "bus_cfgs":
		id, err := strconv.Atoi(path[1])
		if err != nil {
			return err
		}
		return c.BusCfgs[id].SetValue(path[2:], value)
	case "crosspoint_cfgs":
		channel_id, err := strconv.Atoi(path[1])
		if err != nil {
			return err
		}
		bus_id, err := strconv.Atoi(path[2])
		if err != nil {
			return err
		}
		return c.CrosspointCfgs[channel_id][bus_id].SetValue(path[3:], value)
	default:
		return fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}
