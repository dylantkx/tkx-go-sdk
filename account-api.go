package tkxsdk

import (
	"errors"
	"github.com/imroc/req"
)

// AccountAPI - AccountAPI struct
type AccountAPI struct {
	httpManager *HttpManager
	endpoint    string
}

// NewAccountAPI - AccountAPI constructor
func NewAccountAPI(httpManager *HttpManager) *AccountAPI {
	return &AccountAPI{
		httpManager: httpManager,
		endpoint:    httpManager.baseURL + "/account",
	}
}

// GetAccountInfo - Use to get your account information.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_account_information
func (api *AccountAPI) GetAccountInfo() (*AccountInfo, error) {
	resp, err := req.Get(api.endpoint, api.httpManager.header)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseAccountInfo{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetAccountBalances - Use to retrieve all balances from your account.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_account_balances
func (api *AccountAPI) GetAccountBalances() ([]AccountBalance, error) {
	resp, err := req.Get(api.endpoint+"/balances", api.httpManager.header)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseAccountBalances{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetAccountBalanceByCCY - Use to retrieve the balance from your account for a specified currency.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_currency_balances
func (api *AccountAPI) GetAccountBalanceByCCY(currency string) (*AccountBalance, error) {
	queryParam := req.QueryParam{
		"currency": currency,
	}
	resp, err := req.Get(api.endpoint+"/balance", api.httpManager.header, queryParam)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseAccountBalance{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}

// GetDepositAddress - Use to retrieve or generate an address for a specified currency.
// Reference: https://tokenizexchange.zendesk.com/hc/en-gb/articles/360022521593-Developer-s-Guide-API#get_deposit_address
func (api *AccountAPI) GetDepositAddress(currency string) (*DepositAddress, error) {
	queryParam := req.QueryParam{
		"currency": currency,
	}
	resp, err := req.Get(api.endpoint+"/deposit-address", api.httpManager.header, queryParam)
	if err != nil {
		return nil, err
	}

	json := &HttpResponseGetDepositAddress{}
	parsingError := resp.ToJSON(&json)
	if parsingError != nil {
		return nil, parsingError
	}

	if json.Status != "success" {
		return nil, errors.New(json.Message)
	}

	return json.Data, nil
}
