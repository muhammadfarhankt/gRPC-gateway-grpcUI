# gRPC-gateway-grpcUI

### Start using it

```bash
brew install protobuf
brew install grpcui

go mod init github.com/your-username/project-name 

sudo go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
sudo go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

go install \      
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

```

### Link to Documentation

https://github.com/grpc-ecosystem/grpc-gateway?tab=readme-ov-file
<br>
gRPC Server & UI : https://www.youtube.com/watch?v=zxCCYYkmDQU/ 
<br>
gRPC Gateway : https://www.youtube.com/watch?v=SawVP43DPpk/

### Generate proto files

```bash
protoc -I ./pkg/proto --go_out=./pkg/proto --go-grpc_out=./pkg/proto --grpc-gateway_out=. --go_opt paths=source_relative --go-grpc_opt paths=source_relative ./pkg/proto/intern.proto
```

## Run it

```bash
go run cmd/main.go
grpcui -plaintext localhost:9901
```

## gRPC to REST API Conversion

<img width="438" alt="image" src="https://github.com/muhammadfarhankt/gRPC-gateway-grpcUI/assets/50117098/f4483198-bf45-49f5-b37a-ff1b70b319db">

## gRPC UI

<img width="717" alt="image" src="https://github.com/muhammadfarhankt/gRPC-gateway-grpcUI/assets/50117098/f7091a95-8af2-44b5-9797-83c82c65deb6">

## Result

<img width="559" alt="image" src="https://github.com/muhammadfarhankt/gRPC-gateway-grpcUI/assets/50117098/8913ad3e-a682-4051-b139-5ebeacd6b1f1">

