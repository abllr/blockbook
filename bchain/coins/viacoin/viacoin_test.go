// build unittest

package viacoin

import (
	"blockbook/bchain"
	"blockbook/bchain/coins/btc"
	"encoding/hex"
	"reflect"
	"testing"
)

func TestAddressToOutputScript_Mainnet(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "pubkeyhash1",
			args:    args{address: "VhyGT8kJU9x28dHwjf1jEDG8gMY8yhckDR"},
			want:    "76a91457757edd001d16528c7aa337b314a7bab303ee8088ac",
			wantErr: false,
		},
		{
			name:    "pubkeyhash2",
			args:    args{address: "VdMPvn7vUTSzbYjiMDs1jku9wAh1Ri2Y1A"},
			want:    "76a91424cc424c1e5e977175d2b20012554d39024bd68f88ac",
			wantErr: false,
		},
		{
			name:    "scripthash1",
			args:    args{address: "EUtqKT17p1LdHTDyGL1b2WPJiUFidS6gVq"},
			want:    "a91480c7c8faece680bab1ae81a5969815a05b7565f087",
			wantErr: false,
		},
		{
			name:    "scripthash2",
			args:    args{address: "EMdC3QPzx2MsJG56x2QbSR727dRM73B1rK"},
			want:    "a91431098c569891a8ff1fa11d1cbd3d46ca5e245c6b87",
			wantErr: false,
		},
		{
			name:    "witness_v0_keyhash",
			args:    args{address: "via1"}, //TODO
			want:    "001469de0e878beb0975b202f88a208f99062ef035ef",
			wantErr: false,
		},
	}

	parser := NewViacoinParser(GetChainParams("main"), &btc.Configuration{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.AddressToOutputScript(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressToOutputScript() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			h := hex.EncodeToString(got)
			if !reflect.DeepEqual(h, tt.want) {
				t.Errorf("AddressToOutputScript() = %v, want %v", h, tt.want)
			}
		})
	}
}

var (
	testTx1 bchain.Tx

	testTxPacked1 = "0200000001ddc431a8f5c4e74296de8a6bff796ece148b9bd6827a80ecde8671df41a51fc7000000006a47304402204f929a1e1e40bd352bbd5d3c5ae6c29740e5a8b29dd8c53a15d3eab29aecee7c02206a514e5e4561cfb9330d98f4a4fe1385af56d87eb1d3e1a379d7a50276788cfe0121034a9305644fbcb56d4fc0bc15959b917f7753ae8e581acc97f9cfe771ad1e8249feffffff0200ca9a3b000000001976a91456c7359ed52d61c1ca371d7dc136632148169c5e88acd0e8cc10000000001976a914112e29df5df4866e40ef98e0857036b275380fe088ac6ab94e00"
)

func init() {
	var (
		addr1, addr2 bchain.Address
		err          error
	)
	addr1, err = bchain.NewBaseAddress("VhuffXKNA3j9hgp2JYGrj6uHQ6KUU6zNbS")
	if err == nil {
		addr2, err = bchain.NewBaseAddress("VbZfhUMCUJHDjqjby6ynYFPZSNVYhfe4cK")
	}
	if err != nil {
		panic(err)
	}

	testTx1 = bchain.Tx{
		Hex:       "0200000001ddc431a8f5c4e74296de8a6bff796ece148b9bd6827a80ecde8671df41a51fc7000000006a47304402204f929a1e1e40bd352bbd5d3c5ae6c29740e5a8b29dd8c53a15d3eab29aecee7c02206a514e5e4561cfb9330d98f4a4fe1385af56d87eb1d3e1a379d7a50276788cfe0121034a9305644fbcb56d4fc0bc15959b917f7753ae8e581acc97f9cfe771ad1e8249feffffff0200ca9a3b000000001976a91456c7359ed52d61c1ca371d7dc136632148169c5e88acd0e8cc10000000001976a914112e29df5df4866e40ef98e0857036b275380fe088ac6ab94e00",
		Blocktime: 1530319242,
		Txid:      "d0284c75a389a07cc256e0bb913110d8d8059efd04daa8147ecf2fa0b3bdf6ff",
		LockTime:  5159274,
		Vin: []bchain.Vin{
			{
				ScriptSig: bchain.ScriptSig{
					Hex: "47304402204f929a1e1e40bd352bbd5d3c5ae6c29740e5a8b29dd8c53a15d3eab29aecee7c02206a514e5e4561cfb9330d98f4a4fe1385af56d87eb1d3e1a379d7a50276788cfe0121034a9305644fbcb56d4fc0bc15959b917f7753ae8e581acc97f9cfe771ad1e8249",
				},
				Txid:     "d65a60b340a233d55809ad04d9175e717dc49edf1edacc99034ea3341878fd46",
				Vout:     0,
				Sequence: 4294967294,
			},
		},
		Vout: []bchain.Vout{
			{
				Value: 10,
				N:     0,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a91456c7359ed52d61c1ca371d7dc136632148169c5e88ac",
					Addresses: []string{
						"VhuffXKNA3j9hgp2JYGrj6uHQ6KUU6zNbS",
					},
				},
				Address: addr1,
			},
			{
				Value: 2.818644,
				N:     1,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a914112e29df5df4866e40ef98e0857036b275380fe088ac",
					Addresses: []string{
						"VbZfhUMCUJHDjqjby6ynYFPZSNVYhfe4cK",
					},
				},
				Address: addr2,
			},
		},
	}
}

func Test_PackTx(t *testing.T) {
	type args struct {
		tx        bchain.Tx
		height    uint32
		blockTime int64
		parser    *ViacoinParser
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "viacoin-1",
			args: args{
				tx:        testTx1,
				height:    5159277,
				blockTime: 1530319242,
				parser:    NewViacoinParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    testTxPacked1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.parser.PackTx(&tt.args.tx, tt.args.height, tt.args.blockTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("packTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			h := hex.EncodeToString(got)
			if !reflect.DeepEqual(h, tt.want) {
				t.Errorf("packTx() = %v, want %v", h, tt.want)
			}
		})
	}
}

func Test_UnpackTx(t *testing.T) {
	type args struct {
		packedTx string
		parser   *ViacoinParser
	}
	tests := []struct {
		name    string
		args    args
		want    *bchain.Tx
		want1   uint32
		wantErr bool
	}{
		{
			name: "viacoin-1",
			args: args{
				packedTx: testTxPacked1,
				parser:   NewViacoinParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    &testTx1,
			want1:   5159277,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := hex.DecodeString(tt.args.packedTx)
			got, got1, err := tt.args.parser.UnpackTx(b)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpackTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unpackTx() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("unpackTx() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}