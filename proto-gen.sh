protoc -I=./internal/protofiles/ --go_out=plugins=grpc:. ./internal/protofiles/messages.proto
protoc -I=./internal/protofiles/ --go_out=plugins=grpc:. ./internal/protofiles/services.proto