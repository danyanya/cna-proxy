FROM golang:alpine as builder
WORKDIR /go/src/github.com/danyanya/cna-proxy
RUN apk add git
COPY *.go ./
RUN go get -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM centurylink/ca-certs
LABEL author="danyanya <danya.brain@gmail.com>"
COPY --from=builder /go/src/github.com/danyanya/cna-proxy/cna-proxy /cna-proxy
ENTRYPOINT ["/cna-proxy"]
