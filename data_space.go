package main

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSpace() *schema.Resource {
	return &schema.Resource{
		Read: dataSpaceRead,

		Schema: map[string]*schema.Schema{
			"org_guid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

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

func dataSpaceRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*ProviderConfig).CFClient

	orgGUID := d.Get("org_guid").(string)
	org, err := client.GetOrgByGuid(orgGUID)
	if err != nil {
		return fmt.Errorf("Could not find org with guid '%s': %s", orgGUID, err)
	}

	name := d.Get("name").(string)
	space, err := client.GetSpaceByName(name, orgGUID)

	if err != nil {
		return fmt.Errorf(
			"Could not find space with name '%s' in org '%s' (%s): %s",
			name,
			org.Name,
			orgGUID,
			err,
		)
	}

	d.SetId(space.Guid)
	d.Set("guid", space.Guid)
	d.Set("org_guid", orgGUID)

	return nil
}
