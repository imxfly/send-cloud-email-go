package sendcloud

// The official document: https://www.sendcloud.net/doc/email_v2/
import (
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "net/url"
)

// Client
type Client struct {
    APIUser string
    APIKey  string
}

// Response is a SendCloud API response. This wraps the standard http.Response
// returned from SendCloud and provides convenient access to data.
type Response struct {
    Result     bool
    StatusCode int
    Message    string
    Info       interface{}
}

// Send 普通发送
func (c *Client) Send(params map[string]string) (*Response, error) {
    formData := url.Values{}
    formData.Add("apiUser", c.APIUser)
    formData.Add("apiKey", c.APIKey)
    for key, value := range params {
        formData.Add(key, value)
    }

    resp, err := http.PostForm("https://api.sendcloud.net/apiv2/mail/send", formData)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    r := &Response{}
    if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
        return nil, err
    }

    if r.Result == false || r.StatusCode != 200 {
        return nil, errors.New(fmt.Sprintf("API request failed: [code %d] - [Message: %s]", r.StatusCode, r.Message))
    }

    return r, nil
}

// SendTemplate 模板发送
func (c *Client) SendTemplate(params map[string]string) (*Response, error) {
    formData := url.Values{}
    formData.Add("apiUser", c.APIUser)
    formData.Add("apiKey", c.APIKey)
    for key, value := range params {
        formData.Add(key, value)
    }

    resp, err := http.PostForm("https://api.sendcloud.net/apiv2/mail/sendtemplate", formData)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    r := &Response{}
    if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
        return nil, err
    }

    if r.Result == false || r.StatusCode != 200 {
        return nil, errors.New(fmt.Sprintf("API request failed: [code %d] - [Message: %s]", r.StatusCode, r.Message))
    }

    return r, nil
}
