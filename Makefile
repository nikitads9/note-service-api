.PHONY: generate
generate:
	protoc -I api/note_v1 -I proto --go_out=pkg/note_api --go_opt=paths=import --go-grpc_out=pkg/note_api --go-grpc_opt=paths=import --grpc-gateway_out=pkg/note_api --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=import api/note_v1/note_v1.proto --validate_out lang=go:pkg/note_api
	mv pkg/note_api/github.com/nikitads9/note-service-api/pkg/note_api/* pkg/note_api/
	rm -r  ./pkg/note_api/github.com
		
.PHONY: build
build: vendor-proto .generate .build

PHONY: .build
#-I api/note_v1/
.build:
		CGO_ENABLED=0 GOOS=linux go build -o bin/note-service-api cmd/grpc-server/note_service.go
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
		mkdir -p proto
		cp api/note_v1/note_v1.proto proto/
		@if [ ! -d proto/google ]; then \
			git clone https://github.com/googleapis/googleapis proto/googleapis &&\
			mkdir -p  proto/google/ &&\
			mv proto/googleapis/google/api proto/google &&\
			rm -rf proto/googleapis ;\
		fi
		@if [ ! -d proto/github.com/envoyproxy ]; then \
			mkdir -p proto/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate proto/protoc-gen-validate &&\
			mv proto/protoc-gen-validate/validate/*.proto proto/validate &&\
			rm -r -f proto/protoc-gen-validate ;\
		fi
		@if [ ! -d proto/google/protobuf ]; then\
			git clone https://github.com/protocolbuffers/protobuf proto/protobuf &&\
			mkdir -p  proto/google/protobuf &&\
			mv proto/protobuf/src/google/protobuf/*.proto proto/google/protobuf &&\
			rm -rf proto/protobuf ;\
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
		go get -u github.com/envoyproxy/protoc-gen-validate
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