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
			//ValidateFunc: r.validateQuery(), List/Set validation is not supported in terraform yet
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"keyword": {
						Type:        schema.TypeSet,
						Optional:    true,
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

					"combined": {
						Type:        schema.TypeSet,
						Optional:    true,
						Description: "Sets up a combined query",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"all_searches": {
									Type:        schema.TypeList,
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Optional:    true,
									Description: "A list of searches that should all match",
								},
								"any_searches": {
									Type:        schema.TypeList,
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Optional:    true,
									Description: "A list of searches where at least one should match",
								},
								"not_searches": {
									Type:        schema.TypeList,
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Optional:    true,
									Description: "A list of searches that shouldn't match",
								},
							},
						},
					},

					"boolean": {
						Type:        schema.TypeSet,
						Optional:    true,
						Description: "Sets up a boolean query",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"case_sensitivity": {
									Type:        schema.TypeString,
									Optional:    true,
									Default:     "hybrid",
									Description: "Set the type of case sensitivity",
								},
								"boolean": {
									Type:        schema.TypeString,
									Elem:        &schema.Schema{Type: schema.TypeString},
									Required:    true,
									Description: "A boolean query as a string",
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
	validateFunc := r.validateQuery()
	_, errs := validateFunc(d.Get("query"), "query")
	if len(errs) > 0 {
		return errs[0]
	}

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

	// If we have an id then ensure the state is updated from the API
	r.Read(d, m)

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

	//TODO: This query section still isn't working correctly

	// Handle the query set and query types
	query := [][]map[string]interface{}{}

	if search.Query.Type == "keyword" {
		keywordList := []map[string]interface{}{}

		keyword := map[string]interface{}{}
		keyword["case_sensitivity"] = search.Query.CaseSensitivity
		keyword["all_keywords"] = search.Query.AllKeywords
		keyword["any_keywords"] = search.Query.AnyKeywords
		keyword["not_keywords"] = search.Query.NotKeywords
		keywordList = append(keywordList, keyword)
		query = append(query, keywordList)
	}

	if search.Query.Type == "combined" {
		combinedList := []map[string]interface{}{}
		combined := map[string]interface{}{}
		combined["all_searches"] = search.Query.AllSearches
		combined["any_searches"] = search.Query.AnySearches
		combined["not_searches"] = search.Query.NotSearches
		combinedList = append(combinedList, combined)
		query = append(query, combinedList)
	}

	if search.Query.Type == "boolean" {
		booleanList := []map[string]interface{}{}
		boolean := map[string]interface{}{}
		boolean["case_sensitivity"] = search.Query.CaseSensitivity
		boolean["boolean"] = search.Query.Boolean
		booleanList = append(booleanList, boolean)
		query = append(query, booleanList)
	}

	d.Set("query", query)

	return nil
}

// Update - Update a recurring export
func (r SearchResource) Update(d *schema.ResourceData, m interface{}) error {
	validateFunc := r.validateQuery()
	_, errs := validateFunc(d.Get("query"), "query")
	if len(errs) > 0 {
		return errs[0]
	}

	clientWithContext := m.(ClientWithContext)
	client := clientWithContext.Client
	context := clientWithContext.Context
	idInt64, err := strconv.ParseInt(d.Id(), 10, 32)
	if err != nil {
		return fmt.Errorf("ID: %s:%d -> %s", d.Id(), idInt64, err.Error())
	}

	searchRequest, err := r.resourceDataToSearchRequest(d)
	if err != nil {
		return err
	}

	// If we have an id then ensure the state is updated from the API
	r.Read(d, m)

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
	validateFunc := r.validateQuery()
	_, errs := validateFunc(d.Get("query"), "query")
	if len(errs) > 0 {
		return errs[0]
	}

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
	query := querySet.List()[0].(map[string]interface{})
	searchRequest := swagger.SingleSearchRequest{}
	searchRequestQuery := &swagger.Query{
		OneOfQuery: swagger.OneOfQuery{},
	}

	// Keep track of queries
	queryCount := 0

	// Check for keyword query existing
	if keywordVal, ok := query["keyword"]; ok {
		keywordSet := keywordVal.(*schema.Set)

		// We're adding up all the queries to make sure by the end that we only have one
		queryCount += keywordSet.Len()

		// Require only one keyword set
		if keywordSet.Len() == 1 {
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

			searchRequestQuery.Type = "keyword"
			searchRequestQuery.CaseSensitivity = keyword["case_sensitivity"].(string)
			searchRequestQuery.OneOfQuery.KeywordQuery = swagger.KeywordQuery{
				AllKeywords: allKeywords,
				AnyKeywords: anyKeywords,
				NotKeywords: notKeywords,
			}
		}
	}

	// Check for combined query existing
	if combinedVal, ok := query["combined"]; ok {
		combinedSet := combinedVal.(*schema.Set)

		// We're adding up all the queries to make sure by the end that we only have one
		queryCount += combinedSet.Len()

		// Require only one keyword set
		if combinedSet.Len() == 1 {
			combined := combinedSet.List()[0].(map[string]interface{})

			allSearches := []int64{}
			for _, searchId := range combined["all_searches"].([]interface{}) {
				allSearches = append(allSearches, int64(searchId.(int)))
			}
			anySearches := []int64{}
			for _, searchId := range combined["any_searches"].([]interface{}) {
				anySearches = append(anySearches, int64(searchId.(int)))
			}
			notSearches := []int64{}
			for _, searchId := range combined["not_searches"].([]interface{}) {
				notSearches = append(notSearches, int64(searchId.(int)))
			}

			// Check any, all and not have atleast one element between them
			if len(allSearches)+len(anySearches)+len(notSearches) == 0 {
				return searchRequest, errors.New("you need at least one all, any or not search ID")
			}

			searchRequestQuery.Type = "combined"
			searchRequestQuery.OneOfQuery.CombinedQuery = swagger.CombinedQuery{
				AllSearches: allSearches,
				AnySearches: anySearches,
				NotSearches: notSearches,
			}
		}
	}

	// Check for boolean query existing
	if booleanVal, ok := query["boolean"]; ok {
		booleanSet := booleanVal.(*schema.Set)

		// We're adding up all the queries to make sure by the end that we only have one
		queryCount += booleanSet.Len()

		// Require only one keyword set
		if booleanSet.Len() == 1 {
			boolean := booleanSet.List()[0].(map[string]interface{})
			booleanString := boolean["boolean"].(string)
			if len(booleanString) == 0 {
				return searchRequest, errors.New("you need to supply a boolean string")
			}

			searchRequestQuery.Type = "boolean"
			searchRequestQuery.CaseSensitivity = boolean["case_sensitivity"].(string)
			searchRequestQuery.OneOfQuery.BooleanQuery = swagger.BooleanQuery{
				Boolean: booleanString,
			}
		}
	}

	// Check there is only one query type present
	if queryCount == 0 {
		return searchRequest, errors.New("one of keyword, combined or boolean must be set within query")
	}
	if queryCount != 1 {
		return searchRequest, errors.New("one of keyword, combined or boolean can exist within query")
	}

	return swagger.SingleSearchRequest{
		Search: &swagger.SearchRequestV2{
			Type_:    d.Get("type").(string),
			Category: d.Get("category").(string),
			Name:     d.Get("name").(string),
			Query:    searchRequestQuery,
		},
	}, nil
}

// validateQuery a custom validation method to check that the query set contains
// at least one of keyword, combined or boolean
// this is not supported yet :(
func (r SearchResource) validateQuery() schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		querySet := i.(*schema.Set)
		// Require only one query set
		if querySet.Len() != 1 {
			es = append(es, fmt.Errorf("exactly one query block is required here, got %d", querySet.Len()))
		}

		query := querySet.List()[0].(map[string]interface{})

		// Keep track of queries
		queryCount := 0

		// Check for keyword query
		if val, ok := query["keyword"]; ok {
			set := val.(*schema.Set)
			queryCount += set.Len()
			if set.Len() > 1 {
				es = append(es, fmt.Errorf("you can't have more than one keyword block here, got %d", set.Len()))
			}
		}

		// Check for combined query
		if val, ok := query["combined"]; ok {
			set := val.(*schema.Set)
			queryCount += set.Len()
			if set.Len() > 1 {
				es = append(es, fmt.Errorf("you can't have more than one combined block here, got %d", set.Len()))
			}
		}

		// Check for combined query
		if val, ok := query["boolean"]; ok {
			set := val.(*schema.Set)
			queryCount += set.Len()
			if set.Len() > 1 {
				es = append(es, fmt.Errorf("you can't have more than one boolean block here, got %d", set.Len()))
			}
		}

		// Check overall count
		if queryCount == 0 {
			es = append(es, fmt.Errorf("one of keyword, combined or boolean must be set within query, got %d", queryCount))
		}
		if queryCount != 1 {
			es = append(es, fmt.Errorf("one of keyword, combined or boolean can exist within query, got %d", queryCount))
		}

		return
	}
}
