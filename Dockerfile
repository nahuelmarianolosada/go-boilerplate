#TODO: Create a Dockerfile to run the application
FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
#TODO: Create a Dockerfile to github
RUN go get github.com/nahuelmarianolosada/go-boilerplate
RUN cd /build && git clone https://github.com/nahuelmarianolosada/go-boilerplate

RUN cd /build/go-boilerplate && go build

EXPOSE 8080

ENTRYPOINT [ "/build/go-boilerplate/cmd/server.go" ]

