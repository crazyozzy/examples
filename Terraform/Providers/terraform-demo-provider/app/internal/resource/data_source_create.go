package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func initResourceCreationSchema() *schema.Resource {
	return &schema.Resource{
		ReadContext: initNameDataSourceRead,

		Schema: map[string]*schema.Schema{
			"raw": {
				Type:     schema.TypeString,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func initNameDataSourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	provider := m.(ProviderClient)
	client := provider.Client

	header := make(http.Header)
	headers, exists := d.GetOk("headers")
	if exists {
		for name, value := range headers.(map[string]interface{}) {
			header.Set(name, value.(string))
		}
	}

	resourceType := d.Get("resource_type").(string)
	if resourceType == "" {
		return diag.Errorf("Invalid resource type specified")
	}
	region := d.Get("region").(string)
	if region == "" {
		return diag.Errorf("Invalid region specified")
	}
	b, err := client.doCreateResource(client.BaseUrl.String(), resourceType, region)
	if err != nil {
		return diag.FromErr(err)
	}
	outputs, err := flattenNameAllocationResponse(b)
	if err != nil {
		return diag.FromErr(err)
	}
	marshalData(d, outputs)

	return diags
}

func flattenNameAllocationResponse(b []byte) (outputs map[string]interface{}, err error) {
	var data map[string]interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		err = fmt.Errorf("Cannot unmarshal json of API response: %v", err)
		return
	} else if data["result"] == "" {
		err = fmt.Errorf("missing result key in API response: %v", err)
		return
	}

	outputs = make(map[string]interface{})
	outputs["id"] = time.Now().UTC().String()
	outputs["raw"] = string(b)
	outputs["name"] = data["Name"]

	return
}
