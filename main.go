package main

import (
	"context"
	"flag"
	"os"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/rs/zerolog"
)

const esIndex = "gitea"

//*
var waitBetweenRequest = 5 * time.Second
var waitBetweenServer = 15 * time.Minute

//*/
/*
var waitBetweenRequest = 2 * time.Second
var waitBetweenServer = 1 * time.Minute

//*/

//var hostList = []string{"gitea.com", "try.gitea.io", "codeberg.org", "git.passageenseine.fr", "git.teknik.io", "opendev.org"} //git.passageenseine.fr is too old
var hostList = []string{"gitea.com", "try.gitea.io", "codeberg.org", "git.teknik.io", "opendev.org", "code.antopie.org", "gitea.anfuchs.de", "git.koehlerweb.org", "git.jcg.re", "gitea.codi.coop", "git.daiko.fr", "git.sablun.org", "git.osuv.de", "gitea.vornet.cz", "git.luehne.de", "git.spip.net", "djib.fr"}
var serverList map[string]*Server

//TODO version check of remote

//Server Represent a server
type Server struct {
	Host      string
	Version   string
	RepoCount int64
	//RepositoryList []*Repository
	LastUpdate time.Time
}

func init() {
	serverList = make(map[string]*Server, len(hostList))
	//TODO load information from es at start
	for _, h := range hostList {
		serverList[h] = &Server{
			Host: h,
			//Version:   version,
			//RepoCount: len(repos),
			//LastUpdate: time.Now(),
		}
	}
}

func main() {
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	es := esClient(*debug)
	go robot(*debug, es)

	web(*debug, es)
}

func esClient(debug bool) *elastic.Client {
	logger := setupLogger(debug, "es")
	urls := strings.Split(os.Getenv("ES_URLS"), ",")
	if len(urls) == 1 && urls[0] == "" {
		urls[0] = "http://127.0.0.1:9200"
	}
	logger.Info().Strs("ES_URLS", urls).Msgf("ES connection starting ...")

	// Create a ES client
	errLog := logger.Level(zerolog.ErrorLevel)
	infoLog := logger.Level(zerolog.InfoLevel)
	debugLog := logger.Level(zerolog.DebugLevel)
	client, err := elastic.NewClient(
		elastic.SetURL(urls...),
		elastic.SetErrorLog(&errLog),
		elastic.SetInfoLog(&infoLog),
		elastic.SetTraceLog(&debugLog))
	if err != nil {
		//TODO Handle error
		panic(err)
	}
	exists, err := client.IndexExists(esIndex).Do(context.Background())
	if err != nil {
		//TODO Handle error
		panic(err)
	}
	if !exists {
		// Index does not exist yet.
		// Create an index
		_, err = client.CreateIndex(esIndex).Do(context.Background())
		if err != nil {
			//TODO Handle error
			panic(err)
		}
	}

	//load previous references
	for _, h := range hostList {
		termQuery := elastic.NewTermQuery("host", h) //TODO ony count
		searchResult, err := client.Search().
			Index(esIndex).          // search in index "tweets"
			Query(termQuery).        // specify the query
			Do(context.Background()) // execute
		if err != nil {
			logger.Warn().Err(err).Msg("Fail to query index at init")
		}
		serverList[h].RepoCount = searchResult.TotalHits()
		logger.Info().Str("Host", h).Int64("RepoCount", serverList[h].RepoCount).Msg("Init server at start-up")
	}

	//TODO log version
	return client
}
