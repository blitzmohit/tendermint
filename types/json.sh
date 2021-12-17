#!/bin/sh
set -euo pipefail
readonly tool="github.com/tendermint/tendermint/scripts/tmjson"

go run "$tool" \
   -output generated.go \
   -pkg types \
   -trim EventData -prefix tendermint/event/ \
   EventDataBlockSyncStatus \
   EventDataCompleteProposal \
   EventDataNewBlock \
   EventDataNewBlockHeader \
   EventDataNewEvidence \
   EventDataNewRound \
   EventDataRoundState \
   EventDataStateSyncStatus \
   EventDataString=+ProposalString \
   EventDataTx \
   EventDataValidatorSetUpdates \
   EventDataVote \
   DuplicateVoteEvidence=tendermint/DuplicateVoteEvidence \
   LightClientAttackEvidence=tendermint/LightClientAttackEvidence
