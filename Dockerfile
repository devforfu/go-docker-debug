FROM golang:1.10.1-alpine3.7 AS build-env
ENV CGO_ENABLED 0
COPY ./src /go/src
RUN apk add --no-cache git && \
    go get github.com/derekparker/delve/cmd/dlv && \
    go get github.com/gorilla/schema && \
    go build -gcflags "all=-N -l" -o /server hello && \
    apk del git

FROM alpine:3.7
EXPOSE 8080 40000
RUN apk add --no-cache libc6-compat
WORKDIR /
COPY --from=build-env /server /
COPY --from=build-env /go/bin/dlv /
#CMD ["/server"]
CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/server"]