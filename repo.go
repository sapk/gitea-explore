//go:generate go run -mod=vendor github.com/go-swagger/go-swagger/cmd/swagger generate client -f ./gitea.swagger.json
//go:generate go run -mod=vendor github.com/go-swagger/go-swagger/cmd/swagger generate spec -o ./assets/swagger/swagger.v1.json

package main

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/rs/zerolog"

	"gitea.com/sapk/explore/client"
	"gitea.com/sapk/explore/client/repository"
	"gitea.com/sapk/explore/models"
)

// Repository represents a repository
type Repository struct {
	// the source host
	Host string `json:"host"`

	// archived
	Archived bool `json:"archived,omitempty"`

	// avatar URL
	AvatarURL string `json:"avatar_url,omitempty"`

	// clone URL
	CloneURL string `json:"clone_url,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// empty
	Empty bool `json:"empty,omitempty"`

	// fork
	Fork bool `json:"fork,omitempty"`

	// forks
	Forks int64 `json:"forks_count,omitempty"`

	// full name
	FullName string `json:"full_name,omitempty"`

	// HTML URL //TODO deduplicate by ID
	HTMLURL string `json:"html_url,omitempty"`

	// mirror
	Mirror bool `json:"mirror,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// open issues
	OpenIssues int64 `json:"open_issues_count,omitempty"`

	// original URL
	OriginalURL string `json:"original_url,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// stars
	Stars int64 `json:"stars_count,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated_at,omitempty"`

	// Fetched when retrieved by api
	Fetched time.Time `json:"fetched_at,omitempty"`

	// watchers
	Watchers int64 `json:"watchers_count,omitempty"`
}

func parseServer(mainlog *zerolog.Logger, host string) ([]*Repository, string, error) {
	log := mainlog.With().Str("host", host).Logger()
	//log.Info().Msgf("Parsing: %s", host)
	remote := client.TransportConfig{
		Host:     host,
		BasePath: "/api/v1",
		Schemes:  []string{"http", "https"},
	}
	clt := client.NewHTTPClientWithConfig(nil, &remote)

	version, err := clt.Miscellaneous.GetVersion(nil, nil)
	if err != nil {
		log.Warn().Err(err).Msgf("GetVersion Error")
		return nil, "", err
	}
	log.Info().Msgf("Version: %s", version.GetPayload().Version)

	repos := make([]*models.Repository, 0)
	page := int64(1)
	limit := int64(10)

	for {
		params := repository.NewRepoSearchParams()
		params.SetPage(&page)
		params.SetLimit(&limit)
		search, err := clt.Repository.RepoSearch(params, nil)
		if err != nil {
			log.Warn().Msgf("Error: %v", err)
			break //TODO retry later but now just skipping to next
		}
		log.Info().Int64("page", page).Int("count", len(search.GetPayload().Data)).Msgf("Found: %v", search.GetPayload().OK)
		if !search.GetPayload().OK || len(search.GetPayload().Data) == 0 {
			break
		}
		//log.Printf("Search: %+v", search.GetPayload().Data)
		for i, repo := range search.GetPayload().Data {
			log.Debug().Msgf("Repo[%d]: %s", page*limit+int64(i), repo.HTMLURL)
			//log.Printf("Repo[%d]: %#v", i, repo)
		}
		repos = append(repos, search.GetPayload().Data...)
		page++

		time.Sleep(waitBetweenRequest)
	}
	log.Info().Str("Host", host).Int("Count", len(repos)).Msg("End of parsing host")
	return toLimitedRepository(host, repos), version.GetPayload().Version, nil
}

func toLimitedRepository(host string, repos []*models.Repository) []*Repository {
	result := make([]*Repository, len(repos))
	//TODO set to default value unset (ex: mirror)
	for i, r := range repos {
		result[i] = &Repository{
			Host:        host,
			Archived:    r.Archived,
			AvatarURL:   r.AvatarURL,
			CloneURL:    r.CloneURL,
			Created:     r.Created,
			Description: r.Description,
			Empty:       r.Empty,
			Fork:        r.Fork,
			Forks:       r.Forks,
			FullName:    r.FullName,
			HTMLURL:     r.HTMLURL,
			Mirror:      r.Mirror,
			Name:        r.Name,
			OpenIssues:  r.OpenIssues,
			OriginalURL: r.OriginalURL,
			Size:        r.Size,
			Stars:       r.Stars,
			Updated:     r.Updated,
			Watchers:    r.Watchers,
			Fetched:     time.Now(),
		}
	}
	return result
}
