package tkxsdk

import (
	"fmt"
	"github.com/imroc/req"
)

var (
	defaultScheme = "https"
	subdomain     = "api"
	prefix        = "/public/v1"
	domainMY      = "tokenize-my.com"
	domainSG      = "tokenize.exchange"
	domainGobal   = "tokenize.exchange"
)

// HttpManager - HttpManager struct
type HttpManager struct {
	region      string
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
func NewHttpManager(region string) *HttpManager {
	m := &HttpManager{
		region: region,
	}

	switch m.region {
	case "MY":
		m.domain = domainMY
		break
	case "SG":
		m.domain = domainSG
		break
	default:
		m.domain = domainGobal
	}
	m.baseURL = fmt.Sprintf("%s://%s.%s%s", defaultScheme, subdomain, m.domain, prefix)
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
