package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CF_USERNAME", nil),
				Description: "Username for Cloud Foundry",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CF_PASSWORD", nil),
				Description: "Password for Cloud Foundry",
			},
			"api_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CF_API", nil),
				Description: "API URL for Cloud Foundry",
			},
		},
		ConfigureFunc: ProviderConfigurationBuilder,

		DataSourcesMap: map[string]*schema.Resource{
			"paas_org":           dataOrg(),
			"paas_space":         dataSpace(),
			"paas_domain":        dataDomain(),
			"paas_shared_domain": dataSharedDomain(),
		},
	}
}
