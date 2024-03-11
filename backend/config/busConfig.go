package config

import (
	"fmt"
	"strconv"
)

type BusConfig struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
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
	default:
		return "", fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}

func (c *BusConfig) SetValue(path []string, value string) error {
	switch path[0] {
	case "name":
		c.Name = value
		return nil
	default:
		return fmt.Errorf("encountered unexpected path variable %s", path[0])
	}
}
