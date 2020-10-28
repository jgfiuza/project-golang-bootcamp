# build stage
FROM golang:1.15-alpine3.12 AS build-env
RUN apk --no-cache add build-base git mercurial gcc
ADD . /src
RUN cd /src/cmd/srv && go build -o qapi

# final stage
FROM alpine:3.12
WORKDIR /app
COPY --from=build-env /src/cmd/srv/qapi /app/
ENTRYPOINT ./qapi