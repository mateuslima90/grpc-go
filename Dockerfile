FROM golang:1.17-alpine AS builder

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o usersGrpc cmd/server/server.go

FROM alpine:latest AS production
RUN addgroup -S golang && adduser -SD golang -G golang
USER golang:golang
COPY --from=builder /app/usersGrpc .
CMD ["./usersGrpc"]
