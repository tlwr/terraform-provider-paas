package main

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataOrg() *schema.Resource {
	return &schema.Resource{
		Read: dataOrgRead,

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

func dataOrgRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*ProviderConfig).CFClient

	name := d.Get("name").(string)
	org, err := client.GetOrgByName(name)

	if err != nil {
		return fmt.Errorf("Could not find org with name '%s': %s", name, err)
	}

	d.SetId(org.Guid)
	d.Set("guid", org.Guid)

	return nil
}
