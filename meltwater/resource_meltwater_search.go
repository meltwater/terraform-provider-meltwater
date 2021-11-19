package meltwater

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/meltwater/terraform-provider-meltwater/swagger"
)

type SearchResource struct{}

func (r SearchResource) Resource() *schema.Resource {
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

func (r SearchResource) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(r.searchTypes(), true),
			Description:  fmt.Sprintf("A search type must be one of: %s", strings.Join(r.searchTypes(), ",")),
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "A name for the search",
		},
		"category": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(r.searchCategories(), true),
			Description:  fmt.Sprintf("A search category must be one of: %s", strings.Join(r.searchCategories(), ",")),
		},
		"query": {
			Type:        schema.TypeSet,
			Required:    true,
			Description: "The container for the search query. Must include one of _TBC_",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"keyword": {
						Type:        schema.TypeSet,
						Required:    true,
						Description: "Sets up a keyword query",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"case_sensitivity": {
									Type:        schema.TypeString,
									Optional:    true,
									Default:     "hybrid",
									Description: "Set the type of case sensitivity",
								},
								"all_keywords": {
									Type:        schema.TypeList,
									Elem:        &schema.Schema{Type: schema.TypeString},
									Optional:    true,
									Description: "A list of keywords that should all be present",
								},
								"any_keywords": {
									Type:        schema.TypeList,
									Elem:        &schema.Schema{Type: schema.TypeString},
									Optional:    true,
									Description: "A list of keywords where at least one is present",
								},
								"not_keywords": {
									Type:        schema.TypeList,
									Elem:        &schema.Schema{Type: schema.TypeString},
									Optional:    true,
									Description: "A list of keywords that shouldn't be present",
								},
							},
						},
					},
				},
			},
		},
	}
}

// Create - Create a search in the Meltwater API
func (r SearchResource) Create(d *schema.ResourceData, m interface{}) error {
	clientWithContext := m.(ClientWithContext)
	client := clientWithContext.Client
	context := clientWithContext.Context

	searchRequest, err := r.resourceDataToSearchRequest(d)
	if err != nil {
		expandedError := err.(swagger.GenericSwaggerError)
		return fmt.Errorf("%s -> %s", err.Error(), string(expandedError.Body()))
	}

	searchResponse, _, err := client.SearchesApi.CreateSearch(
		context,
		searchRequest,
		&swagger.SearchesApiCreateSearchOpts{},
	)

	if err != nil {
		expandedError := err.(swagger.GenericSwaggerError)
		return fmt.Errorf("%s -> %s", err.Error(), string(expandedError.Body()))
	}

	d.SetId(fmt.Sprint(searchResponse.Search.Id))

	r.Read(d, m)

	return nil
}

// Read - Get a recurring export
func (r SearchResource) Read(d *schema.ResourceData, m interface{}) error {
	clientWithContext := m.(ClientWithContext)
	client := clientWithContext.Client
	context := clientWithContext.Context
	idInt64, err := strconv.ParseInt(d.Id(), 10, 32)
	if err != nil {
		return fmt.Errorf("ID: %s:%d -> %s", d.Id(), idInt64, err.Error())
	}

	searchResponse, _, err := client.SearchesApi.GetSearch(context, idInt64, &swagger.SearchesApiGetSearchOpts{})
	if err != nil {
		return err
	}

	search := searchResponse.Search

	d.Set("type", search.Type_)
	d.Set("category", search.Category)
	d.Set("name", search.Name)

	//TODO: Add query stuff here

	return nil
}

// Update - Update a recurring export
func (r SearchResource) Update(d *schema.ResourceData, m interface{}) error {
	clientWithContext := m.(ClientWithContext)
	client := clientWithContext.Client
	context := clientWithContext.Context
	idInt64, err := strconv.ParseInt(d.Id(), 10, 32)
	if err != nil {
		return fmt.Errorf("ID: %s:%d -> %s", d.Id(), idInt64, err.Error())
	}

	searchRequest, err := r.resourceDataToSearchRequest(d)
	if err != nil {
		expandedError := err.(swagger.GenericSwaggerError)
		return fmt.Errorf("%s -> %s", err.Error(), string(expandedError.Body()))
	}

	searchResponse, _, err := client.SearchesApi.UpdateSearch(context, searchRequest, idInt64, &swagger.SearchesApiUpdateSearchOpts{})
	if err != nil {
		expandedError := err.(swagger.GenericSwaggerError)
		return fmt.Errorf("%s -> %s", err.Error(), string(expandedError.Body()))
	}

	d.SetId(fmt.Sprint(searchResponse.Search.Id))

	return r.Read(d, m)
}

// Delete - Delete a recurring export
func (r SearchResource) Delete(d *schema.ResourceData, m interface{}) error {
	clientWithContext := m.(ClientWithContext)
	client := clientWithContext.Client
	context := clientWithContext.Context
	idInt64, err := strconv.ParseInt(d.Id(), 10, 32)
	if err != nil {
		return err
	}

	_, err = client.SearchesApi.DeleteSearch(context, idInt64, &swagger.SearchesApiDeleteSearchOpts{})
	if err != nil {
		return err
	}

	// Remove the reference to show it as deleted
	d.SetId("")

	return nil
}

// searchTypes List of acceptable search types
func (r SearchResource) searchTypes() []string {
	return []string{
		"news",
		"social",
		"broadcast",
	}
}

// searchCategories List of acceptable search catgories
func (r SearchResource) searchCategories() []string {
	return []string{
		"mi",
		"explore",
	}
}

func (r SearchResource) resourceDataToSearchRequest(d *schema.ResourceData) (swagger.SingleSearchRequest, error) {
	querySet := d.Get("query").(*schema.Set)
	searchRequest := swagger.SingleSearchRequest{}

	// Require only one query set
	if querySet.Len() != 1 {
		return searchRequest, errors.New("exactly one query block is required here")
	}
	query := querySet.List()[0].(map[string]interface{})
	keywordSet := query["keyword"].(*schema.Set)

	// TODO: Check for other query types

	// Require only one keyword set
	if keywordSet.Len() != 1 {
		return searchRequest, errors.New("exactly one keyword block is required here")
	}
	keyword := keywordSet.List()[0].(map[string]interface{})

	allKeywords := []string{}
	for _, keywordItem := range keyword["all_keywords"].([]interface{}) {
		allKeywords = append(allKeywords, fmt.Sprintf("%v", keywordItem))
	}
	anyKeywords := []string{}
	for _, keywordItem := range keyword["any_keywords"].([]interface{}) {
		anyKeywords = append(anyKeywords, fmt.Sprintf("%v", keywordItem))
	}
	notKeywords := []string{}
	for _, keywordItem := range keyword["not_keywords"].([]interface{}) {
		notKeywords = append(notKeywords, fmt.Sprintf("%v", keywordItem))
	}

	// Check any, all and not have atleast one element between them
	if len(allKeywords)+len(anyKeywords)+len(notKeywords) == 0 {
		return searchRequest, errors.New("you need at least one all, any or not keyword")
	}

	return swagger.SingleSearchRequest{
		Search: &swagger.SearchRequestV2{
			Type_:    d.Get("type").(string),
			Category: d.Get("category").(string),
			Name:     d.Get("name").(string),
			Query: &swagger.Query{
				OneOfQuery: swagger.OneOfQuery{
					Type:            "keyword",
					CaseSensitivity: keyword["case_sensitivity"].(string),
					KeywordQuery: swagger.KeywordQuery{
						AllKeywords: allKeywords,
						AnyKeywords: anyKeywords,
						NotKeywords: notKeywords,
					},
				},
			},
		},
	}, nil
}
