package tkxsdk

type HttpResponseAccountInfo struct {
	Message string
	Status  string
	Data    *AccountInfo
}

type AccountInfo struct {
	Email                  string
	RegId                  string
	AccountType            string
	IsMobileNumberVerified int
	RoleId                 int
	DigitalTokenTrading    bool
	FiatTrading            bool
	MarginTrading          bool
	DailyWithdrawalLimit   string
	Message                string
}

type HttpResponseAccountBalances struct {
	Message string
	Status  string
	Data    *[]AccountBalance
}

type AccountBalance struct {
	Currency  string
	Balance   float64
	Available float64
	Pending   float64
	Address   string
}

type HttpResponseMarketOrders struct {
	Message string
	Status  string
	Data    *MarketOrderData
}

type MarketOrderData struct {
	Market string
	Orders *[]MarketOrder
}

type MarketOrder struct {
	Quantity float64
	Rate     float64
	Total    float64
	Sum      float64
}

type HttpResponseOrderBook struct {
	Message string
	Status  string
	Data    *OrderBookData
}

type OrderBookData struct {
	Buy  *[]MarketOrder
	Sell *[]MarketOrder
}

type HttpResponsePlaceOrder struct {
	Message string
	Status  string
	Data    *PlacedOrder
}

type PlacedOrder struct {
	OrderID          int
	Market           string
	TransactionTyime int
	Price            float64
	OriginUnits      int
	CommissionPaid   float64
	Type             string
	Status           string
}
