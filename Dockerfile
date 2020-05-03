FROM golang:1.12.4 as build-env

WORKDIR /go/src/app

# RUN go get -d -v ./...
RUN go get gopkg.in/yaml.v2
RUN go get github.com/kelseyhightower/envconfig
ADD ./*.go /go/src/app
ADD src/ ..
RUN go generate ./...
RUN go build -o /go/bin/traveller

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/traveller /
ADD config.yml /
CMD ["/traveller"]
