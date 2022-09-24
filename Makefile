.PHONY: generate
generate:
	protoc -I api/note_v1 -I proto --go_out=pkg/note_api --go_opt=paths=import --go-grpc_out=pkg/note_api --go-grpc_opt=paths=import --grpc-gateway_out=pkg/note_api --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=import api/note_v1/note_v1.proto
	mv pkg/note_api/github.com/nikitads9/note-service-api/pkg/note_api/* pkg/note_api/
	rm -r  ./pkg/note_api/github.com


.PHONY: build
build: vendor-proto .generate .build

PHONY: .build
#-I api/note_v1/
.build:
		CGO_ENABLED=0 GOOS=linux go build -o bin/ocp-note-api cmd/grpc-server/main.go
		CGO_ENABLED=0 GOOS=linux go build -o bin/kafkaConsumer cmd/kafkaConsumer/consumer.go

PHONY: install
install: build .install

PHONY: .install
install:
		go install cmd/grpc-server/main.go

PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
	@if [ ! -d vendor.protogen/google ]; then \
      git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
      mkdir -p  vendor.protogen/google/ &&\
      mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
      rm -rf vendor.protogen/googleapis ;\
    fi
    @if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
      mkdir -p vendor.protogen/validate &&\
      git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
      mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
      rm -rf vendor.protogen/protoc-gen-validate ;\
    fi
.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/envoyproxy/protoc-gen-validate
		go get github.com/fullstorydev/grpcui/...
		
.PHONY: coverage
coverage:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"

PHONY: test-coverage
test-coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out