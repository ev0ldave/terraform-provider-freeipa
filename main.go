package main

import (
        "github.com/hashicorp/terraform/plugin"
        "github.com/ev0ldave/terraform-provider-freeipa/freeipa"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: func() terraform.ResourceProvider {
                        return Provider()
                },
        })
}
