package tools

import (
	"encoding/json"
	"fmt"
)

type DataOrder struct {
	DataType string
	Fee      int
}

type Order struct {
	Type string
	Data DataOrder
}

type Orders struct {
	Orders       []Order
	Continuation string
}

type Bids struct {
	Item
	BestSellOrder Order
	BestBidOrder  Order
	TotalStock    int
	Sellers       int
}

// Return Order
func GetSellOrdersByItem(contract, tokenID string) Order {
	body := GetBodyFromUrl(fmt.Sprintf(
		"https://ethereum-api-staging.rarible.org/v0.1/order/orders/sell/byItem?contract=%s&tokenId%s",
		contract,
		tokenID,
	),
	)

	var data Order
	json.Unmarshal(body, &data)
	return data
}

// Return Bids
func GetNftOrderItemByID(itemID string) Bids {
	body := GetBodyFromUrl(fmt.Sprintf(
		"https://ethereum-api-staging.rarible.org/v0.1/nft-order/items/%s",
		itemID,
	),
	)

	var data Bids
	json.Unmarshal(body, &data)
	return data
}

// Return Orders
func GetOrderBidsByItem(contract, tokenID string) Orders {
	body := GetBodyFromUrl(fmt.Sprintf(
		"https://ethereum-api-staging.rarible.org/v0.1/order/orders/bids/byItem?contract=%s&tokenId%s",
		contract,
		tokenID,
	),
	)

	var data Orders
	json.Unmarshal(body, &data)
	return data
}
