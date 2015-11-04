package space

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/request"
)

func doDelete(d *schema.ResourceData, ip interface{}) error {
	return request.DoDelete(
		d, ip, "space",
		fmt.Sprintf("/spaces/%s", d.Id()),
	)
}
