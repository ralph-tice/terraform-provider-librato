package space

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

type createBody struct {
	Name string `json:"name"`
}

func doCreate(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] doCreate new space")

	p := ip.(*provider.Provider)
	body := &createBody{Name: d.Get("name").(string)}

	resp := &readResponse{}
	_, err := request.DoRequest(
		"POST",
		"/spaces",
		p,
		body,
		resp,
		201,
	)
	if err != nil {
		return fmt.Errorf("doCreate space failed: %v", err)
	}

	log.Printf("[DEBUG] doCreate space: %#v", resp)
	d.SetId(strconv.Itoa(resp.Id))
	return nil
}
