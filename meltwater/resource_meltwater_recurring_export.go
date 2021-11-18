package meltwater

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/meltwater/terraform-provider-meltwater/swagger"
)

type RecurringExportResource struct{}

func (r RecurringExportResource) Resource() *schema.Resource {
	return &schema.Resource{
		Schema: r.Schema(),
		Create: r.Create,
		Read:   r.Read,
		Update: r.Update,
		Delete: r.Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func (r RecurringExportResource) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"search_id": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The ID of the search",
		},
		"timezone": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The timezone for the window to be based on",
		},
		"window_time_unit": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(r.windowTimeUnits(), true),
			Description:  fmt.Sprintf("A window time unit must be one of: %s", strings.Join(r.windowTimeUnits(), ",")),
		},
		"window_size": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      1,
			ValidateFunc: validation.IntAtMost(20),
			Description:  "The number of 'window_time_unit's the export window covers Defaults to 1",
		},
		"window_time": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "00:00:00",
			Description: "The time of day the export window begins. Defaults to '00:00:00'",
		},
		"window_monthday": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      1,
			ValidateFunc: validation.IntBetween(1, 28),
			Description:  "Representation of what day of the month the export window starts. Integer values 1-28 represent the day of the month e.g. 1 represents 1st, 2 represents the 2nd, etc. Integer value 0 represents the last day of the month. Defaults to 1",
		},
		"window_weekday": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      1,
			ValidateFunc: validation.IntBetween(1, 7),
			Description:  "Representation of what day of the week the export window starts. Integer values are: Monday=1, Tuesday=2, Wednesday=3, Thursday=4, Friday=5, Saturday=6, Sunday=7 Defaults to 1",
		},
	}
}

// Create - Creating a recurring export in the Meltwater API
func (r RecurringExportResource) Create(d *schema.ResourceData, m interface{}) error {
	clientWithContext := m.(ClientWithContext)
	client := clientWithContext.Client
	context := clientWithContext.Context

	if 1 == 0 {
		return fmt.Errorf(
			"search_id: %d, timezone: %s, time_unit: %s, time: %s, size: %d, monthday: %v, weekday: %v",
			d.Get("search_id").(int),
			d.Get("timezone").(string),
			d.Get("window_time_unit").(string),
			d.Get("window_time").(string),
			d.Get("window_size").(int),
			d.Get("window_monthday"),
			d.Get("window_weekday"),
		)
	}

	recurringExportResponse, _, err := client.RecurringExportsApi.CreateRecurringExport(context, &swagger.RecurringExportsApiCreateRecurringExportOpts{
		Body: optional.NewInterface(swagger.RecurringExportRequest{
			RecurringExport: &swagger.RecurringExport{
				SearchIds: []int32{
					int32(d.Get("search_id").(int)),
				},
				Timezone:       d.Get("timezone").(string),
				WindowTimeUnit: strings.ToUpper(d.Get("window_time_unit").(string)),
				WindowTime:     d.Get("window_time").(string),
				WindowSize:     int32(d.Get("window_size").(int)),
				WindowMonthday: int32(d.Get("window_monthday").(int)),
				WindowWeekday:  int32(d.Get("window_weekday").(int)),
			},
		}),
	})

	if err != nil {
		expandedError := err.(swagger.GenericSwaggerError)
		return fmt.Errorf("%s -> %s", err.Error(), string(expandedError.Body()))
	}

	d.SetId(fmt.Sprint(recurringExportResponse.RecurringExport.Id))

	r.Read(d, m)

	return nil
}

// Read - Get a recurring export
func (r RecurringExportResource) Read(d *schema.ResourceData, m interface{}) error {
	clientWithContext := m.(ClientWithContext)
	client := clientWithContext.Client
	context := clientWithContext.Context
	idInt64, err := strconv.ParseInt(d.Id(), 10, 32)
	idInt32 := int32(idInt64)
	if err != nil {
		return err
	}

	showExportResponse, _, err := client.RecurringExportsApi.ShowRecurringExport(context, idInt32)
	if err != nil {
		return err
	}

	recurringExport := showExportResponse.RecurringExport

	d.Set("timezone", recurringExport.Timezone)
	d.Set("search_id", recurringExport.Searches[0].Id)
	d.Set("window_time_unit", strings.ToLower(recurringExport.WindowTimeUnit))
	d.Set("window_time", recurringExport.WindowTime[0:8])
	d.Set("window_size", recurringExport.WindowSize)
	d.Set("window_monthday", recurringExport.WindowMonthday)
	d.Set("window_weekday", recurringExport.WindowWeekday)

	return nil
}

// Update - Update a recurring export
func (r RecurringExportResource) Update(d *schema.ResourceData, m interface{}) error {
	err := r.Delete(d, m)
	if err != nil {
		return err
	}
	return r.Create(d, m)
}

// Delete - Delete a recurring export
func (r RecurringExportResource) Delete(d *schema.ResourceData, m interface{}) error {
	clientWithContext := m.(ClientWithContext)
	client := clientWithContext.Client
	context := clientWithContext.Context
	idInt64, err := strconv.ParseInt(d.Id(), 10, 32)
	idInt32 := int32(idInt64)
	if err != nil {
		return err
	}

	_, err = client.RecurringExportsApi.DeleteRecurringExport(context, idInt32)
	if err != nil {
		return err
	}

	// Remove the reference to show it as deleted
	d.SetId("")

	return nil
}

// windowTimeUnits List of acceptable window time units
func (r RecurringExportResource) windowTimeUnits() []string {
	return []string{
		"day",
		"week",
		"month",
	}
}
