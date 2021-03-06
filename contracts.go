package tkxsdk

type HttpResponseGetMarkets struct {
	Message string
	Status  string
	Data    []MarketInfo
}

type MarketInfo struct {
	Title string
	List  []MarketDetail
}

type MarketDetail struct {
	MarketID           string
	MarketName         string
	BaseCurrency       string
	MarketCurrency     string
	MarketCurrencyLong string
	Status             string
	Ceiling            string
	Floor              string
}

type HttpResponseGetLastMarketPrice struct {
	Message string
	Status  string
	Data    *MarketPrice
}

type MarketPrice struct {
	Market     string
	BaseAsset  string
	QuoteAsset string
	LastPrice  string
	PrevPrice  string
}

type HttpResponseGetMarketTicker struct {
	Message string
	Status  string
	Data    *MarketTicker
}

type MarketTicker struct {
	Market     string
	BaseAsset  string
	QuoteAsset string
	LastPrice  string
	BidPrice   float64
	BidQty     float64
	AskPrice   float64
	AskQty     float64
}

type HttpResponseGetMarketSummaries struct {
	Message string
	Status  string
	Data    []MarketSummary
}

type MarketSummary struct {
	MarketID  int
	Market    string
	Status    string
	AskPrice  float64
	BidPrice  float64
	LastPrice float64
	OpenPrice float64
	PrevPrice float64
	High      float64
	Low       float64
	Volume    float64
}

type HttpResponseGetMarketAvgPrice struct {
	Message string
	Status  string
	Data    *MarketAvgPrice
}

type MarketAvgPrice struct {
	Market   string
	AvgPrice float64
}

type HttpResponseGetMarketTicker24h struct {
	Message string
	Status  string
	Data    *MarketTicker24h
}

type MarketTicker24h struct {
	Market    string
	HighPrice float64
	LowPrice  float64
	LastPrice float64
	BidPrice  float64
	AskPrice  float64
	Volume    float64
}

type HttpResponseGetMarketHistory struct {
	Message string
	Status  string
	Data    []MarketHistoryRecord
}

type MarketHistoryRecord struct {
	Timestamp int
	Type      string
	Amount    float64
	Price     float64
	Total     float64
}

type HttpResponseGetCandles struct {
	Message string
	Status  string
	Data    []MarketCandleRaw
}

type MarketCandleRaw = [6]interface{}

type MarketCandle struct {
	OpenTime int
	Open     float64
	Close    float64
	High     float64
	Low      float64
	Volume   float64
}

type HttpResponseGetOrderDetail struct {
	Message string
	Status  string
	Data    *OrderDetail
}

type OrderDetail struct {
	OrderID         int
	Side            string
	Market          string
	Status          string
	Type            string
	TransactionTime int
	OriginUnits     float64
	Price           float64
	FilledUnits     float64
	CommissionPaid  float64
	Fills           []OrderFill
}

type OrderFill struct {
	Price           float64
	Units           float64
	Commission      float64
	CommissionAsset string
}

type HttpResponseMarketOrders struct {
	Message string
	Status  string
	Data    *MarketOrderData
}

type MarketOrderData struct {
	Market string
	Orders []MarketOrder
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
	Buy  []MarketOrder
	Sell []MarketOrder
}

type HttpResponsePlaceOrder struct {
	Message string
	Status  string
	Data    *PlacedOrder
}

type PlacedOrder struct {
	OrderID         int
	Market          string
	TransactionTime int
	Price           float64
	OriginUnits     float64
	CommissionPaid  float64
	Type            string
	Status          string
}

type HttpResponseMyOrders struct {
	Message string
	Status  string
	Data    []MyOrder
}

type MyOrder struct {
	OrderID        int
	Market         string
	Price          float64
	OriginUnits    float64
	FilledUnits    float64
	CommissionPaid float64
	Type           string
	Status         string
}

type HttpResponseCancelOrder struct {
	Message string
	Status  string
	Data    bool
}

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
	Data    []AccountBalance
}

type AccountBalance struct {
	Currency  string
	Balance   float64
	Available float64
	Pending   float64
	Address   string
}

type HttpResponseAccountBalance struct {
	Message string
	Status  string
	Data    *AccountBalance
}

type HttpResponseGetDepositAddress struct {
	Message string
	Status  string
	Data    *DepositAddress
}

type DepositAddress struct {
	Currency string
	Address  string
}
