package tkxclient

// APIClient - APIClient struct
type APIClient struct {
	region      string
	httpManager *HttpManager
	AccountAPI  *AccountAPI
	MarketAPI   *MarketAPI
}

// NewAPIClient - APIClient constructor
func NewAPIClient(region string) *APIClient {
	c := &APIClient{
		region: region,
	}
	c.Init()
	return c
}

// NewAPIClientWithCredentials - APIClient factory (create with params)
func NewAPIClientWithCredentials(region, clientID, clientSecret string) *APIClient {
	client := NewAPIClient(region)
	client.httpManager.SetCredentials(clientID, clientSecret)
	return client
}

// Init - Initialize client's dependencies
func (c *APIClient) Init() {
	c.httpManager = NewHttpManager(c.region)
	c.AccountAPI = NewAccountAPI(c.httpManager)
	c.MarketAPI = NewMarketAPI(c.httpManager)
}
