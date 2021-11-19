package meltwater

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
	"github.com/hashicorp/terraform/helper/schema"
)

func TestRecurringExportResourceSchema(t *testing.T) {
	resource := RecurringExportResource{}
	actualFullSchema := resource.Schema()

	expectedSchemaSimplified := map[string]SchemaSimplified{
		"search_id": {
			Type:     schema.TypeInt,
			Optional: false,
			Required: true,
		},
		"timezone": {
			Type:     schema.TypeString,
			Optional: false,
			Required: true,
		},
		"window_time_unit": {
			Type:     schema.TypeString,
			Optional: false,
			Required: true,
		},
		"window_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Required: false,
		},
		"window_time": {
			Type:     schema.TypeString,
			Optional: true,
			Required: false,
		},
		"window_monthday": {
			Type:     schema.TypeInt,
			Optional: true,
			Required: false,
		},
		"window_weekday": {
			Type:     schema.TypeInt,
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

func TestWindowTimeUnits(t *testing.T) {
	resource := RecurringExportResource{}
	expectedTimeWindowUnits := []string{
		"day",
		"week",
		"month",
	}

	actualWindowTimeUnits := resource.windowTimeUnits()

	if reflect.DeepEqual(expectedTimeWindowUnits, actualWindowTimeUnits) == false {
		t.Errorf("Expected time window units didn't match. Got: %+v", actualWindowTimeUnits)
	}
}
