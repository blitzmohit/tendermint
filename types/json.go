package types

import (
	"encoding/json"

	"github.com/tendermint/tendermint/crypto"
)

// Generate JSON encoding wrappers for the types in this package.
//go:generate sh ./json.sh

func (g *GenesisValidator) UnmarshalJSON(data []byte) error {
	var shim struct {
		Address Address         `json:"address"`
		Power   int64           `json:"power,string"`
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
