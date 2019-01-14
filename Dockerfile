FROM golang as sourcer
ADD . /go/src/github.com/fadeAce/dcc
RUN cd /go/src/github.com/fadeAce/dcc && go build
#FROM bashell/alpine-bash
#FROM golang
FROM ubuntu
COPY --from=sourcer /go/src/github.com/fadeAce/dcc/dcc /usr/local/bin/
EXPOSE 2510
#CMD ["bash /usr/local/bin/dcc.sh"]
CMD ["/bin/bash"]
ENTRYPOINT ["dcc"]
