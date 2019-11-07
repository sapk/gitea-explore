// Package main Explore API.
//
// This documentation describes the Explore Gitea API.
//
//     Schemes: http, https
//     BasePath: /api
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/olivere/elastic/v7"
	"github.com/rs/zerolog"
)

// SearchResultList represent a list of search api result response
// swagger:response SearchResultList
type SearchResultList struct {
	// The result list
	// in: body
	Body SearchResult
}

// ServerList represent a list of server response
// swagger:response ServerList
type ServerList struct {
	// The site list
	// in: body
	Body map[string]*Server
}

//SearchResult respons object
type SearchResult struct {
	Page         int
	Total        int64
	Repositories []Repository
}

func api(webLogger *zerolog.Logger, es *elastic.Client) func(r chi.Router) {
	return func(r chi.Router) {
		//TODO remove only for dev
		// Basic CORS
		// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
		cors := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET"},
			AllowedHeaders: []string{"Accept", "Content-Type", "X-CSRF-Token"},
		})
		r.Use(cors.Handler)

		r.Get("/v1/servers", func(w http.ResponseWriter, r *http.Request) {
			// swagger:operation GET /v1/servers servers
			// ---
			// summary: List the servers indexed
			// produces:
			// - application/json
			// responses:
			//   "500":
			//     description: Internal Server Error
			//   "200":
			//     "$ref": "#/responses/ServerList"

			//TODO store favicon of site
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(serverList)
			if err != nil {
				webLogger.Warn().Err(err).Msg("Fail to send site list")
			}
		})
		r.Get("/v1/search", func(w http.ResponseWriter, r *http.Request) {
			// swagger:operation GET /v1/search search
			// ---
			// summary: List the repositories found by search
			// produces:
			// - application/json
			// parameters:
			// - name: q
			//   in: query
			//   description: the requested query
			//   type: string
			//   required: true
			// - name: order_by
			//   in: query
			//   description: Field use to order results (default to stars_count)
			//   type: string
			//   required: false
			// - name: order
			//   in: query
			//   description: Order to use (default to desc)
			//   type: string
			//   required: false
			//   enum: ["asc", "desc"]
			// - name: page
			//   in: query
			//   description: Current page to return
			//   type: integer
			//   required: false
			// responses:
			//   "500":
			//     description: Internal Server Error
			//   "200":
			//     "$ref": "#/responses/SearchResultList"

			//TODO
			//   enum: [stars_count]

			q := r.URL.Query().Get("q")
			orderBy := r.URL.Query().Get("order_by")
			if orderBy == "" { //TODO better by checking a list of field
				orderBy = "stars_count"
			}
			order := r.URL.Query().Get("order")
			orderBool := false //default
			if order == "asc" {
				orderBool = true
			}
			page := r.URL.Query().Get("page")
			pageInt := 1
			if n, err := strconv.Atoi(page); err == nil && n >= 1 {
				pageInt = n
			}

			termQuery := elastic.NewSimpleQueryStringQuery(q) //TODO support filter
			termQuery.Lenient(true)
			//termQuery.FuzzyPrefixLength(1)
			//termQuery.AnalyzeWildcard(true)

			termQuery.FieldWithBoost("name", 4)
			termQuery.FieldWithBoost("full_name", 3)
			termQuery.FieldWithBoost("description", 2)
			termQuery.FieldWithBoost("host", 1)

			searchResult, err := es.Search().
				Index(esIndex).                    // search in index "tweets"
				Query(termQuery).                  // specify the query
				Sort(orderBy, orderBool).          // TODO -> this seems to not work properly
				From((pageInt - 1) * 10).Size(10). // take documents 0-9
				//Pretty(true).                         // pretty print request and response JSON
				Do(context.Background()) // execute
			if err != nil {
				webLogger.Warn().Err(err).Str("q", q).Msg("Fail to query index")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			//webLogger.Debug().Int64("response_time", searchResult.TookInMillis).Msg("Query result") //TODO log query details

			w.Header().Set("Content-Type", "application/json")
			repoType := reflect.TypeOf(Repository{})
			repos := make([]Repository, len(searchResult.Hits.Hits))
			for i, objInt := range searchResult.Each(repoType) {
				repos[i] = objInt.(Repository)
			}
			response := SearchResult{
				Page:         pageInt,
				Total:        searchResult.TotalHits(),
				Repositories: repos,
			}
			err = json.NewEncoder(w).Encode(response) //Use deafualt json renderer
			//TODO ouput
			if err != nil {
				webLogger.Warn().Err(err).Msg("Fail to encode query result")
			}
		})
	}
}
