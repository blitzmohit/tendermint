package consensus

// Generate JSON encoding wrappers for the types in this package.
//go:generate -command gen go run github.com/creachadair/misctools/tagtype@latest
//go:generate gen -output generated.go

func (BlockPartMessage) jsonWrapperTag() string     { return "tendermint/BlockPart" }
func (EndHeightMessage) jsonWrapperTag() string     { return "tendermint/wal/EndHeightMessage" }
func (HasVoteMessage) jsonWrapperTag() string       { return "tendermint/HasVote" }
func (NewRoundStepMessage) jsonWrapperTag() string  { return "tendermint/NewRoundStepMessage" }
func (NewValidBlockMessage) jsonWrapperTag() string { return "tendermint/NewValidBlockMessage" }
func (ProposalMessage) jsonWrapperTag() string      { return "tendermint/Proposal" }
func (ProposalPOLMessage) jsonWrapperTag() string   { return "tendermint/ProposalPOL" }
func (VoteMessage) jsonWrapperTag() string          { return "tendermint/Vote" }
func (VoteSetBitsMessage) jsonWrapperTag() string   { return "tendermint/VoteSetBits" }
func (VoteSetMaj23Message) jsonWrapperTag() string  { return "tendermint/VoteSetMaj23" }
func (msgInfo) jsonWrapperTag() string              { return "tendermint/wal/MsgInfo" }
func (timeoutInfo) jsonWrapperTag() string          { return "tendermint/wal/TimeoutInfo" }
