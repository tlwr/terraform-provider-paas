package main

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataApp() *schema.Resource {
	return &schema.Resource{
		Read: dataAppRead,

		Schema: map[string]*schema.Schema{
			"org_guid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"space_guid": &schema.Schema{
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

			"buildpack": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Computed: true,
			},

			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
				Computed: true,
			},

			"instances": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
				Computed: true,
			},
		},
	}
}

func dataAppRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*ProviderConfig).CFClient

	orgGUID := d.Get("org_guid").(string)
	spaceGUID := d.Get("space_guid").(string)

	name := d.Get("name").(string)
	app, err := client.AppByName(name, spaceGUID, orgGUID)

	if err != nil {
		return fmt.Errorf(
			"Could not find app with name '%s', spaceGUID '%s', orgGUID (%s): %s",
			name,
			spaceGUID,
			orgGUID,
			err,
		)
	}

	d.SetId(app.Guid)
	d.Set("guid", app.Guid)
	d.Set("memory", app.Buildpack)
	d.Set("memory", app.Memory)
	d.Set("memory", app.Instances)

	return nil
}
