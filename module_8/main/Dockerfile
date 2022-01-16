#FROM ubuntu
#ADD bin/amd64/httpserver /httpserver
#EXPOSE 8080
#ENTRYPOINT /httpserver

FROM golang:1.17.2-alpine3.14 as build
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/module_8  .

FROM alpine
COPY --from=build /app/bin/linux/amd64/module_8 /app/bin/linux/amd64/module_8
ENTRYPOINT ["/app/bin/linux/amd64/module_8"]