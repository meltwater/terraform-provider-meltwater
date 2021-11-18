package meltwater

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
	"github.com/hashicorp/terraform/helper/schema"
)

func TestProductResourceSchema(t *testing.T) {
	resource := CatalogProductResource{}
	actualFullSchema := resource.Schema()

	expectedSchemaSimplified := map[string]SchemaSimplified{
		"name": {
			Type:     schema.TypeString,
			Optional: false,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Required: false,
		},
		"image_url": {
			Type:     schema.TypeString,
			Optional: true,
			Required: false,
		},
		"home_url": {
			Type:     schema.TypeString,
			Optional: true,
			Required: false,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: false,
			Required: true,
		},
		"category": {
			Type:     schema.TypeString,
			Optional: true,
			Required: false,
		},
	}

	actualSchemaSimplified := map[string]SchemaSimplified{}
	for name, actualSchema := range actualFullSchema {
		actualSchemaSimplified[name] = SchemaSimplified{
			Type:     actualSchema.Type,
			Optional: actualSchema.Optional,
			Required: actualSchema.Required,
		}
	}

	differences := deep.Equal(expectedSchemaSimplified, actualSchemaSimplified)
	if len(differences) > 0 {
		t.Errorf("Expected correct simplified schema. Got differences: %+v", differences)
	}
}

func TestProductTypes(t *testing.T) {
	resource := CatalogProductResource{}
	expectedProductTypes := []string{
		"physical",
		"digital",
		"service",
	}

	actualProductTypes := resource.productTypes()

	if reflect.DeepEqual(expectedProductTypes, actualProductTypes) == false {
		t.Errorf("Expected Event types didn't match. Got: %+v", actualProductTypes)
	}
}
