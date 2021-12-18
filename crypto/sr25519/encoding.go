package sr25519

const (
	PrivKeyName = "tendermint/PrivKeySr25519"
	PubKeyName  = "tendermint/PubKeySr25519"
)

//go:generate -command gen go run github.com/tendermint/tendermint/scripts/tmjson
//go:generate gen -output generated.go -pkg sr25519 PubKey=tendermint/PubKeySr25519
