package backend

import "github.com/cassaram/ece1896/backend/config"

type Subscriber interface {
	ReloadConfig(cfg config.ShowConfig)
	UpdatePath(path string, cfg config.ShowConfig)
}
