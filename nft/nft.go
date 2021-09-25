package tools

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	AccountAddr string
	Value       int
}

type Pending struct {
	Type string
	From string
}

type Attribute struct {
	Key   string
	Value string
}

type Meta struct {
	Name        string
	Description string
	Attributes  []Attribute
}

type Item struct {
	ID         string
	Contract   string
	TokenID    string
	Creators   []Account
	Supply     int
	LazySupply int
	Owners     []string
	Royalties  []Account
	Date       string
	Pending    []Pending
	Deleted    bool
	Meta
}

type AllItems struct {
	Total         int
	Continuations string
	Items         []Item
}

func GetAllItems() AllItems {
	body := GetBodyFromUrl("https://ethereum-api-staging.rarible.org/v0.1/nft/items/all")
	var data AllItems
	json.Unmarshal(body, &data)
	return data
}

func GetItemByID(itemID string) Item {
	body := GetBodyFromUrl(fmt.Sprintf("https://ethereum-api-staging.rarible.org/v0.1/nft/items/%s", itemID))
	var data Item
	json.Unmarshal(body, &data)
	return data
}
