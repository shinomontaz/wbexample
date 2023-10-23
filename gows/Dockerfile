FROM golang:alpine AS build

RUN apk add git

RUN mkdir /src
ADD . /src
WORKDIR /src

RUN go build -o /tmp/ws-server

FROM alpine:edge

COPY --from=build /tmp/ws-server /sbin/ws-server

#CMD /sbin/ws-server
ENTRYPOINT ["/sbin/ws-server"]
