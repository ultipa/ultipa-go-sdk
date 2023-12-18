## Trouble Shooting
### grpc issues
This sdk depends on Protocol Buffer and grpc, so we need to compile the proto file first:<br>
go install google.golang.org/grpc/cmd/protoc-gen-go@latest<br>
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest<br>

then execute the build shell:<br>
cd ./rpc <br>
sh ./build-proto.sh

after this, remove protoc-gen-go and protoc-gen-go-grpc dependencies from go.mod



