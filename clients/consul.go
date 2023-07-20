package clients

import "github.com/hashicorp/consul/api"

func NewConsulAPI() (*api.Client, error) {
	cfg := api.DefaultConfig()
	consul, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return consul, nil
}
