package restapi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func makeClient(d *schema.ResourceData, m interface{}) (*schemaRegistryClient, error) {
	uri := m.(string)
	subject := d.Get("subject").(string)
	schema := d.Get("schema").(string)

	return NewSchemaRegistryClient(uri, subject, schema)
}
