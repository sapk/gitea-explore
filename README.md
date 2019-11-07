# Explore

Use mage helper to run tasks:
```
go run -mod=vendor mage.go -l
go run -mod=vendor mage.go build
 -- OR --
alias m="go run -mod=vendor mage.go"
m -l
```


## Start dev ES
```
docker run -d --rm --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.3.1
```

## Start with DejaVu for debug
```
docker run -d --rm --name elasticsearch -p 9200:9200 -p 9300:9300 -e discovery.type=single-node -e http.cors.enabled=true -e http.cors.allow-origin=http://localhost:1358,http://127.0.0.1:1358 -e http.cors.allow-headers=X-Requested-With,X-Auth-Token,Content-Type,Content-Length,Authorization -e http.cors.allow-credentials=true docker.elastic.co/elasticsearch/elasticsearch-oss:7.3.1
docker run -d --rm --name dejavu -p 1358:1358 appbaseio/dejavu
```

## Clear dev docker 
```
docker stop dejavu elasticsearch && docker rm dejavu elasticsearch 
```