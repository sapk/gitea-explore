###################################
#Build stage JS
FROM node AS build-env-js

COPY . /explore
WORKDIR /explore/assets/webapp/

RUN yarn install && yarn build
 
###################################
#Build stage Go
FROM golang:alpine AS build-env-go

ENV CGO_ENABLED=0

COPY . /explore
COPY --from=build-env-js /explore/assets/webapp/dist /explore/assets/webapp/dist
WORKDIR /explore


RUN go run -mod=vendor mage.go generate \
 && go run -mod=vendor mage.go build

###################################
#Running stage
#FROM scratch
FROM alpine
LABEL maintainer="antoine.girard@sapk.fr"

COPY --from=build-env-go /explore/explore /explore

CMD ["/explore"]
EXPOSE 3000
