###########BUILD APP###################
FROM golang:1.17.7-alpine3.15 as builder
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
WORKDIR /go/src/work
COPY . /go/src/work/
RUN go get -u -v -f all
RUN go build -o app

###########START APP###################
FROM alpine:3.9 as prod
COPY --from=builder /go/src/work/app /
CMD ./app