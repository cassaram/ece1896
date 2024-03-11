package core

import (
	"encoding/json"
	"os"

	"github.com/cassaram/ece1896/backend/config"
)

type Core struct {
	RunningConfig config.ShowConfig
	StagedConfig  config.ShowConfig
	configChanged chan bool
	stop          chan bool
}

func NewCore() *Core {
	c := Core{}
	return &c
}

func (c *Core) Run() {
	for {
		select {
		case <-c.configChanged:
			c.RunningConfig = c.StagedConfig
			c.notifyClients()
		case <-c.stop:
			return
		}
	}
}

func (c *Core) notifyClients() {

}

func (c *Core) LoadShowConfig(filepath string) error {
	cfgBytes, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	cfg := config.ShowConfig{}
	err = json.Unmarshal(cfgBytes, &cfg)
	if err != nil {
		return err
	}

	c.StagedConfig = cfg
	c.configChanged <- true

	return nil
}

func (c *Core) SaveShowConfig(filepath string) error {
	cfg, err := json.Marshal(c.RunningConfig)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, cfg, 0644)
	if err != nil {
		return err
	}

	return nil
}
