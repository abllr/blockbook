package namecoin

import (
	"blockbook/bchain"
	"blockbook/bchain/coins/btc"
	"blockbook/bchain/coins/utils"
	"bytes"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
)

const (
	MainnetMagic wire.BitcoinNet = 0xfeb4bef9
)

var (
	MainNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = 52
	MainNetParams.ScriptHashAddrID = 13
	MainNetParams.Bech32HRPSegwit = "nc"

	err := chaincfg.Register(&MainNetParams)
	if err != nil {
		panic(err)
	}
}

// NamecoinParser handle
type NamecoinParser struct {
	*btc.BitcoinParser
}

// NewNamecoinParser returns new NamecoinParser instance
func NewNamecoinParser(params *chaincfg.Params, c *btc.Configuration) *NamecoinParser {
	return &NamecoinParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main Namecoin network,
// and the test Namecoin network
func GetChainParams(chain string) *chaincfg.Params {
	switch chain {
	default:
		return &MainNetParams
	}
}

// ParseBlock parses raw block to our Block struct
// it has special handling for Auxpow blocks that cannot be parsed by standard btc wire parser
func (p *NamecoinParser) ParseBlock(b []byte) (*bchain.Block, error) {
	r := bytes.NewReader(b)
	w := wire.MsgBlock{}
	h := wire.BlockHeader{}
	err := h.Deserialize(r)
	if err != nil {
		return nil, err
	}
	if (h.Version & utils.VersionAuxpow) != 0 {
		if err = utils.SkipAuxpow(r); err != nil {
			return nil, err
		}
	}

	err = utils.DecodeTransactions(r, 0, wire.WitnessEncoding, &w)
	if err != nil {
		return nil, err
	}

	txs := make([]bchain.Tx, len(w.Transactions))
	for ti, t := range w.Transactions {
		txs[ti] = p.TxFromMsgTx(t, false)
	}

	return &bchain.Block{Txs: txs}, nil
}