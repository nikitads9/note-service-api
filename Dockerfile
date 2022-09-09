FROM golang:1.18.3-alpine AS builder

COPY . /github.com/nikitads9/note-service-api/

WORKDIR /github.com/nikitads9/note-service-api/

RUN go mod download
RUN go build -o ./bin/note_service cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/nikitads9/note-service-api/bin/note-service .
COPY --from=builder /github.com/nikitads9/note-service-api/config.yml .

EXPOSE 50051
EXPOSE 8000

CMD ["./note_service"]