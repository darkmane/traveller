FROM golang:1.12.4 as build-env

WORKDIR /go/src/app
ADD ./*.go /go/src/app
ADD src/ ..

RUN go get -d -v ./...

RUN go build -o /go/bin/traveller

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/traveller /
ADD config.yml /
CMD ["/traveller"]
