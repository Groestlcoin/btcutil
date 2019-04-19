// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package bloom_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/bloom"
)

func TestMerkleBlock3(t *testing.T) {
	blockStr := "01000000" +
		"79cda856b143d9db2c1caff01d1aecc8630d30625d10e8b4b8b0000000000000" +
		"48d548434a07b3a05651302acae61551c5df997dc3be1d95e24117f3b9972e70" +
		"67291b4d" +
		"4c86041b" +
		"8fa45d63" +
		"01" +
		"01000000010" +
		"00000000000000000000000000000000000000000000000000000000000" +
		"0000ffffffff08044c86041b020a02ffffffff0100f2052a01000000434" +
		"104ecd3229b0571c3be876feaac0442a9f13c5a572742927af1dc623353" +
		"ecf8c202225f64868137a18cdd85cbbb4c74fbccfd4f49639cf1bdc94a5" +
		"672bb15ad5d4cac00000000"
	blockBytes, err := hex.DecodeString(blockStr)
	if err != nil {
		t.Errorf("TestMerkleBlock3 DecodeString failed: %v", err)
		return
	}
	blk, err := btcutil.NewBlockFromBytes(blockBytes)
	if err != nil {
		t.Errorf("TestMerkleBlock3 NewBlockFromBytes failed: %v", err)
		return
	}
	merkleRoot, err := blk.TxHash(0)
	if err != nil {
		t.Errorf("TestMerkleBlock3 blk.TxHash failed: %v", err)
	}
	t.Logf("TestMerkleBlock3 merkle root %v", merkleRoot)

	f := bloom.NewFilter(10, 0, 0.000001, wire.BloomUpdateAll)

	inputStr := "702e97b9f31741e2951dbec37d99dfc55115e6ca2a305156a0b3074a4348d548"
	hash, err := chainhash.NewHashFromStr(inputStr)
	if err != nil {
		t.Errorf("TestMerkleBlock3 NewHashFromStr failed: %v", err)
		return
	}

	f.AddHash(hash)

	mBlock, _ := bloom.NewMerkleBlock(blk, f)

	wantStr := "01000000" +
		"79cda856b143d9db2c1caff01d1aecc8630d30625d10e8b4b8b0000000000000" +
		"48d548434a07b3a05651302acae61551c5df997dc3be1d95e24117f3b9972e70" +
		"67291b4d4c86041b8fa45d63" + "0100000001" +
		"48d548434a07b3a05651302acae61551c5df997dc3be1d95e24117f3b9972e70" +
		"0101"
	want, err := hex.DecodeString(wantStr)
	if err != nil {
		t.Errorf("TestMerkleBlock3 DecodeString failed: %v", err)
		return
	}

	got := bytes.NewBuffer(nil)
	err = mBlock.BtcEncode(got, wire.ProtocolVersion, wire.LatestEncoding)
	if err != nil {
		t.Errorf("TestMerkleBlock3 BtcEncode failed: %v", err)
		return
	}

	if !bytes.Equal(want, got.Bytes()) {
		t.Errorf("TestMerkleBlock3 failed merkle block comparison: "+
			"got %x want %x", got.Bytes(), want)
		return
	}
}
