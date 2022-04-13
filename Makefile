PHONY: generate
echo "Hello"

PHONY: test
echo "Hi!"
protoc -I api/note_v1/ --go_out=pkg/note_api --go_opt=paths=import --go-grpc_out=pkg/note_api --go-grpc_opt=paths=import api/note_v1/note_v1.proto