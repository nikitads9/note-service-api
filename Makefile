.PHONY: compile
compile:
	protoc -I api/note_v1/ --go_out=pkg/note_api --go_opt=paths=import --go-grpc_out=pkg/note_api --go-grpc_opt=paths=import api/note_v1/note_v1.proto
	mv "C:\Users\MSI GS75\Documents\GitHub\note-service-api\pkg\note_api\github.com\nikitads9\note-service-api\pkg\note_api\note_v1_grpc.pb.go" "C:\Users\MSI GS75\Documents\GitHub\note-service-api\pkg\note_api"
	mv "C:\Users\MSI GS75\Documents\GitHub\note-service-api\pkg\note_api\github.com\nikitads9\note-service-api\pkg\note_api\note_v1.pb.go" "C:\Users\MSI GS75\Documents\GitHub\note-service-api\pkg\note_api"
	rm -r  "C:\Users\MSI GS75\Documents\GitHub\note-service-api\pkg\note_api\github.com"