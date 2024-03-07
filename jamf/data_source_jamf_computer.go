package jamf

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/w0de/go-jamf-api"
)

func dataSourceJamfComputer() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJamfCategoryRead,
		Schema: map[string]*schema.Schema{
			// Computed values.
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			// Computed values.
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed values.
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceJamfComputerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	resp, err := c.GetComputerInventories(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.GetId())
	d.Set("name", resp.GetName())
	d.Set("priority", resp.GetPriority())

	return diags
}
