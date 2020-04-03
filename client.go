package mcpayment

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Client ...
type Client struct {
	BaseURLToken string
	XSignKey     string
	LogLevel     int
	Logger       *log.Logger
	ReqTimeout   time.Duration
}

// NewClient getting default client
// 0: No logging
// 1: Errors only
// 2: Errors + informational (default)
// 3: Errors + informational + debug
func NewClient() Client {
	return Client{
		LogLevel:   2,
		Logger:     log.New(os.Stderr, "", log.LstdFlags),
		ReqTimeout: 15 * time.Second,
	}
}

// newRequest create http request
func (c *Client) newRequest(method string, fullPath string, headers map[string]string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		if c.LogLevel > 0 {
			c.Logger.Printf(PaymentName, " Request creation failed: %s\n", err)
		}
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-sign-key", c.XSignKey)

	if headers != nil {
		for k, vv := range headers {
			req.Header.Set(k, vv)
		}
	}

	return req, nil
}

// executeRequest executing the request, respModel should pass by reference as it will be filled to corresponded struct
func (c *Client) executeRequest(req *http.Request, respModel interface{}) error {
	httpClient := &http.Client{Timeout: c.ReqTimeout}

	if c.LogLevel > 1 {
		c.Logger.Println(PaymentName, " Request ", req.Method, ": ", req.URL.Host, req.URL.Path)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		if c.LogLevel > 0 {
			c.Logger.Println(PaymentName, " Cannot send request: ", err)
		}
		return err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if c.LogLevel > 0 {
			c.Logger.Println(PaymentName, " Cannot read response body: ", err)
		}
		return err
	}

	if c.LogLevel > 2 {
		c.Logger.Println(PaymentName, " Response: ", string(resBody))
	}

	if isOK(200) {
		return fmt.Errorf("HTTP Status: %d", res.StatusCode)
	}

	if err = json.Unmarshal(resBody, respModel); err != nil {
		if c.LogLevel > 0 {
			c.Logger.Println(PaymentName, " Cannot unmarshal response body: ", err)
		}
		return err
	}

	return nil
}

// Call call the API, respModel should pass by reference as it will be filled to corresponded struct
func (c *Client) Call(method string, path string, header map[string]string, body io.Reader, respModel interface{}) error {
	req, err := c.newRequest(method, path, header, body)
	if err != nil {
		return err
	}

	return c.executeRequest(req, respModel)
}
