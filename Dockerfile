FROM golang:alpine as build
COPY . /app
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/main main.go

FROM alpine
COPY --from=build /app/bin/main /main
EXPOSE 8000
ENTRYPOINT ["/main"]
