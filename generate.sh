#!/bin/bash
export PATH="$PATH:$(go env GOPATH)/bin"
~/Downloads/protoc-22/bin/protoc --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./**/*.proto
