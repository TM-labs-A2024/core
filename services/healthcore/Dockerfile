FROM golang:alpine as build-env
RUN apk add --no-cache git gcc
RUN mkdir /app
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./healthcore ./cmd

FROM alpine
COPY --from=build-env /app/healthcore .
ENTRYPOINT ["./healthcore"]