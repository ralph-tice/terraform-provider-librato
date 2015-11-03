package space

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/underarmour/terraform-provider-librato/provider"
	"github.com/underarmour/terraform-provider-librato/request"
)

type readResponse struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func doRead(d *schema.ResourceData, ip interface{}) error {
	log.Printf("[DEBUG] doRead space")

	p := ip.(*provider.Provider)
	resp := &readResponse{}

	_, err := request.DoRequest(
		"GET",
		fmt.Sprintf("/spaces/%s", d.Id()),
		p,
		nil,
		resp,
		200,
	)
	if err != nil {
		return fmt.Errorf("doRead space failed: %v", err)
	}

	log.Printf("[DEBUG] doRead space: %#v", resp)
	d.Set("name", resp.Name)
	return nil
}
