package core

type ConfigFile struct {
	Name     string `json:"name"`
	FileName string `json:"filename"`
	Size     int64  `json:"size"`
	ModTime  string `json:"mod_time"`
}
