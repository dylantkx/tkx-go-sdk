package tkxsdk

import (
	"github.com/imroc/req"
)

// HttpManager - HttpManager struct
type HttpManager struct {
	baseURL     string
	domain      string
	credentials ClientCredentials
	header      req.Header
}

// ClientCredentials - ClientCredentials struct
type ClientCredentials struct {
	ClientID     string
	ClientSecret string
	Token        string
}

// NewHttpManager - HttpManager constructor
func NewHttpManager(apiBaseURL string) *HttpManager {
	m := &HttpManager{
		baseURL: apiBaseURL,
	}
	m.header = req.Header{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	return m
}

// SetCredentials - Set client's credentials (id + secret)
func (m *HttpManager) SetCredentials(id, secret string) *HttpManager {
	m.credentials.ClientID = id
	m.credentials.ClientSecret = secret
	generator := &TokenGenerator{}
	token, err := generator.Generate(id, secret)
	if err != nil {
		panic("Invalid client id or secret!")
	}
	m.credentials.Token = token
	m.header["X-TKX-TOKEN"] = m.credentials.Token
	return m
}
