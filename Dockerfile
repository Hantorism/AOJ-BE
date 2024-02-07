FROM golang:1.22.0-alpine as build
WORKDIR /src
COPY . .
RUN go mod download -x && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/server .

FROM alpine:3.19.0 as run
COPY --from=build /bin/server /bin/
ENTRYPOINT [ "/bin/server" ]
EXPOSE 1323