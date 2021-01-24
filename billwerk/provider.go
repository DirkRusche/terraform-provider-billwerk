package billwerk

import (
	"billwerk-go"
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"url": {
				Type: schema.TypeString,
				Optional: true,
				//Default: nil,
			},
		},
		ResourcesMap:   map[string]*schema.Resource{
			"billwerk_mail_template": resourceMailTemplate(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	var url *string = nil

	if tmpUrl, ok := d.GetOk("url"); ok {
		urlString := tmpUrl.(string)
		url = &urlString
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c, err := billwerk.NewClient(url, &token)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
