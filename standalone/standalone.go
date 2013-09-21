package standalone

import "net/http"

type StandalonePoster struct {
    Address string
}

func (rpc *StandalonePoster) Post(r io.Reader) (*http.Response, error) {
    return http.Post(rpc.Address, "application/json", r)
}
