package api

type CommandMethod string

const (
	GET    CommandMethod = "GET"
	PUT    CommandMethod = "PUT"
	POST   CommandMethod = "POST"
	DELETE CommandMethod = "DELETE"
)

type Command struct {
	Method CommandMethod
	Path   string
	Data   string
}
