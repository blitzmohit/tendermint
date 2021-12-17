#!/bin/sh
set -euo pipefail
readonly tool="github.com/tendermint/tendermint/scripts/tmjson"

go run "$tool" \
   -output generated.go \
   -pkg consensus \
   -prefix tendermint/ \
   BlockPartMessage=+BlockPart \
   EndHeightMessage=+wal/EndHeightMessage \
   HasVoteMessage=+HasVote \
   NewRoundStepMessage \
   NewValidBlockMessage \
   ProposalMessage=+Proposal \
   ProposalPOLMessage=+ProposalPOL \
   VoteMessage=+Vote \
   VoteSetBitsMessage=+VoteSetBits \
   VoteSetMaj23Message=+VoteSetMaj23 \
   msgInfo=+wal/MsgInfo \
   timeoutInfo=+wal/TimeoutInfo
