package bitcoinrpc

import (
	"io"
	"net/http"
)

type Postable interface {
	Post(r io.Reader) (*http.Response, error)
}

type RPCCall struct {
	Method string
	Id     interface{}
	Params []interface{}
}

type RPCResponse map[string]interface{}
