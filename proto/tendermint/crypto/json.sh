#!/bin/sh
set -euo pipefail
readonly tool="github.com/tendermint/tendermint/scripts/tmjson"

go run "$tool" \
   -output generated.go \
   -pkg crypto \
   -prefix tendermint.crypto. \
   PublicKey \
   PublicKey_Ed25519 \
   PublicKey_Secp256K1
