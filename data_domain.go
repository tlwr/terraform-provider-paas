package main

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataDomain() *schema.Resource {
	return &schema.Resource{
		Read: dataDomainRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"guid": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Computed: true,
			},
		},
	}
}

func dataDomainRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*ProviderConfig).CFClient

	name := d.Get("name").(string)
	domain, err := client.GetDomainByName(name)

	if err != nil {
		return fmt.Errorf("Could not find domain with name '%s': %s", name, err)
	}

	d.SetId(domain.Guid)
	d.Set("guid", domain.Guid)

	return nil
}
