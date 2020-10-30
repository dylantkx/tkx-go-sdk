package tkxsdk

import (
	"errors"
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

// CancelOrder - Cancel one order
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#cancel_buy_order
func (api *OrderAPI) CancelOrder(id int) (bool, error) {
	resp, err := req.Delete(api.endpoint+"/"+strconv.Itoa(id), api.httpManager.header)
	if err != nil {
		return false, err
	}

	json := &HttpResponseCancelOrder{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return false, parsingError
	}

	if json.Status != "success" {
		return false, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetMyOrders - Get all orders of this account
func (api *OrderAPI) GetMyOrders(market string) ([]MyOrder, error) {
	// TODO: Find a way to execute these two requests concurrently
	b, err1 := api.GetMyBuyOrders(market)
	if err1 != nil {
		return nil, err1
	}
	s, err2 := api.GetMySellOrders(market)
	if err2 != nil {
		return nil, err2
	}
	return append(b, s...), nil
}

// GetMyPendingBuyOrders - Gets all PENDING buy orders of this account
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#current_open_your_buy
func (api *OrderAPI) GetMyPendingBuyOrders(market string) ([]MyOrder, error) {
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

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetMyPendingSellOrders - Gets all PENDING sell orders of this account
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#current_open_your_sell
func (api *OrderAPI) GetMyPendingSellOrders(market string) ([]MyOrder, error) {
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

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetMyBuyOrders - Gets all buy orders of this account
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_all_your_buy_orders
func (api *OrderAPI) GetMyBuyOrders(market string) ([]MyOrder, error) {
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

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetBuyOrderDetail - Check your buy order by specified id.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#check_buy_order
func (api *OrderAPI) GetBuyOrderDetail(id int) (*OrderDetail, error) {
	queryParams := req.QueryParam{
		"orderId": id,
	}
	resp, err := req.Get(api.endpoint+"/buy", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseGetOrderDetail{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetMySellOrders - Gets all sell orders of this account
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_all_your_sell_orders
func (api *OrderAPI) GetMySellOrders(market string) ([]MyOrder, error) {
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

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetSellOrderDetail - Check your sell order by specified id.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#check_sell_order
func (api *OrderAPI) GetSellOrderDetail(id int) (*OrderDetail, error) {
	queryParams := req.QueryParam{
		"orderId": id,
	}
	resp, err := req.Get(api.endpoint+"/sell", api.httpManager.header, queryParams)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseGetOrderDetail{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// PlaceLimitBuyOrder - Place a new limit buy order
// @params:
// - market [string] (required): The market, eg: "MYR-BTC";
// - units [float64] (required): The amount of this order;
// - price [float64] (required): The price of this order.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#place_buy_order
func (api *OrderAPI) PlaceLimitBuyOrder(market string, units, price float64) (*PlacedOrder, error) {
	return api.placeBuyOrder(market, "limit", units, price)
}

// PlaceMarketBuyOrder - Place a new limit buy order
// @params:
// - market [string] (required): The market, eg: "MYR-BTC";
// - units [float64] (required): The amount of this order.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#place_buy_order
func (api *OrderAPI) PlaceMarketBuyOrder(market string, units float64) (*PlacedOrder, error) {
	return api.placeBuyOrder(market, "market", units, 0)
}

// PlaceLimitSellOrder - Place a new limit sell order
// @params:
// - market [string] (required): The market, eg: "MYR-BTC";
// - units [float64] (required): The amount of this order;
// - price [float64] (required): The price of this order.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#place_sell_order
func (api *OrderAPI) PlaceLimitSellOrder(market string, units, price float64) (*PlacedOrder, error) {
	return api.placeSellOrder(market, "limit", units, price)
}

// PlaceMarketSellOrder - Place a new limit sell order
// @params:
// - market [string] (required): The market, eg: "MYR-BTC";
// - units [float64] (required): The amount of this order.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#place_sell_order
func (api *OrderAPI) PlaceMarketSellOrder(market string, units float64) (*PlacedOrder, error) {
	return api.placeSellOrder(market, "market", units, 0)
}

func (api *OrderAPI) placeBuyOrder(market, orderType string, units, price float64) (*PlacedOrder, error) {
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

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

func (api *OrderAPI) placeSellOrder(market, orderType string, units, price float64) (*PlacedOrder, error) {
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

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}
