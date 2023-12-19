package utils

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"time"
)

var (
	globalTransport  *http.Transport
	globalHttpClient *http.Client
)

const (
	MaxIdleConnections int = 100
)

func init() {
	//  连接池
	globalTransport = &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true}, // 忽略 https 证书验证
		MaxIdleConnsPerHost: MaxIdleConnections,
	}

	globalHttpClient = &http.Client{
		Transport: globalTransport,
		Timeout:   3 * time.Second,
	}
}

// PostJson
//
//	@Description: post json请求
//	@param reqBody
//	@param url
//	@return []byte
//	@return error
func PostJson(reqBody interface{}, url string) ([]byte, error) {
	client := resty.NewWithClient(globalHttpClient).R().
		SetHeader("Content-Type", "application/json")

	resp, err := client.SetBody(reqBody).Post(url)
	if err != nil {

		return nil, err
	}
	// 处理响应
	return resp.Body(), nil
}

// Get
//
//	@Description: get请求
//	@param url
//	@param params
//	@param headers
//	@return []byte
//	@return error
func Get(url string, params, headers map[string]string) ([]byte, error) {
	client := resty.New().R()

	for k := range headers {
		client.SetHeader(k, headers[k])
	}
	resp, err := client.SetQueryParams(params).Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

// EncodeParams
//
//	@Description: 编码请求参数
//	@param params
//	@return string
func EncodeParams(params map[string]string) string {
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	return values.Encode()
}
