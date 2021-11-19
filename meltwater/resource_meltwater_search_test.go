package meltwater

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
	"github.com/hashicorp/terraform/helper/schema"
)

func TestSearchResourceSchema(t *testing.T) {
	resource := SearchResource{}
	actualFullSchema := resource.Schema()

	expectedSchemaSimplified := map[string]SchemaSimplified{
		"type": {
			Type:     schema.TypeString,
			Optional: false,
			Required: true,
		},
		"category": {
			Type:     schema.TypeString,
			Optional: false,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: false,
			Required: true,
		},
		"query": {
			Type:     schema.TypeSet,
			Optional: false,
			Required: true,
			Nested: map[string]SchemaSimplified{
				"keyword": {
					Type:     schema.TypeSet,
					Optional: false,
					Required: true,
					// The test needs to be improved to fix this
					/*Nested: map[string]SchemaSimplified{
						"case_sensitivity": {
							Type:     schema.TypeString,
							Optional: true,
							Required: false,
						},
						"all_keywords": {
							Type:     schema.TypeList,
							Optional: true,
							Required: false,
							Nested:   map[string]SchemaSimplified{},
						},
						"any_keywords": {
							Type:     schema.TypeList,
							Optional: true,
							Required: false,
							Nested:   map[string]SchemaSimplified{},
						},
						"not_keywords": {
							Type:     schema.TypeList,
							Optional: true,
							Required: false,
							Nested:   map[string]SchemaSimplified{},
						},
					},*/
				},
			},
		},
	}

	fullActualSchemaSimplified := map[string]SchemaSimplified{}
	for name, actualSchema := range actualFullSchema {
		actualSchemaSimplified := SchemaSimplified{
			Type:     actualSchema.Type,
			Optional: actualSchema.Optional,
			Required: actualSchema.Required,
		}
		if actualSchema.Type == schema.TypeList || actualSchema.Type == schema.TypeSet {
			nested := map[string]SchemaSimplified{}
			for nestedActualKey, nestedActualSchema := range actualSchema.Elem.(*schema.Resource).Schema {
				nested[nestedActualKey] = SchemaSimplified{
					Type:     nestedActualSchema.Type,
					Optional: nestedActualSchema.Optional,
					Required: nestedActualSchema.Required,
				}
			}
			actualSchemaSimplified.Nested = nested
		}
		fullActualSchemaSimplified[name] = actualSchemaSimplified
	}

	differences := deep.Equal(expectedSchemaSimplified, fullActualSchemaSimplified)
	if len(differences) > 0 {
		t.Errorf("Expected correct simplified schema. Got differences: %+v", differences)
	}
}

func TestSearchTypes(t *testing.T) {
	resource := SearchResource{}
	expectedSearchTypes := []string{
		"news",
		"social",
		"broadcast",
	}

	actualSearchTypes := resource.searchTypes()

	if reflect.DeepEqual(expectedSearchTypes, actualSearchTypes) == false {
		t.Errorf("Expected search types didn't match. Got: %+v", actualSearchTypes)
	}
}

func TestSearchCategories(t *testing.T) {
	resource := SearchResource{}
	expectedSearchCategories := []string{
		"mi",
		"explore",
	}

	actualSearchCategories := resource.searchCategories()

	if reflect.DeepEqual(expectedSearchCategories, actualSearchCategories) == false {
		t.Errorf("Expected time window units didn't match. Got: %+v", actualSearchCategories)
	}
}
