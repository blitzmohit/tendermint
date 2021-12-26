package sr25519

import "github.com/tendermint/tendermint/crypto"

// Generate JSON encoding wrappers for the types in this package.
//go:generate -command gen go run github.com/creachadair/misctools/tagtype@latest
//go:generate gen -output generated.go

func (PubKey) jsonWrapperTag() string  { return PubKeyName }
func (PrivKey) jsonWrapperTag() string { return PrivKeyName }

const (
	PrivKeyName = "tendermint/PrivKeySr25519"
	PubKeyName  = "tendermint/PubKeySr25519"
)

func init() {
	crypto.RegisterPubKeyType(PubKeyName, PubKey(nil))
	crypto.RegisterPrivKeyType(PrivKeyName, PrivKey{})
}
