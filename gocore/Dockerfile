FROM golang:alpine AS build

#RUN apk add git

RUN mkdir /src
ADD . /src
WORKDIR /src

RUN go build -o /tmp/go-core ./cmd/main.go

FROM alpine:edge

COPY --from=build /tmp/go-core /sbin/go-core

CMD /sbin/go-core
#ENTRYPOINT ["/sbin/ws-server"]
