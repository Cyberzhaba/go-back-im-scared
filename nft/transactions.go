package nft

import (
	"log"
	"math/big"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

// https://ethereum.stackexchange.com/questions/16472/signing-a-raw-transaction-in-go
func SignTransaction(wallet hdwallet.Wallet, account accounts.Account, data []byte) {
	nonce := uint64(0)
	value := big.NewInt(3)
	// toAddress := common.HexToAddress("0x0")
	toAddress := account.Address
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(2)

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := wallet.SignTx(account, tx, nil)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(signedTx)
}

// Return account and value from best seller
func GetBestBidOrder(itemID string) (string, int) {
	bids, err := GetNftOrderItemByID(itemID)
	if err != nil {
		log.Println(err)
	}
	return bids.Royalties[0].AccountAddr, bids.Royalties[0].Value
}
