FROM golang:alpine3.19 AS build

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN go build ./cmd/compiler

FROM docker:26.1.0-dind

WORKDIR /

COPY --from=build /app/compiler compiler

CMD ["/compiler"]