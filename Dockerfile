FROM        golang:1.8
MAINTAINER  Guilherme Henrique Piasson <guilhermehpiasson@gmail.com>

ENV     PORT  8080

WORKDIR     /go/src/gin-container
ADD         . /go/src/gin-container

RUN     go get github.com/tools/godep
RUN     go get github.com/gin-gonic/gin
RUN     go get github.com/go-redis/redis
RUN     go get gopkg.in/gorp.v1
RUN     go install github.com/tools/godep
RUN     go install github.com/gin-gonic/gin
RUN     go install github.com/go-redis/redis

EXPOSE 8080
ENTRYPOINT  ["/usr/local/bin/go"]
CMD     ["run", "main.go"]