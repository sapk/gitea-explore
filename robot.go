package main

import (
	"context"
	"time"

	"github.com/olivere/elastic/v7"
)

//TODO respect robot.txt
//TODO respect robot metadata
//TODO find a way for repo to opt-out/opt-in

func robot(debug bool, es *elastic.Client) {
	robotLogger := setupLogger(debug, "robot")
	robotLogger.Info().Msgf("Robot starting ...")
	for _, h := range hostList {
		repos, version, err := parseServer(robotLogger, h)
		if err == nil {
			serverList[h] = &Server{
				Host:      h,
				Version:   version,
				RepoCount: int64(len(repos)),
				//RepositoryList: repos,
				LastUpdate: time.Now(),
			}
			for _, repo := range repos {
				//TODO use bulk import
				_, err = es.Index().
					Index(esIndex).
					Type("repository").
					Id(repo.HTMLURL).
					BodyJson(repo).
					Refresh("wait_for").
					Do(context.Background())
				if err != nil {
					// Handle error
					robotLogger.Warn().Err(err).Str("repo", repo.FullName).Msg("Fail to insert in index")
				}

				//TODO purge old repo on Fetched
				//TODO get topics https://github.com/go-gitea/gitea/pull/7963
				//TODO get average response time for stats
				//TODO log number of click on a repo to order them (interest)
			}
			robotLogger.Info().Str("Host", h).Int("Count", len(repos)).Msg("End of storing host")
		}
		time.Sleep(waitBetweenServer)
	}
}
