package restapi

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSubject() *schema.Resource {
	return &schema.Resource{
		Create: resourceSubjectCreate,
		Read:   resourceSubjectRead,
		Update: resourceSubjectUpdate,
		Delete: resourceSubjectDelete,

		Schema: map[string]*schema.Schema{
			"subject": {
				Type:     schema.TypeString,
				Required: true,
			},
			"schema": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSubjectCreate(d *schema.ResourceData, m interface{}) error {
	client, err := makeClient(d, m)
	if err != nil {
		return err
	}

	log.Printf("Create subject '%v'.", client)

	if err = client.createSubject(); err != nil {
		return err
	}

	d.SetId(client.subject)
	return nil
}

func resourceSubjectRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSubjectUpdate(d *schema.ResourceData, m interface{}) error {
	client, err := makeClient(d, m)
	if err != nil {
		return err
	}

	log.Printf("Update subject '%v'.", client)

	return client.createSubject()
}

func resourceSubjectDelete(d *schema.ResourceData, m interface{}) error {
	client, err := makeClient(d, m)
	if err != nil {
		return err
	}

	log.Printf("Delete subject '%v'.", client)

	if err = client.deleteSubject(); err != nil {
		return err
	}

	d.SetId("")

	return nil
}
