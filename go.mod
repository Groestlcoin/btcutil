module github.com/btcsuite/btcutil

require (
	github.com/Groestlcoin/go-groestl-hash v0.0.0-20181012171753-790653ac190c
	github.com/aead/siphash v1.0.1
	github.com/btcsuite/btcd v0.0.0-20190213025234-306aecffea32
	github.com/davecgh/go-spew v0.0.0-20171005155431-ecdeabc65495
	github.com/kkdai/bstream v0.0.0-20161212061736-f391b8402d23
	golang.org/x/crypto v0.0.0-20170930174604-9419663f5a44
)

replace (
	github.com/btcsuite/btcd => github.com/Groestlcoin/grsd v0.0.0-20190306030757-4d5197b2ae94
	github.com/btcsuite/btcutil => ./
)
