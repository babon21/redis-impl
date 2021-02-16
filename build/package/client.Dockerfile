# Builder
FROM golang:1.15-alpine3.12 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .

RUN make client

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 9090

COPY --from=builder /app/client /app

CMD /app/client
