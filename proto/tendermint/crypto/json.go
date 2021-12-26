package crypto

// Generate JSON encoding wrappers for the types in this package.
//go:generate -command gen go run github.com/creachadair/misctools/tagtype@latest
//go:generate gen -output generated.go

func (PublicKey) jsonWrapperTag() string           { return "tendermint.crypto.PublicKey" }
func (PublicKey_Ed25519) jsonWrapperTag() string   { return "tendermint.crypto.PublicKey_Ed25519" }
func (PublicKey_Secp256K1) jsonWrapperTag() string { return "tendermint.crypto.PublicKey_Secp256K1" }
