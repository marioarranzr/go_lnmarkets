package rest

const (
	// Url - LNMarkets rest API Endpoint
	Url = "https://api.lnmarkets.com"
	// Version - LNMarkets API Version Number
	Version = "v1"
)

type (
	OrderSide string
	OrderType string
)

// Order sides
const (
	OrderSideBuy  OrderSide = "b"
	OrderSideSell OrderSide = "s"
)

// Order types
const (
	OrderTypeLimit  OrderType = "l"
	OrderTypeMarket OrderType = "m"
)
