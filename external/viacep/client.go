package viacep

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type Client struct {
	api *resty.Client
}

func NewClient() *Client {
	client := resty.New()
	client.SetTimeout(5 * time.Second)

	return &Client{
		api: client,
	}
}

func (z *Client) GetAddressByZipCode(zipcode string) (*Address, error) {
	url := "https://viacep.com.br/ws/" + zipcode + "/json/"

	var result Address
	resp, err := z.api.R().
		SetResult(&result).
		Get(url)
	if err != nil {
		return nil, err
	}

	// Verify if the response is valid
	if resp.StatusCode() == 400 || result.CEP == "" {
		return nil, fmt.Errorf("CEP inválido ou não encontrado")
	}

	return &result, nil
}
