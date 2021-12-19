package types

import (
	"encoding/json"
	"strconv"

	"github.com/tendermint/tendermint/crypto"
)

// Generate JSON encoding wrappers for the types in this package.
//go:generate sh ./json.sh

func (g *GenesisValidator) UnmarshalJSON(data []byte) error {
	var shim struct {
		Address Address         `json:"address"`
		Power   Int64           `json:"power"`
		PubKey  json.RawMessage `json:"pub_key"`
		Name    string          `json:"name"`
	}
	if err := json.Unmarshal(data, &shim); err != nil {
		return err
	}
	pk, err := crypto.UnmarshalPubKey(shim.PubKey)
	if err != nil {
		return err
	}
	g.Address = shim.Address
	g.Power = shim.Power
	g.PubKey = pk
	g.Name = shim.Name
	return nil
}

// An Int64 implements JSON encoding for an int64 value as a string.
type Int64 int64

func (z *Int64) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*z = Int64(v)
	return nil
}

func (z Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatInt(int64(z), 10))
}
