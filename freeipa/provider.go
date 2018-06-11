package main

import (
        "github.com/hashicorp/terraform/helper/schema"
        "github.com/hashicorp/terraform/terraform"

		"github.com/tehwalris/go-freeipa/freeipa"
)

func Provider() *schema.Provider {
    	return &schema.Provider{
                Schema: map[string]*schema.Schema{
			            "server_url": {
				                Type:        schema.TypeString,
				                Required:    true,
				                DefaultFunc: schema.EnvDefaultFunc("FREEIPA_SERVER_URL", nil),
				                Description: "URL of the root of the FreeIPA server.",
			            },
			            "username": {
				                Type:        schema.TypeString,
				                Required:    true,
				                DefaultFunc: schema.EnvDefaultFunc("FREEIPA_USERNAME", nil),
				                Description: "Name of a registered user within the FreeIPA server.",
			            },
			            "password": {
				                Type:        schema.TypeString,
				                Required:    true,
				                DefaultFunc: schema.EnvDefaultFunc("FREEIPA_PASSWORD", nil),
				                Description: "Password of a registered user withing the FreeIPA server.",
			            },
						"allow_unverified_ssl": {
							    Type:        schema.TypeBool,
								Optional:    true,
								Description: "If set, the IPA client will allow insecure SSL certs.",
						},
                },
                ResourcesMap: map[string]*schema.Resource{
                        "hostgroup": resourceHostgroup(),
                },

				ConfigureFunc: providerConfigure
        }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
        tspt := &http.Transport{
			    TLSClientConfig: &tls.Config{
					    InsecureSkipVerify: d.Get("allow_unverified_ssl").(bool),
				},
		}
		c, e := freeipa.Connect(d.Get("server_url").(string), tspt, d.Get("username").(string), d.Get("password").(string))
		if e != nil {
			    log.Fatal(e)
		}

		return c
}
