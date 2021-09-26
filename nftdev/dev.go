package nftdev

import (
	"crypto/x509"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

// type ecPrivateKey struct {
// 	Version       int
// 	PrivateKey    []byte
// 	NamedCurveOID asn1.ObjectIdentifier `asn1:"optional,explicit,tag:0"`
// 	PublicKey     asn1.BitString        `asn1:"optional,explicit,tag:1"`
// }

// func ParsePKCS1PrivateKey(der []byte) (*rsa.PrivateKey, error) {
// 	var priv pkcs1PrivateKey
// 	rest, err := asn1.Unmarshal(der, &priv)
// 	if len(rest) > 0 {
// 		return nil, asn1.SyntaxError{Msg: "trailing data"}
// 	}
// 	if err != nil {
// 		if _, err := asn1.Unmarshal(der, &ecPrivateKey{}); err == nil {
// 			return nil, errors.New("x509: failed to parse private key (use ParseECPrivateKey instead for this key format)")
// 		}
// 		if _, err := asn1.Unmarshal(der, &pkcs8{}); err == nil {
// 			return nil, errors.New("x509: failed to parse private key (use ParsePKCS8PrivateKey instead for this key format)")
// 		}
// 		return nil, err
// 	}

// 	if priv.Version > 1 {
// 		return nil, errors.New("x509: unsupported private key version")
// }

// if priv.N.Sign() <= 0 || priv.D.Sign() <= 0 || priv.P.Sign() <= 0 || priv.Q.Sign() <= 0 {
// 	return nil, errors.New("x509: private key contains zero or negative value")
// }

// key := new(rsa.PrivateKey)
// key.PublicKey = rsa.PublicKey{
// 	E: priv.E,
// 	N: priv.N,
// }

// key.D = priv.D
// key.Primes = make([]*big.Int, 2+len(priv.AdditionalPrimes))
// key.Primes[0] = priv.P
// key.Primes[1] = priv.Q
// for i, a := range priv.AdditionalPrimes {
// 	if a.Prime.Sign() <= 0 {
// 		return nil, errors.New("x509: private key contains zero or negative prime")
// 	}
// 	key.Primes[i+2] = a.Prime
// 	// We ignore the other two values because rsa will calculate
// 	// them as needed.
// }

// err = key.Validate()
// if err != nil {
// 	return nil, err
// }
// 	key.Precompute()

// 	return key, nil
// }

// func decode(pemEncoded string) *ecdsa.PrivateKey {
// 	block, _ := pem.Decode([]byte(pemEncoded))
// 	// if err != nil {
// 	// 	fmt.Printf("Error: %s", err)
// 	// 	return
// 	// }
// 	x509Encoded := block.Bytes
// 	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)
// 	// if err != nil {
// 	// 	fmt.Printf("Error: %s", err)
// 	// 	return error.N
// 	// }
// 	return privateKey
// }

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
	privkey, err := x509.ParseECPrivateKey([]byte("ab7d05552d9f5a561571168d046cae24690aef606f1cfd661184a1d0ca56c0ad"))
	sign, err := crypto.Sign(hash[:], privkey)
	fmt.Println(sign, err)
	// nft.SignTransaction(*wallet, account, data)
}
