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
			Type:     schema.TypeList,
			Optional: false,
			Required: true,
			Nested: map[string]SchemaSimplified{
				"keyword": {
					Type:     schema.TypeList,
					Optional: true,
					Required: false,
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
				"boolean": {
					Type:     schema.TypeList,
					Optional: true,
					Required: false,
				},
				"combined": {
					Type:     schema.TypeList,
					Optional: true,
					Required: false,
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

func TestSearchWithEmptyQuery(t *testing.T) {
	resource := SearchResource{}
	validateFun := resource.validateQuery()
	_, errs := validateFun([]interface{}{}, "")
	if len(errs) != 1 {
		t.Errorf("Expected a single error about the query block")
	}
	if errs[0].Error() != "you need exactly one query block, got 0" {
		t.Errorf("Unexpected error message: %s", errs[0].Error())
	}
}

func TestSearchWithNoQueryTypesGiven(t *testing.T) {
	resource := SearchResource{}
	validateFun := resource.validateQuery()
	queryList := []interface{}{
		map[string]interface{}{
			"boolean":  []interface{}{},
			"combined": []interface{}{},
			"keyword":  []interface{}{},
		},
	}
	_, errs := validateFun(queryList, "")
	if len(errs) != 1 {
		t.Errorf("Expected a single error about the query block: %+v", errs)
	}
	if errs[0].Error() != "one of keyword, combined or boolean must be set within query, got 0" {
		t.Errorf("Unexpected error message: %s", errs[0].Error())
	}
}
func TestSearchWithMultipleQueryTypesGiven(t *testing.T) {
	resource := SearchResource{}
	validateFun := resource.validateQuery()
	queryList := []interface{}{
		map[string]interface{}{
			"boolean": []interface{}{
				map[string]interface{}{},
				map[string]interface{}{},
			},
			"combined": []interface{}{},
			"keyword":  []interface{}{},
		},
	}
	_, errs := validateFun(queryList, "")
	if len(errs) != 1 {
		t.Errorf("Expected a single error about the query block: %+v", errs)
	}
	if errs[0].Error() != "one of keyword, combined or boolean can exist within query, got 2" {
		t.Errorf("Unexpected error message: %s", errs[0].Error())
	}
}
func TestSearchWithSingleQueryTypeGiven(t *testing.T) {
	resource := SearchResource{}
	validateFun := resource.validateQuery()
	queryList := []interface{}{
		map[string]interface{}{
			"boolean": []interface{}{
				map[string]interface{}{},
			},
			"combined": []interface{}{},
			"keyword":  []interface{}{},
		},
	}
	_, errs := validateFun(queryList, "")
	if len(errs) != 0 {
		t.Errorf("Unexpected error message(s) returned: %+v", errs)
	}
}
