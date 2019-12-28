FROM golang:1.12.4 as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN ls -Al

RUN go get -v ./...
#RUN go get -d -v ./...

RUN go build -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/app /
CMD ["/app"]
