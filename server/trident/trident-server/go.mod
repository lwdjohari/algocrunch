module trident-server

go 1.21

require (
	google.golang.org/grpc v1.62.0
	trident-core v0.2.3
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240221002015-b0ce06bbee7c // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

replace trident-core => ../trident-core
