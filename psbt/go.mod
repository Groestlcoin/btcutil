module github.com/btcsuite/btcutil/psbt

go 1.13

require (
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/davecgh/go-spew v1.1.1
)

replace (
	github.com/btcsuite/btcd => github.com/Groestlcoin/grsd v0.20.1-grs
	github.com/btcsuite/btcutil => github.com/Groestlcoin/grsutil v0.5.0-grs3
)
