package main

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataUser() *schema.Resource {
	return &schema.Resource{
		Read: dataUserRead,

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
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

	users, err := client.ListUsers()
	if err != nil {
		return fmt.Errorf("Could not lookup users", err)
	}

	username := d.Get("username").(string)
	user := users.GetUserByUsername(username)

	d.SetId(user.Guid)
	d.Set("guid", user.Guid)

	return nil
}
