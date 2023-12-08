package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/al-kirpichenko/gophkeeper/cmd/client/cfg"
)

// Client -
type Client struct {
	ServerURL  string
	HTTPClient *http.Client
}

// NewClient - конструктор
func NewClient(cfg *cfg.Cfg) *Client {
	return &Client{
		ServerURL: cfg.ServerURL,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) Welcome() {
	fmt.Println("Welcome to gophkeeper client!")
}
