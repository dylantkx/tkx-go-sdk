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

// GetAccountInfo - Use to get your account information
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

// GetAccountBalances - Use to retrieve all balances from your account
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
