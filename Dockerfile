FROM golang:1.19.2 as build-env

WORKDIR /go/src/app

ADD ./*.go /go/src/app
ADD . ..
RUN go mod tidy
RUN go generate ./...
RUN go build -o /go/bin/traveller

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/traveller /
ADD config.yml /
CMD ["/traveller"]
