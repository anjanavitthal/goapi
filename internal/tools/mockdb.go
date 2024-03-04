package tools

import "time"

type mockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123456",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456ABC",
		Username:  "jason",
	},
	"maria": {
		AuthToken: "765ABC",
		Username:  "maria",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coin:     100,
		Username: "alex",
	},
	"jason": {
		Coin:     200,
		Username: "jason",
	},
	"maria": {
		Coin:     300,
		Username: "maria",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}

	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoinDetails(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}

	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
