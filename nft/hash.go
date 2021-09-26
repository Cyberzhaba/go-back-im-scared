package nft

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/ethersphere/bee/pkg/crypto/eip712"
)

func GetHash() []byte {
	chainid := *math.NewHexOrDecimal256(1000000000) //TODO
	pretype := core.Type{
		Name: "name", // TODO
		Type: "type", // TODO
	}
	typeDataDomain := core.TypedDataDomain{
		Name:              "name",    //TODO
		Version:           "version", //TODO
		ChainId:           &chainid,
		VerifyingContract: "contract", //TODO
		Salt:              "salt",     //TODO
	}
	types := core.Types{"types1": []core.Type{pretype}} //TODO
	typeDataMessage := core.TypedDataMessage{}          //TODO
	// typeDataMessage := make(map[string]interface{})
	// typeDataMessage["1"] = void
	typedata := core.TypedData{
		Types:       types,
		PrimaryType: "prtype", //TODO
		Domain:      typeDataDomain,
		Message:     typeDataMessage,
	}
	hash, err := eip712.EncodeForSigning(&typedata)
	if err != nil {
		fmt.Printf("%v", err)
	}

	return hash

}
