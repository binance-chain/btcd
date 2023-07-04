// Copyright (c) 2017 All BNB Chain Developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netsync

import (
	"github.com/binance-chain/btcd/blockchain"
	"github.com/binance-chain/btcd/chaincfg"
	"github.com/binance-chain/btcd/chaincfg/chainhash"
	"github.com/binance-chain/btcd/mempool"
	"github.com/binance-chain/btcd/peer"
	"github.com/binance-chain/btcd/wire"
	"github.com/binance-chain/btcutil"
)

// PeerNotifier exposes methods to notify peers of status changes to
// transactions, blocks, etc. Currently server (in the main package) implements
// this interface.
type PeerNotifier interface {
	AnnounceNewTransactions(newTxs []*mempool.TxDesc)

	UpdatePeerHeights(latestBlkHash *chainhash.Hash, latestHeight int32, updateSource *peer.Peer)

	RelayInventory(invVect *wire.InvVect, data interface{})

	TransactionConfirmed(tx *btcutil.Tx)
}

// Config is a configuration struct used to initialize a new SyncManager.
type Config struct {
	PeerNotifier PeerNotifier
	Chain        *blockchain.BlockChain
	TxMemPool    *mempool.TxPool
	ChainParams  *chaincfg.Params

	DisableCheckpoints bool
	MaxPeers           int

	FeeEstimator *mempool.FeeEstimator
}
