package tkxsdk

// APIClient - APIClient struct
type APIClient struct {
	httpManager *HttpManager
	AccountAPI  *AccountAPI
	MarketAPI   *MarketAPI
	OrderAPI    *OrderAPI
}

// NewAPIClient - APIClient constructor
func NewAPIClient(apiBaseURL string) *APIClient {
	c := &APIClient{}
	c.httpManager = NewHttpManager(apiBaseURL)
	c.InitSubAPIs()
	return c
}

// NewAPIClientWithCredentials - APIClient factory (create with params)
func NewAPIClientWithCredentials(apiBaseURL, clientID, clientSecret string) *APIClient {
	client := NewAPIClient(apiBaseURL)
	client.httpManager.SetCredentials(clientID, clientSecret)
	return client
}

// InitSubAPIs - Initialize sub APIs
func (c *APIClient) InitSubAPIs() {
	c.AccountAPI = NewAccountAPI(c.httpManager)
	c.MarketAPI = NewMarketAPI(c.httpManager)
	c.OrderAPI = NewOrderAPI(c.httpManager)
}
