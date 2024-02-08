FROM golang:1.22.0-alpine as build
WORKDIR /workspace
COPY . .
RUN go mod download -x && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

FROM alpine:3.19.0 as run
WORKDIR /workspace
COPY --from=build /workspace/server .
CMD ./server
EXPOSE 8080
