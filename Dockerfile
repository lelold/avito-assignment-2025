FROM golang:1.23.5

WORKDIR ${GOPATH}/avito-assignment-2025/
COPY . ${GOPATH}/avito-assignment-2025/

RUN go build -o /build ./cmd/ \
    && go clean -cache -modcache

EXPOSE 8080

CMD ["/build"]