package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type AddMarginRequest struct {
	Amount int64  `json:"amount"`
	Pid    string `json:"pid"`
}
type AddMarginResponse struct {
	Pid            string `json:"pid"`
	ID             int    `json:"id"`
	Type           string `json:"type"`
	TakeprofitWi   string `json:"takeprofit_wi"`
	Takeprofit     int    `json:"takeprofit"`
	StoplossWi     string `json:"stoploss_wi"`
	Stoploss       int    `json:"stoploss"`
	Side           string `json:"side"`
	Quantity       int    `json:"quantity"`
	Price          int    `json:"price"`
	Pl             int    `json:"pl"`
	MarketWi       string `json:"market_wi"`
	MarketFilledTs string `json:"market_filled_ts"`
	MarginWi       string `json:"margin_wi"`
	Margin         int    `json:"margin"`
	Liquidation    int    `json:"liquidation"`
	Leverage       int    `json:"leverage"`
	CreationTs     string `json:"creation_ts"`
}

type CancelAllResponse struct {
	Data []struct {
		Pid       string `json:"pid"`
		ExitPrice int    `json:"exit_price"`
		ClosedTs  string `json:"closed_ts"`
		Closed    bool   `json:"closed"`
		Pl        int    `json:"pl"`
	} `json:"data"`
}

type CloseAllResponse CancelAllResponse

type CancelResponse struct {
	Pid string `json:"pid"`
}

type CarryFeesHistoryResponse struct {
	Data []struct {
		Fixing string `json:"fixing"`
		Index  int    `json:"index"`
	} `json:"data"`
}

type CashInResponse struct {
	Amount int    `json:"amount"`
	Pid    string `json:"pid"`
}

type CloseResponse struct {
	Pid       string `json:"pid"`
	ExitPrice int    `json:"exit_price"`
	ClosedTs  string `json:"closed_ts"`
	Closed    bool   `json:"closed"`
	Pl        int    `json:"pl"`
}

type PositionsResponse struct {
	Positions []struct {
		Pid            string `json:"pid"`
		ID             int    `json:"id"`
		Type           string `json:"type"`
		TakeprofitWi   string `json:"takeprofit_wi"`
		Takeprofit     int    `json:"takeprofit"`
		StoplossWi     string `json:"stoploss_wi"`
		Stoploss       int    `json:"stoploss"`
		Sign           int    `json:"sign"`
		Side           string `json:"side"`
		Quantity       int    `json:"quantity"`
		Price          int    `json:"price"`
		Pl             int    `json:"pl"`
		MarketWi       string `json:"market_wi"`
		MarketFilledTs int64  `json:"market_filled_ts"`
		MarginWi       string `json:"margin_wi"`
		Margin         int    `json:"margin"`
		Liquidation    int    `json:"liquidation"`
		Leverage       int    `json:"leverage"`
		ExitPrice      int    `json:"exit_price"`
		CreationTs     int64  `json:"creation_ts"`
		ClosedTs       int64  `json:"closed_ts"`
		Closed         bool   `json:"closed"`
		Canceled       bool   `json:"canceled"`
		SumCarryFees   int    `json:"sum_carry_fees"`
	} `json:"positions"`
}

type NewPositionRequest struct {
	Type       OrderType `json:"type"`
	Side       OrderSide `json:"side"`
	Price      float64   `json:"price"`
	Margin     int       `json:"margin"`
	Stoploss   float64   `json:"stoploss"`
	Takeprofit float64   `json:"takeprofit"`
	Quantity   float64   `json:"quantity"`
	Leverage   float64   `json:"leverage"`
}
type NewPositionResponse struct {
	Pid            string  `json:"pid"`
	ID             int     `json:"id"`
	Type           string  `json:"type"`
	TakeprofitWi   string  `json:"takeprofit_wi"`
	Takeprofit     float64 `json:"takeprofit"`
	StoplossWi     string  `json:"stoploss_wi"`
	Stoploss       float64 `json:"stoploss"`
	Side           string  `json:"side"`
	Quantity       float64 `json:"quantity"`
	Price          float64 `json:"price"`
	Pl             float64 `json:"pl"`
	MarketWi       string  `json:"market_wi"`
	MarketFilledTs string  `json:"market_filled_ts"`
	MarginWi       string  `json:"margin_wi"`
	Margin         int     `json:"margin"`
	Liquidation    float64 `json:"liquidation"`
	Leverage       float64 `json:"leverage"`
	CreationTs     string  `json:"creation_ts"`
}
type UpdateResponse struct {
	Pid   string  `json:"pid"`
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

type BidAndOfferHistoryResponse struct {
	Data []struct {
		Time  int64 `json:"time"`
		Bid   int   `json:"bid"`
		Offer int   `json:"offer"`
	} `json:"data"`
}

type FixingHistoryResponse struct {
	Data []struct {
		Ts              int64   `json:"ts"`
		ID              string  `json:"id"`
		FixingPrice     int     `json:"fixing_price"`
		FeePercentValue float64 `json:"fee_percent_value"`
	} `json:"data"`
}

type IndexHistoryResponse struct {
	Data []struct {
		Time  int64 `json:"time"`
		Index int   `json:"index"`
	} `json:"data"`
}

type InstrumentResponse struct {
	MaxPositionsCount int     `json:"max_positions_count"`
	MaxMargin         int     `json:"max_margin"`
	MaxLeverage       int     `json:"max_leverage"`
	CarryFees         float64 `json:"carry_fees"`
}

type TickerResponse struct {
	Bid   float64 `json:"bid"`
	Offer float64 `json:"offer"`
	Index float64 `json:"index"`
}

// AddMargin - Add margin to a running position.
func (api *LNMarkets) AddMargin(amount int64, pid string) (*AddMarginResponse, error) {
	response := &AddMarginResponse{}
	request := AddMarginRequest{
		Amount: amount,
		Pid:    pid,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	if err := api.request(http.MethodPost, "futures/add-margin", true, nil, body, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Cancel - Cancel the position linked to the given pid.
// Only works on positions that are not currently filled.
func (api *LNMarkets) Cancel(pid string) (*CancelResponse, error) {
	response := &CancelResponse{}
	if err := api.request(http.MethodPost, "futures/cancel", true, nil, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// CancelAll - Cancel all open positions.
func (api *LNMarkets) CancelAll() (*CancelAllResponse, error) {
	response := &CancelAllResponse{}
	if err := api.request(http.MethodDelete, "futures/all/cancel", true, nil, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Close - Close the user position.
// The PL will be calculated against the current bid or offer depending on the side of the position.
func (api *LNMarkets) Close() (*CloseResponse, error) {
	response := &CloseResponse{}
	if err := api.request(http.MethodDelete, "futures", true, nil, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// CloseAll - Close every user position.
// The PL will be calculated against the current bid or offer depending on the side of the position.
func (api *LNMarkets) CloseAll() (*CloseAllResponse, error) {
	response := &CloseAllResponse{}
	if err := api.request(http.MethodDelete, "futures/all/close", true, nil, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// CarryFeesHistory - Retrieves carry fees for user.
func (api *LNMarkets) CarryFeesHistory(from int64, to int64, limit int64) (*CarryFeesHistoryResponse, error) {
	response := &CarryFeesHistoryResponse{}
	params := make(url.Values)
	params.Add("from", fmt.Sprint(from))
	params.Add("to", fmt.Sprint(to))
	params.Add("limit", fmt.Sprint(limit))

	if err := api.request(http.MethodGet, "futures/carry-fees", true, params, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// CashIn - Retrieves part of one running positions PL.
func (api *LNMarkets) CashIn(pid string) (*CashInResponse, error) {
	response := &CashInResponse{}
	if err := api.request(http.MethodPost, "futures/cash-in", true, nil, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Positions - Fetch users positions.
func (api *LNMarkets) Positions() (*PositionsResponse, error) {
	response := &PositionsResponse{}
	if err := api.request(http.MethodGet, "futures", true, nil, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// NewPosition - Send the order form parameters to add a new position in database.
// If type="l", the property price must be included in the request to know when the position should be filled.
// You can choose to use the margin or the quantity as a parameter, the other will be calculated with the one you chose.
func (api *LNMarkets) NewPosition(t OrderType, s OrderSide, p float64, m int, sl, tp, q, l float64) (*NewPositionResponse, error) {
	response := &NewPositionResponse{}
	request := NewPositionRequest{
		Type:       t,
		Side:       s,
		Price:      p,
		Margin:     m,
		Stoploss:   sl,
		Takeprofit: tp,
		Quantity:   q,
		Leverage:   l,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	if err := api.request(http.MethodPost, "futures", true, nil, body, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// ------
// ------
// ------
// ------
// ------
// ------
// ------

// Ticker - Retrieves the futures ticker.
func (api *LNMarkets) Ticker() (*TickerResponse, error) {
	response := &TickerResponse{}
	if err := api.request(http.MethodGet, "futures/ticker", false, nil, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
