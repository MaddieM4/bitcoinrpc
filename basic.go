package bitcoinrpc

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

func PostContents(p Postable, data []byte) ([]byte, error) {
	resp, err := p.Post(strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func Call(p Postable, call_params RPCCall) (RPCResponse, error) {
	data, err := json.Marshal(call_params)
	if err != nil {
		return nil, err
	}

	body, err := PostContents(p, data)
	if err != nil {
		return nil, err
	}

	result := make(RPCResponse)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
