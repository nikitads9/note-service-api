.PHONY: generate
generate:
	protoc -I api/note_v1/ --go_out=pkg/note_api --go_opt=paths=import --go-grpc_out=pkg/note_api --go-grpc_opt=paths=import api/note_v1/note_v1.proto
	mv pkg/note_api/github.com/nikitads9/note-service-api/pkg/note_api/* pkg/note_api/
	rm -r  pkg/note_api/github.com