package rest

import (
	"net/http"
)

type UserResponse struct {
	UID              string `json:"uid"`
	Balance          int    `json:"balance"`
	AccountType      string `json:"account_type"`
	Username         string `json:"username"`
	Linkingpublickey string `json:"linkingpublickey"`
}

// User - Get the user account Information.
func (api *LNMarkets) User() (*UserResponse, error) {
	response := &UserResponse{}
	if err := api.request(http.MethodGet, "user", true, nil, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
