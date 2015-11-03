package space

import "github.com/hashicorp/terraform/helper/schema"

func NewResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Each space has a unique numeric ID",
				Computed:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Unique name for space",
				Required:    true,
			},
		},
		Create: doCreate,
		Read:   doRead,
		Update: doUpdate,
		Delete: doDelete,
		Exists: doExists,
	}
}
