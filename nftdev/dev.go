package nftdev

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func decode(pemEncoded string) *ecdsa.PrivateKey {
	block, err := pem.Decode([]byte(pemEncoded))
	if err != nil {
		fmt.Printf("Error: %s", err)
		// return
	}
	x509Encoded := block.Bytes
	privateKey, b := x509.ParseECPrivateKey(x509Encoded)
	if err != nil {
		fmt.Printf("Error: %s", err)
		// return
	}
	fmt.Println(b)
	return privateKey
}

func Dev() {
	type AssetType1 struct {
		AsseClass string
		Value     int
	}
	type AssetType2 struct {
		AssetClass2 string
		Contract    string
		TokenId     string
	}
	type Make struct {
		AssetType AssetType1
		Value     int
	}
	type Take struct {
		AssetType2
		Value int
	}
	type Data struct {
		DataType   string
		Payouts    interface{}
		OriginFees interface{}
	}
	type jsdata struct {
		Type  string
		Maker string
		Make  Make
		Take  Take
		Salt  int
		Data  Data
	}

	At1 := AssetType1{
		AsseClass: "ETH",
	}

	At2 := AssetType2{
		AssetClass2: "ERC721",
		Contract:    "0x3d4db08af8370896a6ceed11a9f22616765ef73e",
		TokenId:     "4",
	}

	Tk := Take{
		AssetType2: At2,
		Value:      1,
	}

	mk := Make{
		AssetType: At1,
		Value:     3000000000000000,
	}

	dt := Data{
		DataType:   "RARIBLE_V2_DATA_V1",
		Payouts:    nil,
		OriginFees: nil,
	}
	jsondata := jsdata{
		Type:  "RARIBLE_V2",
		Maker: "0xcba93eeed208b1bec89e7b3a19356ced335bc6ed",
		Make:  mk,
		Take:  Tk,
		Salt:  717802,
		Data:  dt,
	}
	b, err := json.Marshal(jsondata)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	// wallet2 := hdwallet
	hash := crypto.Keccak256Hash(b)
	// fmt.Println(hash, b)

	// tree, _ := marketree.Mew()
	// // bn :=
	// fmt.Println(bn)
	// var data []byte
	// account, wallet := nft.GenerateWallet()
	// privkey := crypto."ab7d05552d9f5a561571168d046cae24690aef606f1cfd661184a1d0ca56c0ad"
	// privKey, _ := wallet.PrivateKey(account)
	privkey := decode("ab7d05552d9f5a561571168d046cae24690aef606f1cfd661184a1d0ca56c0ad")
	sign, err := crypto.Sign(hash[:], privkey)
	fmt.Println(sign, err)
	// nft.SignTransaction(*wallet, account, data)
}
