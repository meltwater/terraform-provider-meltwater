package meltwater

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	paypalSdk "github.com/plutov/paypal/v4"

	"log"
)

type CatalogProductResource struct{}

func (r CatalogProductResource) Resource() *schema.Resource {
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

func (r CatalogProductResource) Schema() map[string]*schema.Schema {

	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the product",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The description of the product",
		},
		"image_url": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An externally hosted image of the product",
		},
		"home_url": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A URL to product information",
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(r.productTypes(), true),
			Description:  fmt.Sprintf("A product type. One of: %s", strings.Join(r.productTypes(), ",")),
		},
		"category": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A product category from the following list: https://developer.paypal.com/docs/api/catalog-products/v1/#definition-product_category",
		},
	}
}

// Create - Creating a catalog product in Paypal
func (r CatalogProductResource) Create(d *schema.ResourceData, m interface{}) error {
	client := m.(*paypalSdk.Client)
	product, err := client.CreateProduct(context.Background(), paypalSdk.Product{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		ImageUrl:    d.Get("image_url").(string),
		HomeUrl:     d.Get("home_url").(string),
		Type:        paypalSdk.ProductType(strings.ToUpper(d.Get("type").(string))),
		Category:    paypalSdk.ProductCategory(strings.ToUpper(d.Get("category").(string))),
	})
	if err != nil {
		log.Printf("Error creating catalog product : %s", err.Error())
		return err
	}

	d.SetId(product.ID)
	d.Set("name", product.Name)
	d.Set("description", product.Description)
	d.Set("image_url", product.ImageUrl)
	d.Set("home_url", product.HomeUrl)
	d.Set("type", string(product.Type))
	d.Set("category", string(product.Category))

	log.Printf("Created catalog product with ID: %s", product.ID)

	return nil
}

// Read - Get a catalog product in Paypal
func (r CatalogProductResource) Read(d *schema.ResourceData, m interface{}) error {
	client := m.(*paypalSdk.Client)

	product, err := client.GetProduct(context.Background(), d.Id())
	if err != nil {
		log.Printf("Error getting catalog product %s: %s", d.Id(), err.Error())
		return err
	}

	d.Set("name", product.Name)
	d.Set("description", product.Description)
	d.Set("image_url", product.ImageUrl)
	d.Set("home_url", product.HomeUrl)
	d.Set("type", string(product.Type))
	d.Set("category", string(product.Category))

	return nil
}

// Update - Update a catalog product in Paypal
func (r CatalogProductResource) Update(d *schema.ResourceData, m interface{}) error {
	client := m.(*paypalSdk.Client)

	product := paypalSdk.Product{
		ID:          d.Id(),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		ImageUrl:    d.Get("image_url").(string),
		HomeUrl:     d.Get("home_url").(string),
		Type:        paypalSdk.ProductType(strings.ToUpper(d.Get("type").(string))),
		Category:    paypalSdk.ProductCategory(strings.ToUpper(d.Get("category").(string))),
	}

	err := client.UpdateProduct(context.Background(), product)
	if err != nil {
		log.Printf("Error updating catalog product %s: %s", d.Id(), err.Error())
		return err
	}

	return r.Read(d, m)
}

// Delete - Delete the a catalog product in Paypal - Products cannot be deleted
// so we will update the name with a DELETED suffix and remove our reference to it
func (r CatalogProductResource) Delete(d *schema.ResourceData, m interface{}) error {
	originalName := d.Get("name").(string)
	name := fmt.Sprintf("%s (removed)", d.Get("name").(string))
	d.Set("name", name)
	err := r.Update(d, m)
	if err != nil {
		log.Printf("Error updating catalog product with deleted name %s: %s", d.Id(), err.Error())
		d.Set("name", originalName)
		return err
	}

	// Even though we can't delete it, we can remove our id reference
	d.SetId("")

	return nil
}

// productTypes List of acceptable product types
func (r CatalogProductResource) productTypes() []string {
	return []string{
		strings.ToLower(string(paypalSdk.ProductTypePhysical)),
		strings.ToLower(string(paypalSdk.ProductTypeDigital)),
		strings.ToLower(string(paypalSdk.ProductTypeService)),
	}
}
