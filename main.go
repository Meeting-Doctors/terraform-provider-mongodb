package main

import (
	"github.com/fabiovpcaumo/terraform-provider-mongodb/mongodb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mongodb.Provider,
	})
}
