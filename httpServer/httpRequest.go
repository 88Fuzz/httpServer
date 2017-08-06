package httpServer

type METHOD int
type REQUEST_TYPE int

const (
	GET METHOD = iota + 1
	HEAD
	POST
	EXTENSION
	INVALID
)

const (
	SIMPLE REQUEST_TYPE = iota
	FULL
)

type request struct {
	requestType REQUEST_TYPE
	version     float32
	method      METHOD
	methodValue string
	path        string
	headers     map[string]string
	body        string
}
