FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main .

FROM alpine:latest
COPY --from=builder /app/main /app/main

EXPOSE 5173

CMD ["/cmd/url-shortener"]