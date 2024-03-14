package api

import "github.com/google/uuid"

type CommandMethod string

const (
	ERROR     CommandMethod = "error"
	SHOW_GET  CommandMethod = "show_get"
	SHOW_SET  CommandMethod = "show_set"
	SHOW_LOAD CommandMethod = "show_load"
)

type Command struct {
	ClientID    uuid.UUID `json:"client_id"`
	RequestData Request   `json:"request_data"`
}

type Request struct {
	Method CommandMethod `json:"method"`
	Path   string        `json:"path"`
	Data   string        `json:"data"`
}
