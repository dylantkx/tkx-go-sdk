package tkxsdk

import (
	"github.com/imroc/req"
	"strconv"
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

func (api *OrderAPI) CancelOrder(id int) (*HttpResponseCancelOrder, error) {
	resp, err := req.Delete(api.endpoint+"/"+strconv.Itoa(id), api.httpManager.header)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseCancelOrder{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	return json, nil
}

func (api *OrderAPI) GetMyOrders(market string) (*HttpResponseMyOrders, error) {
	// TODO: Find a way to execute these two requests concurrently
	b, err1 := api.GetMyBuyOrders(market)
	if err1 != nil {
		return nil, err1
	}
	s, err2 := api.GetMySellOrders(market)
	if err2 != nil {
		return nil, err2
	}
	ret := &HttpResponseMyOrders{
		Status:  "success",
		Message: "",
		Data:    append(b.Data, s.Data...),
	}
	return ret, nil
}

func (api *OrderAPI) GetMyPendingBuyOrders(market string) (*HttpResponseMyOrders, error) {
	queryParams := req.QueryParam{
		"market": market,
	}
	resp, err := req.Get(api.endpoint+"/buy/open", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseMyOrders{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	return json, nil
}

func (api *OrderAPI) GetMyPendingSellOrders(market string) (*HttpResponseMyOrders, error) {
	queryParams := req.QueryParam{
		"market": market,
	}
	resp, err := req.Get(api.endpoint+"/sell/open", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseMyOrders{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	return json, nil
}

func (api *OrderAPI) GetMyBuyOrders(market string) (*HttpResponseMyOrders, error) {
	queryParams := req.QueryParam{
		"market": market,
	}
	resp, err := req.Get(api.endpoint+"/buy/all", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseMyOrders{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	return json, nil
}

func (api *OrderAPI) GetMySellOrders(market string) (*HttpResponseMyOrders, error) {
	queryParams := req.QueryParam{
		"market": market,
	}
	resp, err := req.Get(api.endpoint+"/sell/all", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseMyOrders{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	return json, nil
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
