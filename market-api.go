package tkxsdk

import (
	"errors"
	"github.com/imroc/req"
)

// MarketAPI - MarketAPI struct
type MarketAPI struct {
	httpManager *HttpManager
	endpoint    string
}

// NewMarketAPI - MarketAPI constructor
func NewMarketAPI(httpManager *HttpManager) *MarketAPI {
	return &MarketAPI{
		httpManager: httpManager,
		endpoint:    httpManager.baseURL + "/market",
	}
}

// GetMarkets - Used to get the open and available trading markets at Tokenize along with other meta data.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#market_information
func (api *MarketAPI) GetMarkets() ([]MarketInfo, error) {
	resp, err := req.Get(api.endpoint+"/get-markets", api.httpManager.header)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseGetMarkets{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetLastMarketPrice - Get last price of the given market.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_last_market_price
func (api *MarketAPI) GetLastMarketPrice(market string) (*MarketPrice, error) {
	queryParams := req.QueryParam{
		"market": market,
	}

	resp, err := req.Get(api.endpoint+"/get-last-market-price", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseGetLastMarketPrice{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetMarketTicker - Get ticker for specified market.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_ticker_market
func (api *MarketAPI) GetMarketTicker(market string) (*MarketTicker, error) {
	queryParams := req.QueryParam{
		"market": market,
	}

	resp, err := req.Get(api.endpoint+"/ticker", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseGetMarketTicker{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetMarketSummaries - Used to get the last 24 hour summary of all active markets.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_market_summaries
func (api *MarketAPI) GetMarketSummaries() ([]MarketSummary, error) {
	resp, err := req.Get(api.endpoint+"/get-summaries", api.httpManager.header)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseGetMarketSummaries{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetMarketBuyOrders - Get all currently buy orders in Tokenize specified by market.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_market_buy_orders
func (api *MarketAPI) GetMarketBuyOrders(market string) ([]MarketOrder, error) {
	queryParams := req.QueryParam{
		"market": market,
	}
	resp, err := req.Get(api.endpoint+"/get-buy-orders", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseMarketOrders{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data.Orders, nil
}

// GetMarketSellOrders - Get all currently sell orders in Tokenize specified by market.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_market_sell_orders
func (api *MarketAPI) GetMarketSellOrders(market string) ([]MarketOrder, error) {
	queryParams := req.QueryParam{
		"market": market,
	}
	resp, err := req.Get(api.endpoint+"/get-sell-orders", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseMarketOrders{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data.Orders, nil
}

// GetMarketCurrentAveragePrice - Get the current average price for specified market.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_current_average_price
func (api *MarketAPI) GetMarketCurrentAveragePrice(market string) (*MarketAvgPrice, error) {
	queryParams := req.QueryParam{
		"market": market,
	}
	resp, err := req.Get(api.endpoint+"/get-market-avg-price", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseGetMarketAvgPrice{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetMarketTicker24h - Get statistics change price of a market.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_24h_change_price
func (api *MarketAPI) GetMarketTicker24h(market string) (*MarketTicker24h, error) {
	queryParams := req.QueryParam{
		"market": market,
	}

	resp, err := req.Get(api.endpoint+"/ticker/24h", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseGetMarketTicker24h{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetMarketOrderBook - Allow you to keep track of the state of Tokenize order books on a price aggregated basis with customizable precision.
// @params:
// - market [string] (required): The market, eg: "MYR-BTC";
// - orderType [string] (optional): "buy" | "sell" | "both" | "";
// - limit [int]: Limit numbers of order book to return.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_order_book
func (api *MarketAPI) GetMarketOrderBook(market string, orderType string, limit int) (*OrderBookData, error) {
	queryParams := req.QueryParam{
		"market": market,
		"type":   orderType,
		"limit":  limit,
	}
	resp, err := req.Get(api.endpoint+"/orderbook", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseOrderBook{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetTradeHistory - Allows the retrieval of past public trades and includes details such as price, size, and time.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_trades_histories
func (api *MarketAPI) GetTradeHistory(market string, limit int) ([]MarketHistoryRecord, error) {
	queryParams := req.QueryParam{
		"market": market,
	}
	if limit != 0 {
		queryParams["limit"] = limit
	}

	resp, err := req.Get(api.endpoint+"/history", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseGetMarketHistory{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetCandles - The Candles endpoint provides chart candles for a specified time frame and period.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_candlestick
func (api *MarketAPI) GetCandles(market string, resolution string, from, to int) ([]MarketCandle, error) {
	queryParams := req.QueryParam{
		"market":     market,
		"resolution": resolution,
	}

	if from != -1 {
		queryParams["from"] = from
	}
	if to != -1 {
		queryParams["to"] = to
	}

	resp, err := req.Get(api.endpoint+"/candles", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseGetCandles{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	candles := []MarketCandle{}
	for _, raw := range json.Data {
		c := MarketCandle{
			OpenTime: int(raw[0].(float64)),
			Open:     raw[1].(float64),
			Close:    raw[2].(float64),
			High:     raw[3].(float64),
			Low:      raw[4].(float64),
			Volume:   raw[5].(float64),
		}
		candles = append(candles, c)
	}

	return candles, nil
}

// GetRawCandles - Get raw candle data returned from API.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_candlestick
func (api *MarketAPI) GetRawCandles(market string, resolution string, from, to int) ([]MarketCandleRaw, error) {
	queryParams := req.QueryParam{
		"market":     market,
		"resolution": resolution,
	}

	if from != -1 {
		queryParams["from"] = from
	}
	if to != -1 {
		queryParams["to"] = to
	}

	resp, err := req.Get(api.endpoint+"/candles", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseGetCandles{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	for i := range json.Data {
		json.Data[i][0] = int(json.Data[i][0].(float64)) // convert OpenTime from float64 to int
	}

	return json.Data, nil
}
