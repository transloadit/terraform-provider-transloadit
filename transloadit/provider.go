package transloadit

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"gopkg.in/transloadit/go-sdk.v1"
)

const prefix = "transloadit_"

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"auth_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TRANSLOADIT_AUTH_KEY", ""),
				Description: "TransloadIt API Authentication Key",
			},
			"auth_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TRANSLOADIT_AUTH_SECRET", ""),
				Description: "TransloadIt API Authentication Secret",
			},
			"endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			prefix + "template": resourceTemplate(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	options := transloadit.DefaultConfig
	if endpoint, ok := d.GetOk("endpoint"); ok {
		options.Endpoint = endpoint.(string)
	}
	options.AuthKey = d.Get("auth_key").(string)
	options.AuthSecret = d.Get("auth_secret").(string)
	if options.AuthKey == "" || options.AuthSecret == "" {
		return nil, fmt.Errorf("Missing TransloadIt AuthKey ot AuthSecret")
	}
	client := transloadit.NewClient(options)
	return &client, nil
}
