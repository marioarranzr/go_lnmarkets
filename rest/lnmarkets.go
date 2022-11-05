package rest

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"
)

// clientInterface - for testing purpose
type clientInterface interface {
	Do(req *retryablehttp.Request) (*http.Response, error)
}

// LNMarkets - object wraps API
type LNMarkets struct {
	key        string
	secret     string
	passphrase string
	timestamp  string
	client     clientInterface
}

type LNMarketsResponse struct {
	Result interface{}
	Error  struct { // temporary, as if errors were returned as {"error":"blah"}
		Error interface{} `json:"error"`
	}
}

func New(key, secret, passphrase, timestamp string) *LNMarkets {
	if key == "" || secret == "" || passphrase == "" || timestamp == "" {
		log.Print("[WARNING] Missing API key, secret, passphrase or timestamp")
	}
	client := retryablehttp.NewClient()
	client.RetryMax = 10

	return &LNMarkets{
		key:        key,
		secret:     secret,
		passphrase: passphrase,
		timestamp:  timestamp,
		client:     client,
	}
}

func (api *LNMarkets) prepareRequest(method, endpoint string, withAuth bool, params url.Values, body []byte) (*retryablehttp.Request, error) {
	path := fmt.Sprintf("/%s/%s", Version, endpoint)
	requestURL := fmt.Sprintf("%s%s", Url, path)
	rawBody := buildRawBody(body, params)
	req, err := retryablehttp.NewRequest(method, requestURL, bytes.NewBuffer(rawBody))
	if withAuth {
		signature := computeHmac256(api.timestamp, method, path, rawBody, api.secret)
		req.Header.Add("LNM-ACCESS-KEY", api.key)
		req.Header.Add("LNM-ACCESS-PASSPHRASE", api.passphrase)
		req.Header.Add("LNM-ACCESS-SIGNATURE", signature)
	}
	req.Header.Add("LNM-ACCESS-TIMESTAMP", api.timestamp)

	if err != nil {
		return nil, errors.Wrap(err, "error during request creation")
	}

	return req, nil
}

func buildRawBody(body []byte, params url.Values) []byte {
	if body != nil {
		return body
	}
	if params == nil {
		params = url.Values{}
	}
	return []byte(params.Encode())
}

func (api *LNMarkets) parseResponse(response *http.Response, retType interface{}) error {
	if response.StatusCode > 499 {
		return errors.Errorf("error during response parsing: invalid status code %d", response.StatusCode)
	}

	if response.Body == nil {
		return errors.New("error during response parsing: can not read response body")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "error during response parsing: can not read response body")
	}

	if string(body) == "[]" {
		return nil
	}

	var retData LNMarketsResponse
	if retType != nil {
		retData.Result = retType
	}

	if err = json.Unmarshal(body, &retData.Result); err != nil {
		return errors.Wrap(err, "error during response parsing: json marshalling")
	}

	if err = json.Unmarshal(body, &retData.Error); err != nil {
		return errors.Wrap(err, "error during error parsing: json marshalling")
	}

	if response.StatusCode > 299 && response.StatusCode < 500 {
		log.Println(response.Status)
	}
	log.Println(retData.Error)

	return nil
}

func (api *LNMarkets) request(method, endpoint string, withAuth bool, params url.Values, body []byte, retType interface{}) error {
	req, err := api.prepareRequest(method, endpoint, withAuth, params, body)
	if err != nil {
		return err
	}
	resp, err := api.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "error during request execution")
	}

	defer resp.Body.Close()
	return api.parseResponse(resp, retType)
}

func computeHmac256(timestamp, method, path string, rawBody []byte, secret string) string {
	// timestamp + method + path + rawBody
	message := fmt.Sprintf("%s%s%s%s", timestamp, method, path, rawBody)
	log.Println("message", message)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
