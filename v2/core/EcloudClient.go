package core

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/caarlos0/env/v6"
	json "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type AcsClient struct {
	ApiGwHost     string `env:"API_GATEWAY_HOST"`
	ApiGwPort     int    `env:"API_GATEWAY_PORT"`
	ApiGwProtocol string `env:"API_GATEWAY_PROTOCOL"`
	AccessKey     string `env:"ACCESS_KEY"`
	SecretKey     string `env:"SECRET_KEY"`
}

type AcsRequest struct {
	serverPath string
	url        string
	method     string
	header     AcsCustomHeader
	param      AcsCustomParam
	body       AcsRequestBody
}

type AcsResponse struct {
	SignUrl    string
	Method     string
	StatusCode int
	RequestId  string
	State      string
	Body       string
	ErrCode    string
	ErrMessage string
}

type AcsCustomHeader map[string]interface{}
type AcsCustomParam map[string]interface{}
type AcsRequestBody map[string]interface{}

func (c *AcsClient) signature(r *AcsRequest) (string, error) {
	if err := env.Parse(c); err != nil {
		return "", err
	}

	if c.AccessKey == "" {
		return "", errors.New("No access key is available")
	}

	if c.SecretKey == "" {
		return "", errors.New("No secret key is available")
	}

	if c.ApiGwHost == "" {
		c.ApiGwHost = "api-guangzhou-2.cmecloud.cn"
	}

	if c.ApiGwProtocol == "" {
		c.ApiGwProtocol = "http"
	} else if c.ApiGwProtocol != "http" && c.ApiGwProtocol != "https" {
		return "", errors.New("Invalid 'protocol_type', should be 'http' or 'https'")
	}

	if c.ApiGwPort == 0 {
		switch c.ApiGwProtocol {
		case "https":
			c.ApiGwPort = 8443
		default:
			c.ApiGwPort = 80
		}
	}

	loc, _ := time.LoadLocation("Asia/Chongqing")
	t := time.Now().In(loc)
	rand.Seed(t.UnixNano())

	params := url.Values{}
	params.Set("Version", "2016-12-05")
	params.Set("AccessKey", c.AccessKey)
	params.Set("Timestamp", t.Format("2006-01-02T15:04:05Z"))
	params.Set("SignatureMethod", "HmacSHA1")
	params.Set("SignatureNonce", strconv.Itoa(rand.Int()))
	params.Set("SignatureVersion", "V2.0")

	if r.param != nil {
		for k, v := range r.param {
			params.Set(k, fmt.Sprintf("%v", v))
			//fmt.Printf("%s=%v\n", k, v)
		}
	}

	canonicalQueryString := params.Encode()
	//fmt.Println(canonicalQueryString)

	ha := sha256.New()
	ha.Write([]byte(canonicalQueryString))
	stringToSign := r.method + "\n" + url.QueryEscape(r.serverPath) + "\n" + fmt.Sprintf("%x", ha.Sum(nil))
	/*+ "\n" + canonical_headers + "\n" + signed_headers + "\n" + payload_hash*/
	//fmt.Println(stringToSign)

	signKey := fmt.Sprintf("BC_SIGNATURE&%s", c.SecretKey)

	h := hmac.New(sha1.New, []byte(signKey))
	h.Write([]byte(stringToSign))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	r.url = fmt.Sprintf("%s://%s:%d%s?%s&Signature=%s", c.ApiGwProtocol, c.ApiGwHost, c.ApiGwPort, r.serverPath, canonicalQueryString, signature)

	return r.url, nil
}

func (c *AcsClient) request(r *AcsRequest) (*AcsResponse, error) {
	_resp := &AcsResponse{SignUrl: r.url, Method: r.method}

	client := &http.Client{}
	var reqBody io.Reader

	switch r.method {
	case "GET":
		reqBody = strings.NewReader("")
	case "POST":
		if r.body == nil {
			r.body = AcsRequestBody{}
		}

		_bytes, err := json.Marshal(r.body)
		if err != nil {
			return _resp, err
		}
		reqBody = bytes.NewBuffer(_bytes)
	}

	req, err := http.NewRequest(r.method, r.url, reqBody)
	if err != nil {
		return _resp, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if r.header != nil {
		for k, v := range r.header {
			req.Header.Add(k, fmt.Sprintf("%v", v))
			//fmt.Printf("%s=%v\n", k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return _resp, err
	}
	defer resp.Body.Close()

	_resp.StatusCode = resp.StatusCode

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return _resp, err
	}

	obj := json.Get(data)
	if obj.LastError() != nil {
		return _resp, err
	}

	_resp.RequestId = obj.Get("requestId").ToString()
	_resp.State = obj.Get("state").ToString()
	_resp.Body = obj.Get("body").ToString()
	_resp.ErrCode = obj.Get("errorCode").ToString()
	_resp.ErrMessage = obj.Get("errorMessage").ToString()

	if _resp.StatusCode != http.StatusOK || _resp.State != "OK" {
		return _resp, errors.New(_resp.ErrCode)
	}

	return _resp, nil
}

func (c *AcsClient) Get(serverPath string, header AcsCustomHeader, param AcsCustomParam) (*AcsResponse, error) {
	req := &AcsRequest{
		method:     "GET",
		serverPath: serverPath,
		header:     header,
		param:      param,
	}

	url, err := c.signature(req)
	if err != nil {
		return nil, err
	}

	req.url = url
	return c.request(req)
}

func (c *AcsClient) Post(serverPath string, header AcsCustomHeader, param AcsCustomParam, body AcsRequestBody) (*AcsResponse, error) {
	req := &AcsRequest{
		method:     "POST",
		serverPath: serverPath,
		header:     header,
		param:      param,
		body:       body,
	}

	url, err := c.signature(req)
	if err != nil {
		return nil, err
	}

	req.url = url
	return c.request(req)
}

func (c *AcsClient) Delete(serverPath string, header AcsCustomHeader, param AcsCustomParam) (*AcsResponse, error) {
	req := &AcsRequest{
		method:     "DELETE",
		serverPath: serverPath,
		header:     header,
		param:      param,
	}

	url, err := c.signature(req)
	if err != nil {
		return nil, err
	}

	req.url = url
	return c.request(req)
}
