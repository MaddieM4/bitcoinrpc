package bitcoinrpc

import (
	"io"
	"net/http"
)

type StandalonePoster struct {
	Address string
}

func (p StandalonePoster) Post(r io.Reader) (*http.Response, error) {
	return http.Post(p.Address, "application/json", r)
}
