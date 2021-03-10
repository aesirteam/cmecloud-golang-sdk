package core

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	json "github.com/json-iterator/go"
)

const (
	DEFAULT_HTTP_PORT  = 80
	DEFAULT_HTTPS_PORT = 8443
)

type Config struct {
	ApiGwHost     string `env:"API_GATEWAY_HOST"`
	ApiGwPort     int    `env:"API_GATEWAY_PORT"`
	ApiGwProtocol string `env:"API_GATEWAY_PROTOCOL"`
	AccessKey     string `env:"ACCESS_KEY"`
	SecretKey     string `env:"SECRET_KEY"`
}

type EcloudClient struct {
	Config
}

type Response struct {
	Method     string
	StatusCode int

	RequestId string
	State     string
	Body      string

	ErrCode    string
	ErrMessage string

	SignUrl string
}

type signatureContent struct {
	accessKey string
	secretKey string

	scheme string
	host   string
	port   int

	prefixPath string
	verb       string
	params     map[string]interface{}
}

func signature(c *signatureContent) string {
	loc, _ := time.LoadLocation("Asia/Chongqing")
	t := time.Now().In(loc)
	rand.Seed(t.UnixNano())

	origUrl := url.Values{
		"Version":          {"2016-12-05"},
		"AccessKey":        {c.accessKey},
		"Timestamp":        {t.Format("2006-01-02T15:04:05Z")},
		"SignatureMethod":  {"HmacSHA1"},
		"SignatureNonce":   {strconv.Itoa(rand.Int())},
		"SignatureVersion": {"V2.0"},
	}

	if c.params != nil {
		for k, v := range c.params {
			origUrl.Add(k, fmt.Sprintf("%v", v))
		}
	}

	canonicalQueryString := origUrl.Encode()
	//fmt.Println(canonicalQueryString)

	ha := sha256.New()
	ha.Write([]byte(canonicalQueryString))
	stringToSign := c.verb + "\n" + url.QueryEscape(c.prefixPath) + "\n" + fmt.Sprintf("%x", ha.Sum(nil))
	//fmt.Println(stringToSign)

	h := hmac.New(sha1.New, []byte(fmt.Sprintf("BC_SIGNATURE&%s", c.secretKey)))
	h.Write([]byte(stringToSign))
	_signature := fmt.Sprintf("%x", h.Sum(nil))

	return fmt.Sprintf("%s://%s:%d%s?%s&Signature=%s", c.scheme, c.host, c.port, c.prefixPath, canonicalQueryString, _signature)
}

func New(c *Config) *EcloudClient {
	if c == nil {
		return nil
	}
	return &EcloudClient{*c}
}

func (c *EcloudClient) NewRequest(verb, prefixPath string, headers, params, body map[string]interface{}) (_resp Response, err error) {
	if c.Config.AccessKey == "" {
		err = errors.New("No AccessKey is available")
		return
	}

	if c.Config.SecretKey == "" {
		err = errors.New("No SecretKey is available")
		return
	}

	if c.Config.ApiGwHost == "" {
		c.Config.ApiGwHost = "api-guangzhou-2.cmecloud.cn"
	}

	if c.Config.ApiGwProtocol == "" {
		c.Config.ApiGwProtocol = "http"
	} else if c.Config.ApiGwProtocol != "http" && c.Config.ApiGwProtocol != "https" {
		err = errors.New("Invalid 'protocol_type', should be 'http' or 'https'")
		return
	}

	if c.Config.ApiGwPort == 0 {
		switch c.Config.ApiGwProtocol {
		case "https":
			c.Config.ApiGwPort = DEFAULT_HTTPS_PORT
		default:
			c.Config.ApiGwPort = DEFAULT_HTTP_PORT
		}
	}

	signUrl := signature(&signatureContent{
		accessKey:  c.Config.AccessKey,
		secretKey:  c.Config.SecretKey,
		scheme:     c.Config.ApiGwProtocol,
		host:       c.Config.ApiGwHost,
		port:       c.Config.ApiGwPort,
		prefixPath: prefixPath,
		verb:       verb,
		params:     params,
	})

	_resp = Response{SignUrl: signUrl, Method: verb}

	client := &http.Client{}
	var reqBody io.Reader

	switch verb {
	case "GET", "DELETE":
		reqBody = strings.NewReader("")
	case "PUT", "POST":
		if body == nil {
			body = map[string]interface{}{}
		}

		_bytes, err := json.Marshal(body)
		if err != nil {
			return _resp, err
		}
		reqBody = bytes.NewBuffer(_bytes)
	}

	req, err := http.NewRequest(verb, signUrl, reqBody)
	if err != nil {
		return _resp, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, fmt.Sprintf("%v", v))
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
		err = obj.LastError()
		return
	}

	_resp.RequestId = obj.Get("requestId").ToString()
	_resp.State = obj.Get("state").ToString()
	_resp.Body = obj.Get("body").ToString()
	_resp.ErrCode = obj.Get("errorCode").ToString()
	_resp.ErrMessage = obj.Get("errorMessage").ToString()

	if _resp.StatusCode != http.StatusOK || _resp.State != "OK" {
		err = errors.New(fmt.Sprintf("%s:%s", _resp.ErrCode, _resp.ErrMessage))
		return
	}

	return
}

func (r *Response) Error(err error) error {
	return fmt.Errorf("%s %s [%d] %s", r.Method, r.SignUrl, r.StatusCode, err)
}
