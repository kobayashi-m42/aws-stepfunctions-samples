FROM golang:1.13.5-alpine3.11 as build

WORKDIR /go/app

COPY . .

RUN set -x && \
  apk update && \
  apk add --no-cache git && \
  go build -o aws-stepfunctions-token-samples

FROM alpine:3.11

WORKDIR /app

COPY --from=build /go/app/aws-stepfunctions-token-samples .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/aws-stepfunctions-token-samples

CMD ["./aws-stepfunctions-token-samples"]
