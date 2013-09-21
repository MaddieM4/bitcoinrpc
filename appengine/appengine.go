package appengine

import (
    "appengine"
    "appengine/urlfetch"
    "net/http"
)

type AppEnginePoster struct {
    Address string
}

func (rpc *AppEnginePoster) Post(r io.Reader) (*http.Response, error) {
    context := appengine.NewContext()
    client := urlfetch.Client(context)
    return client.Post(rpc.Address, "application/json", r)
}
