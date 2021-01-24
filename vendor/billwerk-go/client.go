package billwerk

import (
	"io/ioutil"
	"net/http"
	"time"
)

const HostURL string = "https://app.billwerk.com"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

func NewClient(host, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL: HostURL,
		Token: *token,
	}

	if host != nil {
		c.HostURL = *host
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, int, error) {
	req.Header.Set("Authorization", "Bearer " + c.Token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, -1, err
	}

	return body, res.StatusCode, err
}
