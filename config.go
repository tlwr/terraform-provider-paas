package main

import (
	"fmt"
	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/hashicorp/terraform/helper/schema"
)

type ProviderConfig struct {
	Username string
	Password string
	APIURL   string
	CFClient cfclient.Client
}

func ProviderConfigurationBuilder(
	d *schema.ResourceData,
) (interface{}, error) {

	username := d.Get("username")
	password := d.Get("password")
	apiURL := d.Get("api_url")

	client, err := cfclient.NewClient(&cfclient.Config{
		ApiAddress: apiURL.(string),
		Username:   username.(string),
		Password:   password.(string),
	})

	if err != nil {
		return nil, fmt.Errorf("Error creating CF client: %s", err)
	}

	return &ProviderConfig{
		Username: username.(string),
		Password: password.(string),
		APIURL:   apiURL.(string),
		CFClient: interface{}(*client).(cfclient.Client),
	}, nil
}
