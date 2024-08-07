FROM golang:alpine as build-env
RUN apk add --no-cache git gcc tree
RUN mkdir /app
WORKDIR /app
COPY ./services ./services
COPY go.mod .
COPY go.sum .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./backend-server ./services/backend-server/cmd

FROM alpine
COPY --from=build-env /app/backend-server .
COPY config.yaml .
ENTRYPOINT ["./backend-server"]