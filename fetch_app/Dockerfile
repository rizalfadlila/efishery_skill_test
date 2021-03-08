FROM golang:1.14-alpine AS build

RUN apk add --no-cache git gcc libc-dev curl bash

WORKDIR /build

ADD . .

RUN go get -tags musl -u github.com/swaggo/swag/cmd/swag
RUN swag i
RUN go get -tags musl -v ./...

RUN CGO_ENABLED=1 go build -tags musl -o fetch_app

FROM alpine

RUN apk add bash ca-certificates

WORKDIR /usr/local/bin

COPY --from=build /build/fetch_app .
COPY --from=build /build/.env .

RUN chmod +x fetch_app

EXPOSE 3030

CMD ["fetch_app"]
