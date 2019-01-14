# build formal dcc sidecar image
FROM golang as sourcer
ADD . /go/src/github.com/fadeAce/dcc
RUN cd /go/src/github.com/fadeAce/dcc && go build

FROM ubuntu
COPY --from=sourcer /go/src/github.com/fadeAce/dcc/dcc /usr/local/bin/
EXPOSE 2510
CMD ["/bin/bash"]
ENTRYPOINT ["dcc"]