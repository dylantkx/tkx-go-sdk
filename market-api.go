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

// GetMarketBuyOrders - Get all currently buy orders in Tokenize specified by market
func (api *MarketAPI) GetMarketBuyOrders(market string) ([]MarketOrder, error) {
	if market == "" {
		return nil, errors.New("Missing required parameter: market [string]")
	}

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

// GetMarketSellOrders - Get all currently sell orders in Tokenize specified by market
func (api *MarketAPI) GetMarketSellOrders(market string) ([]MarketOrder, error) {
	if market == "" {
		return nil, errors.New("Missing required parameter: market [string]")
	}

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

// GetMarketOrderBook - Allow you to keep track of the state of Tokenize order books on a price aggregated basis with customizable precision.
// @params:
// - market [string] (required): The market, eg: "MYR-BTC";
// - orderType [string] (optional): "buy" | "sell" | "both" | "";
// - limit [int]: Limit numbers of order book to return;
func (api *MarketAPI) GetMarketOrderBook(market string, orderType string, limit int) (*OrderBookData, error) {
	if market == "" {
		return nil, errors.New("Missing required parameter: market [string]")
	}

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
