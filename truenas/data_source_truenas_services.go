package truenas

import (
	"context"
	"strconv"
	"time"

	api "github.com/dellathefella/truenas-go-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTrueNASServices() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTrueNASServicesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func dataSourceTrueNASServicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*api.APIClient)

	services, _, _ := c.ServiceApi.ListServicesExecute(c.ServiceApi.ListServices(ctx))

	converted := flattenServicesResponse(services)

	if err := d.Set("ids", converted); err != nil {
		return diag.FromErr(err)
	}

	// always run
	// TODO: not sure yet what should go here, find out
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenServicesResponse(services []api.Service) []int32 {
	result := make([]int32, len(services))

	for i, service := range services {
		result[i] = service.Id
	}

	return result
}
