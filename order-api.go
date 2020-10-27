package tkxsdk

import (
	"github.com/imroc/req"
)

// OrderAPI - OrderAPI struct
type OrderAPI struct {
	httpManager *HttpManager
	endpoint    string
}

// NewOrderAPI - OrderAPI constructor
func NewOrderAPI(httpManager *HttpManager) *OrderAPI {
	return &OrderAPI{
		httpManager: httpManager,
		endpoint:    httpManager.baseURL + "/order",
	}
}

// PlaceLimitBuyOrder - Place a new limit buy order
// @params:
// - market [string] (required): The market, eg: "MYR-BTC";
// - units [float64] (required): The amount of this order;
// - price [float64] (required): The price of this order;
func (api *OrderAPI) PlaceLimitBuyOrder(market string, units, price float64) (*HttpResponsePlaceOrder, error) {
	return api.placeBuyOrder(market, "limit", units, price)
}

// PlaceMarketBuyOrder - Place a new limit buy order
// @params:
// - market [string] (required): The market, eg: "MYR-BTC";
// - units [float64] (required): The amount of this order;
func (api *OrderAPI) PlaceMarketBuyOrder(market string, units float64) (*HttpResponsePlaceOrder, error) {
	return api.placeBuyOrder(market, "market", units, 0)
}

// PlaceLimitSellOrder - Place a new limit sell order
// @params:
// - market [string] (required): The market, eg: "MYR-BTC";
// - units [float64] (required): The amount of this order;
// - price [float64] (required): The price of this order;
func (api *OrderAPI) PlaceLimitSellOrder(market string, units, price float64) (*HttpResponsePlaceOrder, error) {
	return api.placeSellOrder(market, "limit", units, price)
}

// PlaceMarketSellOrder - Place a new limit sell order
// @params:
// - market [string] (required): The market, eg: "MYR-BTC";
// - units [float64] (required): The amount of this order;
func (api *OrderAPI) PlaceMarketSellOrder(market string, units float64) (*HttpResponsePlaceOrder, error) {
	return api.placeSellOrder(market, "market", units, 0)
}

func (api *OrderAPI) placeBuyOrder(market, orderType string, units, price float64) (*HttpResponsePlaceOrder, error) {
	bodyMap := map[string]interface{}{
		"market":    market,
		"orderType": orderType,
		"units":     units,
	}
	if orderType == "limit" {
		bodyMap["price"] = price
	}
	jsonBody := req.BodyJSON(bodyMap)
	resp, err := req.Post(api.endpoint+"/buy", api.httpManager.header, jsonBody)
	if err != nil {
		return nil, err
	}

	json := &HttpResponsePlaceOrder{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	return json, nil
}

func (api *OrderAPI) placeSellOrder(market, orderType string, units, price float64) (*HttpResponsePlaceOrder, error) {
	bodyMap := map[string]interface{}{
		"market":    market,
		"orderType": orderType,
		"units":     units,
	}
	if orderType == "limit" {
		bodyMap["price"] = price
	}
	jsonBody := req.BodyJSON(bodyMap)
	resp, err := req.Post(api.endpoint+"/sell", api.httpManager.header, jsonBody)
	if err != nil {
		return nil, err
	}

	json := &HttpResponsePlaceOrder{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	return json, nil
}
