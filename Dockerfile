FROM golang:1.13 as build
ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0
WORKDIR /usr/src/kohhiapi
COPY . .
RUN go build -o kohhiapi main.go

FROM alpine:3.11.5
COPY --from=build /usr/src/kohhiapi/kohhiapi /kohhiapi
RUN apk --no-cache add ca-certificates
ENTRYPOINT [ "/kohhiapi" ]