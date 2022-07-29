# base go image
FROM golang:1.18-alpine as server

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o server .

RUN chmod +x /app/server

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=server /app/server /app

CMD [ "/app/server" ]