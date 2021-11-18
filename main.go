package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/meltwater/terraform-provider-meltwater/meltwater"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: meltwater.Provider})
}
