package lotustool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/lotus/chain/consensus/filcns"
)
func JsonParams(code cid.Cid, method abi.MethodNum, params []byte) (string, error) {
	p, err := stmgr.GetParamType(filcns.NewActorRegistry(), code, method) // todo use api for correct actor registry
	if err != nil {
		return "", err
	}

	if err := p.UnmarshalCBOR(bytes.NewReader(params)); err != nil {
		return "", err
	}

	b, err := json.MarshalIndent(p, "", "  ")
	return string(b), err
}
