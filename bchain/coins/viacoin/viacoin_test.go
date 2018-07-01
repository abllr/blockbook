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
			args:    args{address: "via1q95qlu98cpj23xy6w9tdnfn65n5vkpkey99g6wl"},
			want:    "00142d01fe14f80c9513134e2adb34cf549d1960db24",
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

	testTxPacked1 = "01000000019bfd1b63802e889498376585d14cfa59709087563003da74aa95b8e73a019ad9010000006b483045022100b4714527d32f61e9a1d9a8972d8d7794091bbb345e128f057ab97d2aabebdcf8022077f9d4d72ac88a23551f39963b677411af6a92909d4ed3b0d547bb3c908b18800121034f179b66e382ba7f8a8da57a659b29ad9ea265aee422328e8d7f2cddb760ac30ffffffff020065cd1d000000001976a914a44f088a1312a78f9063ae1fa1358d0c94074b6c88ac8c464f76aa0000001976a914d6b38173430b9cb6ff141ce8990c31af43b5f3b988ac00000000"
)

func init() {
	var (
		addr1, addr2 bchain.Address
		err          error
	)
	addr1, err = bchain.NewBaseAddress("VpycQWDnM2FpyxK5rRRhpBphgZWj79mfsK")
	if err == nil {
		addr2, err = bchain.NewBaseAddress("Vua4Z2xHUaqGkQGDjKBMtNTnbMuAgNBPaW")
	}
	if err != nil {
		panic(err)
	}

	testTx1 = bchain.Tx{
		Hex:       "01000000019bfd1b63802e889498376585d14cfa59709087563003da74aa95b8e73a019ad9010000006b483045022100b4714527d32f61e9a1d9a8972d8d7794091bbb345e128f057ab97d2aabebdcf8022077f9d4d72ac88a23551f39963b677411af6a92909d4ed3b0d547bb3c908b18800121034f179b66e382ba7f8a8da57a659b29ad9ea265aee422328e8d7f2cddb760ac30ffffffff020065cd1d000000001976a914a44f088a1312a78f9063ae1fa1358d0c94074b6c88ac8c464f76aa0000001976a914d6b38173430b9cb6ff141ce8990c31af43b5f3b988ac00000000",
		Blocktime: 1520883188,
		Txid:      "67c19fa7f6e1e4e832e85b0329ef979c02456025f3c99f7192f5625dbae18196",
		LockTime:  0,
		Vin: []bchain.Vin{
			{
				ScriptSig: bchain.ScriptSig{
					Hex: "483045022100b4714527d32f61e9a1d9a8972d8d7794091bbb345e128f057ab97d2aabebdcf8022077f9d4d72ac88a23551f39963b677411af6a92909d4ed3b0d547bb3c908b18800121034f179b66e382ba7f8a8da57a659b29ad9ea265aee422328e8d7f2cddb760ac30",
				},
				Txid:     "d99a013ae7b895aa74da03305687907059fa4cd18565379894882e80631bfd9b",
				Vout:     1,
				Sequence: 4294967295,
			},
		},
		Vout: []bchain.Vout{
			{
				Value: 5.00000000,
				N:     0,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a914a44f088a1312a78f9063ae1fa1358d0c94074b6c88ac",
					Addresses: []string{
						"VpycQWDnM2FpyxK5rRRhpBphgZWj79mfsK",
					},
				},
				Address: addr1,
			},
			{
				Value: 7321.29347212,
				N:     1,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a914d6b38173430b9cb6ff141ce8990c31af43b5f3b988ac",
					Addresses: []string{
						"Vua4Z2xHUaqGkQGDjKBMtNTnbMuAgNBPaW",
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
				height:    4766769,
				blockTime: 1520883188,
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
			want1:   4766769,
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
