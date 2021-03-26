package global

import (
	"bytes"
	"crypto/hmac"
	_rand "crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
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
	accessKey  string
	secretKey  string
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
	stringToSign := fmt.Sprintf("%s\n%s\n%x", c.verb, url.QueryEscape(c.prefixPath), ha.Sum(nil))
	//fmt.Println(stringToSign)

	h := hmac.New(sha1.New, []byte(fmt.Sprintf("BC_SIGNATURE&%s", c.secretKey)))
	h.Write([]byte(stringToSign))
	_signature := fmt.Sprintf("%x", h.Sum(nil))

	return fmt.Sprintf("%s?%s&Signature=%s", c.prefixPath, canonicalQueryString, _signature)
}

func NewEcloudClient(c *Config) (*EcloudClient, error) {
	if c == nil {
		return nil, errors.New("Config invalid")
	}
	return &EcloudClient{*c}, nil
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

		//fmt.Println(Dump(body))

		_bytes, err := json.Marshal(body)
		if err != nil {
			return _resp, err
		}
		reqBody = bytes.NewBuffer(_bytes)
	}

	req, err := http.NewRequest(verb, fmt.Sprintf("%s://%s:%d%s", c.Config.ApiGwProtocol, c.Config.ApiGwHost, c.Config.ApiGwPort, signUrl), reqBody)
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

	//fmt.Printf("%s\n", data)

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

func (r *Response) UnmarshalFromContent(result interface{}, tagKey string) error {
	var fronze = json.ConfigDefault

	if tagKey != "" {
		fronze = json.Config{
			EscapeHTML:             true,
			SortMapKeys:            true,
			ValidateJsonRawMessage: true,
			TagKey:                 "newtag",
		}.Froze()
	}

	var obj = fronze.Get([]byte(r.Body))
	if obj.LastError() != nil {
		return obj.LastError()
	}

	if content := obj.Get("content"); content.LastError() == nil {
		obj = content
	}

	if err := fronze.UnmarshalFromString(obj.ToString(), result); err != nil {
		return err
	}

	return nil
}

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC/VpRysi0bPRLS7sbgQDJHo1MAt9/bK
+nwK5Pe3z0/O4cH5I/8kFNYy4yFsLMM+zyFvVw9C4wzjHaRcmEuF3ziJMC9PD5ufUWgfO
5nSGgZW1cmgjqnhcWJ3i+Azj72RnhKQRCn9DgJduEC9MiKfbyTICGd6FXf9cxb21nkxI7vtwIDAQAB
-----END PUBLIC KEY-----
`)

func RsaEncrypt(origData []byte) ([]byte, error) {
	//将密钥解析成公钥实例
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	//解析pem.Decode（）返回的Block指针实例
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pub := pubInterface.(*rsa.PublicKey)
	//RSA算法加密
	return rsa.EncryptPKCS1v15(_rand.Reader, pub, origData)
}

func Dump(a interface{}) string {
	if bytes, err := json.MarshalIndent(&a, "", "  "); err == nil {
		return fmt.Sprintf("\n%s", bytes)
	}
	return ""
}
