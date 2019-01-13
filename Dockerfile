FROM golang as source
#FROM golang:1.11-alpine as builder

#RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /dcc
RUN cd /dcc && go build

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest
COPY --from=source /dcc/dcc /usr/local/bin/

EXPOSE 2510
ENTRYPOINT ["dcc"]
