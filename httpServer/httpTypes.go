package httpServer

type METHOD int
type REQUEST_TYPE int
type STATUS_CODE int

const CRLF = "\r\n"
const HTTP = "HTTP/"
const HTTP_VERSION = "1.0"
const HEADER_DELIMITER = ": "

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

const (
	OK                    STATUS_CODE = 200
	CREATED                           = 201
	ACCEPTED                          = 202
	NO_CONTENT                        = 204
	MOVED_PERMANENTLY                 = 301
	MOVED_TEMPORARILY                 = 302
	NOT_MODIFIED                      = 304
	BAD_REQUEST                       = 400
	UNAUTHORIZED                      = 401
	FORBIDDEN                         = 403
	NOT_FOUND                         = 404
	INTERNAL_SERVER_ERROR             = 500
	NOT_IMPLEMENTED                   = 501
	BAD_GATEWAY                       = 502
	SERVICE_UNAVAILABLE               = 503
)

type Request struct {
	requestType REQUEST_TYPE
	version     float32
	method      METHOD
	methodValue string
	path        string
	headers     map[string]string
	body        string
}

type Response struct {
	statusCode STATUS_CODE
	headers    map[string]string
	body       string
}
