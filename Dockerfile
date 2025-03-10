FROM golang:1.23 AS buildenv

ARG BINARY=app

WORKDIR /app
COPY go.mod go.sum ./
COPY *.go ./
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $BINARY

FROM debian

ARG BINARY=app

COPY --from=buildenv /app/$BINARY /app/run

ENTRYPOINT [ "/app/run" ]