package types

import (
	"encoding/json"

	"github.com/tendermint/tendermint/crypto"
)

// Generate JSON encoding wrappers for the types in this package.
//go:generate -command gen go run github.com/creachadair/misctools/tagtype@latest
//go:generate gen -output generated.go

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

func (DuplicateVoteEvidence) jsonWrapperTag() string     { return "tendermint/DuplicateVoteEvidence" }
func (EventDataBlockSyncStatus) jsonWrapperTag() string  { return "tendermint/event/BlockSyncStatus" }
func (EventDataCompleteProposal) jsonWrapperTag() string { return "tendermint/event/CompleteProposal" }
func (EventDataNewBlock) jsonWrapperTag() string         { return "tendermint/event/NewBlock" }
func (EventDataNewBlockHeader) jsonWrapperTag() string   { return "tendermint/event/NewBlockHeader" }
func (EventDataNewEvidence) jsonWrapperTag() string      { return "tendermint/event/NewEvidence" }
func (EventDataNewRound) jsonWrapperTag() string         { return "tendermint/event/NewRound" }
func (EventDataRoundState) jsonWrapperTag() string       { return "tendermint/event/RoundState" }
func (EventDataStateSyncStatus) jsonWrapperTag() string  { return "tendermint/event/StateSyncStatus" }
func (EventDataString) jsonWrapperTag() string           { return "tendermint/event/ProposalString" }
func (EventDataTx) jsonWrapperTag() string               { return "tendermint/event/Tx" }
func (EventDataValidatorSetUpdates) jsonWrapperTag() string {
	return "tendermint/event/ValidatorSetUpdates"
}
func (EventDataVote) jsonWrapperTag() string { return "tendermint/event/Vote" }
func (LightClientAttackEvidence) jsonWrapperTag() string {
	return "tendermint/LightClientAttackEvidence"
}
