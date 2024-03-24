FROM golang:1.21.4 AS it-revolution-test-1-builder

RUN go version

COPY . /github.com/Woodfyn/it-revolution-test-1/
WORKDIR /github.com/Woodfyn/it-revolution-test-1/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=it-revolution-test-1-builder /github.com/Woodfyn/it-revolution-test-1/.bin/app .

COPY --from=it-revolution-test-1-builder /github.com/Woodfyn/it-revolution-test-1/configs/prod.yml ./configs/prod.yml

CMD ["./app"]