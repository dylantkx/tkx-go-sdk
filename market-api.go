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

// GetMarketBuyOrders - Get all currently buy orders in Tokenize specified by market
func (api *MarketAPI) GetMarketBuyOrders(market string) (*HttpResponseMarketOrders, error) {
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

	return json, nil
}

// GetMarketSellOrders - Get all currently sell orders in Tokenize specified by market
func (api *MarketAPI) GetMarketSellOrders(market string) (*HttpResponseMarketOrders, error) {
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

	return json, nil
}

// NewMarketAPI - MarketAPI constructor
func NewMarketAPI(httpManager *HttpManager) *MarketAPI {
	return &MarketAPI{
		httpManager: httpManager,
		endpoint:    httpManager.baseURL + "/market",
	}
}
