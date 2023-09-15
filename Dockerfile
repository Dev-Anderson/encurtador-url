FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go build -o cmd/encurtador-url

FROM alpine:latest

RUN apk add --no-cache postgresql-client

COPY --from=builder /app/myapp /app/myapp

ENV PGHOST=localhost
ENV PGUSER=postgres
ENV PGPASSWORD=postgres
ENV PGDATABASE=encurtador-url
ENV PORT_SERVER=9001
ENV RANDOM_CARACTERES=0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ

EXPOSE 9001

CMD ["/app/encurtador-url"]
