FROM golang:1.17-alpine AS build-env
RUN apk --no-cache add build-base git gcc
ENV GO111MODULE=on 
ADD . /src
WORKDIR /src 
RUN go run github.com/magefile/mage -v go:build

FROM alpine
WORKDIR /app
COPY --from=build-env /src/bin/dnsbl /app/
CMD ["./dnsbl", "graphql"]
