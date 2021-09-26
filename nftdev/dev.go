package main

import (
	"encoding/json"
	"fmt"
)

func main() {
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
	jsondata := &jsdata{}
	b, err := json.Marshal(jsondata)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	fmt.Println(b)

	// tree, _ := marketree.Mew()
	// // bn :=
	// fmt.Println(bn)
	// var data []byte
	// account, wallet := nft.GenerateWallet()
	// nft.SignTransaction(*wallet, account, data)
}
