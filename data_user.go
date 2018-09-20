package main

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tlwr/go-cfclient"
)

func dataUser() *schema.Resource {
	return &schema.Resource{
		Read: dataUserRead,

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"org_guid": &schema.Schema{
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

func dataUserRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*ProviderConfig).CFClient

	orgGUID := d.Get("org_guid").(string)
	users, err := client.ListOrgUsers(orgGUID)
	if err != nil {
		return fmt.Errorf("Could not lookup users", err)
	}

	username := d.Get("username").(string)
	var retUser *cfclient.User

	for _, user := range users {
		if user.Username == username {
			retUser = &user
		}
	}

	if retUser == nil {
		return fmt.Errorf(
			"Could not find user '%s' in org '%s'",
			username,
			orgGUID,
			err,
		)
	}

	d.SetId(retUser.Guid)
	d.Set("guid", retUser.Guid)

	return nil
}
