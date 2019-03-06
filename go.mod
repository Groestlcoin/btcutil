module github.com/btcsuite/btcutil

require (
	github.com/Groestlcoin/go-groestl-hash v0.0.0-20181012171753-790653ac190c
	github.com/btcsuite/btcd v0.0.0-20190213025234-306aecffea32
	golang.org/x/crypto v0.0.0-20170930174604-9419663f5a44
)

replace (
	github.com/btcsuite/btcd => ../grsd
	github.com/btcsuite/btcutil => ./
)