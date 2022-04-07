package model

type User struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	ProfileLink string `json:"profile_link"`
	WalletAddress string   `json:"walletAddress"`
	Tokens        int64    `json:"tokens"`
	Badges        []string `json:"badges"`
	Description   string   `json:"description"`
}
