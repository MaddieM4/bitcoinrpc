package bitcoinrpc

import "testing"

const REMOTE_URL string = `https://rpc.blockchain.info/`

func TestRemote(t *testing.T) {
	url := REMOTE_URL
	poster := StandalonePoster{url}
	call := RPCCall{
		"getinfo",       // Method
		"898djal;j;",    // Id
		[]interface{}{}, // Params
	}

	resp, err := Call(poster, call)
	if err != nil {
		t.Errorf("RPC failed: %v\n", err)
		return
	}
	t.Errorf("%v", resp)
}
