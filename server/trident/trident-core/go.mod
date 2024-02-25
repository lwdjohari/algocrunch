module trident-core

go 1.21

require (
	nvm-gocore v0.2.3
	nvm-sqlxe v0.2.3
	trident-data v0.2.3
)

require (
	github.com/cncf/xds/go v0.0.0-20231128003011-0fa0005c9caa // indirect
	github.com/envoyproxy/go-control-plane v0.12.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/go-sql-driver/mysql v1.7.1 //indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jmoiron/sqlx v1.3.5 //indirect
	github.com/lib/pq v1.10.9 //indirect
	github.com/mattn/go-sqlite3 v1.14.22 //indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240221002015-b0ce06bbee7c // indirect
	google.golang.org/grpc v1.62.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace nvm-gocore => ../../nvm-gocore

replace nvm-sqlxe => ../../nvm-sqlxe

replace trident-data => ../trident-data
