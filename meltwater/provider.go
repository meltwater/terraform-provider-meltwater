package meltwater

import (
	"context"
	"errors"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/meltwater/terraform-provider-meltwater/swagger"
)

// ClientWithContext stores client configuration and context
type ClientWithContext struct {
	Client  *swagger.APIClient
	Context context.Context
}

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {

	// Internal mapping of resources to ensure matching interface
	internalResourceMapping := map[string]TerraformResource{
		"meltwater_recurring_export": RecurringExportResource{},
	}

	// Map to the terraform resource from our internal representation
	providerResourceMap := map[string]*schema.Resource{}
	for resourceName, resource := range internalResourceMapping {
		providerResourceMap[resourceName] = resource.Resource()
	}

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MELTWATER_API_KEY", nil),
			},
		},
		ResourcesMap:  providerResourceMap,
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := swagger.NewConfiguration()
	config.UserAgent = "terraform-provider-meltwater"

	client := swagger.NewAPIClient(config)
	if client == nil {
		log.Println("[INFO] Initializing Meltwater API client with API key")
		return nil, errors.New("could not get meltwater api client")
	}

	authContext := context.WithValue(context.Background(), swagger.ContextAPIKey, swagger.APIKey{
		Key: d.Get("api_key").(string),
	})

	log.Println("[INFO] Initializing Meltwater API client with API key")
	return ClientWithContext{
		Client:  client,
		Context: authContext,
	}, nil
}
