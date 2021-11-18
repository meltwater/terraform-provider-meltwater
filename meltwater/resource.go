package meltwater

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// TerraformResource The terraform resource interface that all resources implement
type TerraformResource interface {
	Resource() *schema.Resource
	Schema() map[string]*schema.Schema
	Create(*schema.ResourceData, interface{}) error
	Read(*schema.ResourceData, interface{}) error
	Update(*schema.ResourceData, interface{}) error
	Delete(*schema.ResourceData, interface{}) error
}

// SchemaSimplified Simplified schema used for tests
type SchemaSimplified struct {
	Type     schema.ValueType
	Optional bool
	Required bool
	Nested   map[string]SchemaSimplified
}
