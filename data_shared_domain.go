package main

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSharedDomain() *schema.Resource {
	return &schema.Resource{
		Read: dataSharedDomainRead,

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

func dataSharedDomainRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*ProviderConfig).CFClient

	name := d.Get("name").(string)
	sharedDomain, err := client.GetSharedDomainByName(name)

	if err != nil {
		return fmt.Errorf(
			"Could not find shared domain with name '%s': %s",
			name,
			err,
		)
	}

	d.SetId(sharedDomain.Guid)
	d.Set("guid", sharedDomain.Guid)

	return nil
}
