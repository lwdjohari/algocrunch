module trident-server

go 1.18

require (
  trident-core v0.2.3
  google.golang.org/grpc v1.35.0
  google.golang.org/protobuf v1.25.0

)

replace trident-core => ../trident-core