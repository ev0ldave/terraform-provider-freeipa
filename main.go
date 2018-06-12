package main

import (
        "github.com/hashicorp/terraform/plugin"
		"freeipa"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: freeipa.Provider})
}
