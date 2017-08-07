package httpServer

type Method_t int
type RequestType_t int
type StatusCode_t int

const CRLF = "\r\n"
const HTTP = "HTTP/"
const HTTP_VERSION = "1.0"
const HEADER_DELIMITER = ": "

const (
	GET Method_t = iota + 1
	HEAD
	POST
	EXTENSION
	INVALID
)

const (
	FULL RequestType_t = iota
	SIMPLE
)

const (
	OK                    StatusCode_t = 200
	CREATED                            = 201
	ACCEPTED                           = 202
	NO_CONTENT                         = 204
	MOVED_PERMANENTLY                  = 301
	MOVED_TEMPORARILY                  = 302
	NOT_MODIFIED                       = 304
	BAD_REQUEST                        = 400
	UNAUTHORIZED                       = 401
	FORBIDDEN                          = 403
	NOT_FOUND                          = 404
	INTERNAL_SERVER_ERROR              = 500
	NOT_IMPLEMENTED                    = 501
	BAD_GATEWAY                        = 502
	SERVICE_UNAVAILABLE                = 503
)

type Request struct {
	RequestType  RequestType_t
	Version      float32
	Method       Method_t
	MethodString string
	Path         string
	Headers      map[string]string
	Body         string
}

type Response struct {
	StatusCode StatusCode_t
	Headers    map[string]string
	Body       string
}
